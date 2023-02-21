package service

import (
	"context"

	"github.com/1nkh3art1/goods-reservation-service/internal/cerror"
	"github.com/1nkh3art1/goods-reservation-service/internal/reservation"
	"github.com/1nkh3art1/goods-reservation-service/internal/reservation/domain"
	"github.com/1nkh3art1/goods-reservation-service/pkg/logger"

	"github.com/pkg/errors"
)

// reservationService представляет собой сервис
// для резервирования/освобождения товаров на скаладах
type reservationService struct {
	storage reservation.Storage
	log     logger.Logger
}

// NewReservationService создает новый экземпляр reservationService
func NewReservationService(storage reservation.Storage, log logger.Logger) *reservationService {
	return &reservationService{storage: storage, log: log}
}

// ReserveGoods резервирует товары на складе
func (s *reservationService) ReserveGoods(ctx context.Context, req *domain.Request) (*domain.Responce, error) {
	if err := req.Validate(); err != nil {
		return nil, errors.WithMessage(err, "request validate")
	}

	if err := s.WarehouseAvailability(ctx, req.WarehouseID); err != nil {
		return nil, errors.WithMessage(err, "service warehouse availability")
	}

	result, err := s.storage.ReserveGoods(ctx, req)
	if err != nil {
		return nil, errors.WithMessage(err, "storage reserve goods")
	}

	return result, nil
}

// ReserveGoods освобождает товары на складе
func (s *reservationService) ReleaseGoods(ctx context.Context, req *domain.Request) (*domain.Responce, error) {
	if err := req.Validate(); err != nil {
		return nil, errors.WithMessage(err, "request validate")
	}

	if err := s.WarehouseAvailability(ctx, req.WarehouseID); err != nil {
		return nil, errors.WithMessage(err, "service warehouse availability")
	}

	result, err := s.storage.ReleaseGoods(ctx, req)
	if err != nil {
		return nil, errors.WithMessage(err, "storage release goods")
	}

	return result, nil
}

// GoodsAmount проверяет остатки товаров на складе
func (s *reservationService) GoodsAmount(ctx context.Context, warehouseID domain.WarehouseID) (*domain.Responce, error) {
	if err := warehouseID.Validate(); err != nil {
		return nil, errors.WithMessage(err, "request validate")
	}

	if err := s.WarehouseAvailability(ctx, warehouseID); err != nil {
		return nil, errors.WithMessage(err, "service warehouse availability")
	}

	result, err := s.storage.GoodsAmount(ctx, warehouseID)
	if err != nil {
		return nil, errors.WithMessage(err, "storage goods amount in warehouse")
	}

	return result, nil
}

// warehouseAvailability проверяет доступен ли в настоящее время склад
// для дальнейших операций резервирования/освобождения
func (s *reservationService) WarehouseAvailability(ctx context.Context, warehouseID domain.WarehouseID) error {
	if err := warehouseID.Validate(); err != nil {
		return errors.WithMessage(err, "request validate")
	}

	available, err := s.storage.WarehouseAvailability(ctx, warehouseID)
	if err != nil {
		return errors.WithMessage(err, "storage warehouse availability")
	}

	if !available {
		return cerror.BadRequest.Newf("warehouse with ID '%d' is unavailable", warehouseID)
	}

	return nil
}
