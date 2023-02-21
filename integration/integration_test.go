package integration

import (
	"bytes"
	"context"
	"io"

	"net/http/httptest"
	"testing"
	"time"

	middleware "github.com/1nkh3art1/goods-reservation-service/internal/reservation/delivery/http"
	jrpc "github.com/1nkh3art1/goods-reservation-service/internal/reservation/delivery/jsonrpc"
	"github.com/go-chi/chi/v5"

	"github.com/1nkh3art1/goods-reservation-service/internal/reservation/delivery/jsonrpc/method"
	"github.com/1nkh3art1/goods-reservation-service/internal/reservation/service"
	"github.com/1nkh3art1/goods-reservation-service/internal/reservation/storage"
	"github.com/1nkh3art1/goods-reservation-service/migration"
	"github.com/1nkh3art1/goods-reservation-service/pkg/logger"
	"github.com/1nkh3art1/goods-reservation-service/pkg/postgres"
	"github.com/osamingo/jsonrpc/v2"
	"github.com/stretchr/testify/suite"
)

type TestSuite struct {
	suite.Suite
	container *PostgresContainer
	server    *httptest.Server
}

func (s *TestSuite) SetupSuite() {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	container, err := NewPostgresContainer(ctx)
	s.Require().NoError(err)

	s.container = container

	err = migration.Migrate(container.GetConnString())
	s.Require().NoError(err)

	poolCfg := &postgres.Config{
		Host:     container.Config.Host,
		Port:     container.Config.MappedPort,
		User:     container.Config.User,
		Password: container.Config.Password,
		DB:       container.Config.Database,
		SSLMode:  container.Config.SSLMode,
	}

	pool, err := postgres.NewConnPool(poolCfg)
	s.Require().NoError(err)

	log := logger.NewLoggerStub()

	storage := storage.NewGoodStorage(pool)
	service := service.NewReservationService(storage, log)

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

	s.Require().NoError(err)

	s.server = httptest.NewServer(m)
}

func (s *TestSuite) TearDownSuite() {
	ctx, ctxCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer ctxCancel()

	s.Require().NoError(s.container.Terminate(ctx))

	s.server.Close()
}

func TestSuite_Run(t *testing.T) {
	if testing.Short() {
		t.Skip("Skip integration tests")
	}

	suite.Run(t, new(TestSuite))
}

func (s *TestSuite) TestAmount() {
	msg := `{"jsonrpc": "2.0", "method": "amount", "params": {"warehouse_id": 1}}`

	res1, err := s.server.Client().Post(s.server.URL+"/rpc", "application/json", bytes.NewBufferString(msg))
	defer res1.Body.Close()
	s.Require().NoError(err)
	s.Require().NotNil(res1)

	data, err := io.ReadAll(res1.Body)
	s.Require().NoError(err)

	expected := `{"jsonrpc":"2.0","result":{"warehouse_id":1,"amount":{"вода":10,"молоко":5,"хлеб":10,"яйца":5}}}`
	s.Require().JSONEq(expected, string(data))

	msg = `{"jsonrpc": "2.0", "method": "amount", "params": {"warehouse_id": "not valid"}}`
	res2, err := s.server.Client().Post(s.server.URL+"/rpc", "application/json", bytes.NewBufferString(msg))
	defer res2.Body.Close()
	s.Require().NoError(err)
	s.Require().NotNil(res2)

	data, err = io.ReadAll(res2.Body)
	s.Require().NoError(err)

	expected = `{"jsonrpc":"2.0","error":{"code":-32602,"message":"Invalid params"}}`
	s.Require().JSONEq(expected, string(data))
}

func (s *TestSuite) TestReserveAndRelease() {
	msg := `{"jsonrpc": "2.0", "method": "reserve", "params": {"warehouse_id": 1, "codes": ["uniqueGoodCode01"]}}`

	res1, err := s.server.Client().Post(s.server.URL+"/rpc", "application/json", bytes.NewBufferString(msg))
	defer res1.Body.Close()
	s.Require().NoError(err)
	s.Require().NotNil(res1)

	data, err := io.ReadAll(res1.Body)
	s.Require().NoError(err)

	expected := `{"jsonrpc":"2.0","result":{"warehouse_id":1,"reserved":{"uniqueGoodCode01":1}}}`
	s.Require().JSONEq(expected, string(data))

	res2, err := s.server.Client().Post(s.server.URL+"/rpc", "application/json", bytes.NewBufferString(msg))
	defer res2.Body.Close()
	s.Require().NoError(err)
	s.Require().NotNil(res2)

	data, err = io.ReadAll(res2.Body)
	s.Require().NoError(err)

	expected = `{"jsonrpc":"2.0","result":{"warehouse_id":1,"reserved":{"uniqueGoodCode01":1}}}`
	s.Require().JSONEq(expected, string(data))

	msg = `{"jsonrpc": "2.0", "method": "release", "params": {"warehouse_id": 1, "codes": ["uniqueGoodCode01"]}}`

	res3, err := s.server.Client().Post(s.server.URL+"/rpc", "application/json", bytes.NewBufferString(msg))
	defer res3.Body.Close()
	s.Require().NoError(err)
	s.Require().NotNil(res3)

	data, err = io.ReadAll(res3.Body)
	s.Require().NoError(err)

	expected = `{"jsonrpc":"2.0","result":{"warehouse_id":1,"released":{"uniqueGoodCode01":1}}}`
	s.Require().JSONEq(expected, string(data))

	res4, err := s.server.Client().Post(s.server.URL+"/rpc", "application/json", bytes.NewBufferString(msg))
	defer res4.Body.Close()
	s.Require().NoError(err)
	s.Require().NotNil(res4)

	data, err = io.ReadAll(res4.Body)
	s.Require().NoError(err)

	expected = `{"jsonrpc":"2.0","result":{"warehouse_id":1,"released":{"uniqueGoodCode01":1}}}`
	s.Require().JSONEq(expected, string(data))

}
