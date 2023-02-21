package app

import (
	"context"
	"os"
	"os/signal"

	"github.com/1nkh3art1/goods-reservation-service/config"
	middleware "github.com/1nkh3art1/goods-reservation-service/internal/reservation/delivery/http"
	jrpc "github.com/1nkh3art1/goods-reservation-service/internal/reservation/delivery/jsonrpc"
	"github.com/1nkh3art1/goods-reservation-service/internal/reservation/delivery/jsonrpc/method"
	"github.com/1nkh3art1/goods-reservation-service/internal/reservation/service"
	"github.com/1nkh3art1/goods-reservation-service/internal/reservation/storage"
	"github.com/1nkh3art1/goods-reservation-service/pkg/logger"
	"github.com/1nkh3art1/goods-reservation-service/pkg/postgres"
	"github.com/1nkh3art1/goods-reservation-service/pkg/server"
	"github.com/go-chi/chi/v5"
	"github.com/osamingo/jsonrpc/v2"
)

// Run инициализирует и запускает все сервисы
func Run(cfg *config.Config) {
	log := logger.NewZapSugaredlogger(&logger.Config{
		Development: cfg.Logger.Development,
		Level:       cfg.Logger.Level,
		Encoding:    cfg.Logger.Encoding,
	})

	log.Init()

	log.Info("logger initialized")

	log.Infof("database: '%s':'%s'", cfg.Postgres.Host, cfg.Postgres.Port)

	pool, err := postgres.NewConnPool(&postgres.Config{
		Host:     cfg.Postgres.Host,
		Port:     cfg.Postgres.Port,
		User:     cfg.Postgres.User,
		Password: cfg.Postgres.Password,
		DB:       cfg.Postgres.DB,
		SSLMode:  cfg.Postgres.SSLMode,
	})
	defer pool.Close()
	if err != nil {
		log.Fatalf("postgres new conn pool: %v", err)
	}

	log.Info("established connection to postgres")

	storage := storage.NewGoodStorage(pool)

	log.Info("good storage initialized")

	service := service.NewReservationService(storage, log)

	log.Info("reservation service initialized")

	mr := jsonrpc.NewMethodRepository()

	if err = jrpc.RegisterMethods(mr,
		method.NewAmountHandler(service, log),
		method.NewReserveHandler(service, log),
		method.NewReleaseHandler(service, log),
	); err != nil {
		log.Fatalf("register methods: %v\n", err)
	}

	log.Info("registered methods")

	m := chi.NewMux()

	m.Use(middleware.RequestID)
	m.Handle("/rpc", mr)

	log.Info("created new serve mux and registered jsonrpc handlers on it")

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	s := server.NewServer(m)

	log.Info("server initialized")

	s.Run()

	log.Info("started http server")

	<-ctx.Done()

	s.Shutdown()
}
