package storage

import (
	"context"

	"log"

	"github.com/1nkh3art1/goods-reservation-service/internal/cerror"
	"github.com/1nkh3art1/goods-reservation-service/internal/reservation/domain"
	"github.com/pkg/errors"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type goodStorage struct {
	pool *pgxpool.Pool
}

func NewGoodStorage(pool *pgxpool.Pool) *goodStorage {
	return &goodStorage{pool: pool}
}

func (s *goodStorage) ReserveGoods(ctx context.Context, req *domain.Request) (*domain.Responce, error) {
	amount, err := s.update(ctx, req, -1)
	if err != nil {
		return nil, err
	}

	return &domain.Responce{Amount: amount}, nil
}

func (s *goodStorage) ReleaseGoods(ctx context.Context, req *domain.Request) (*domain.Responce, error) {
	amount, err := s.update(ctx, req, 1)
	if err != nil {
		return nil, err
	}

	return &domain.Responce{Amount: amount}, nil
}

// GoodsAmount получает идентификатор склада и возвращает количество оставшихся на складе товаров
func (s *goodStorage) GoodsAmount(ctx context.Context, warehouseID domain.WarehouseID) (*domain.Responce, error) {
	rows, err := s.pool.Query(ctx, goodsAmountInWarehouse, int(warehouseID))
	if err != nil {
		return nil, errors.WithMessage(err, "pool query")
	}
	defer rows.Close()

	goods := make(map[string]int, 0)
	for rows.Next() {
		var (
			good   string
			amount int
		)

		if err := rows.Scan(&good, &amount); err != nil {
			return nil, errors.WithMessage(err, "rows scan")
		}

		goods[good] = amount
	}

	if err := rows.Err(); err != nil {
		return nil, errors.WithMessage(err, "rows error")
	}

	return &domain.Responce{Amount: goods}, nil
}

func isConstraintError(err error) bool {
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		if pgErr.Code == "23514" {
			return true
		}
	}

	return false
}

// checkWarehouseAvailability получает идентификатор склада и проверяет доступен ли склад
func (s *goodStorage) WarehouseAvailability(ctx context.Context, warehouseID domain.WarehouseID) (bool, error) {
	var status bool
	if err := s.pool.QueryRow(ctx, warehouseAvailability, warehouseID).Scan(&status); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return false, cerror.NotFound.Newf("warehouse with ID %d not found", warehouseID)
		}

		return false, errors.WithMessage(err, "pool query row")
	}

	return status, nil
}

func (s *goodStorage) updateAmountInWarehouses(
	ctx context.Context,
	tx pgx.Tx,
	warehouseID int,
	code string,
	amount int) error {
	cmd, err := tx.Exec(ctx, updateAmountOfGoodInWarehouse, warehouseID, code, amount)
	if err != nil {
		if isConstraintError(err) {
			return cerror.BadRequest.Newf("good %s is out of stock in warehouse %d", code, warehouseID)
		}

		return errors.WithMessage(err, "pool exec: update amount of good in warehouse")
	}

	if cmd.RowsAffected() == 0 {
		return cerror.NotFound.Newf("good %s is not found in warehouse %d", code, warehouseID)
	}

	return nil
}

func (s *goodStorage) updateAmountInReservations(
	ctx context.Context,
	tx pgx.Tx,
	warehouseID int,
	code string,
	amount int) error {
	cmd, err := tx.Exec(ctx, updateAmountOfGoodInReservations, warehouseID, code, amount)
	if err != nil {
		if isConstraintError(err) {
			return cerror.BadRequest.Newf("good %s is out of reservations in warehouse %d", code, warehouseID)
		}

		return errors.WithMessage(err, "pool exec: update amount of good in reservations")
	}

	if cmd.RowsAffected() == 0 {
		return cerror.NotFound.Newf("good %s is not found in reservations of warehouse %d", code, warehouseID)
	}

	return nil
}

func (s *goodStorage) update(ctx context.Context, req *domain.Request, amount int) (map[string]int, error) {
	tx, err := s.pool.BeginTx(ctx, pgx.TxOptions{IsoLevel: pgx.RepeatableRead})
	if err != nil {
		return nil, errors.WithMessage(err, "pool begin tx")
	}

	defer func() {
		if r := recover(); r != nil {
			log.Println(err)
		}

		_ = tx.Rollback(context.Background())
	}()

	result := make(map[string]int)
	for _, code := range req.Codes {
		err := s.updateAmountInWarehouses(ctx, tx, int(req.WarehouseID), code, amount)
		if err != nil {
			return nil, err
		}

		err = s.updateAmountInReservations(ctx, tx, int(req.WarehouseID), code, -amount)
		if err != nil {
			return nil, err
		}

		result[code]++
	}

	if err := tx.Commit(ctx); err != nil {
		return nil, errors.WithMessage(err, "tx commit")
	}

	return result, nil
}
