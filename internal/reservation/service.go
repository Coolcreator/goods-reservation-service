package reservation

import (
	"context"

	"github.com/1nkh3art1/goods-reservation-service/internal/reservation/domain"
)

type Service interface {
	GoodsAmount(ctx context.Context, warehouseID domain.WarehouseID) (*domain.Responce, error)
	ReserveGoods(ctx context.Context, req *domain.Request) (*domain.Responce, error)
	ReleaseGoods(ctx context.Context, req *domain.Request) (*domain.Responce, error)
}
