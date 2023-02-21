package service

import (
	"context"
	"errors"
	"testing"

	"github.com/1nkh3art1/goods-reservation-service/internal/cerror"
	"github.com/1nkh3art1/goods-reservation-service/internal/reservation/domain"
	"github.com/1nkh3art1/goods-reservation-service/internal/reservation/mock"
	"github.com/1nkh3art1/goods-reservation-service/pkg/logger"

	"github.com/stretchr/testify/require"
)

func TestReleaseGoods(t *testing.T) {
	mockStorage := &mock.Storage{}
	rs := NewReservationService(mockStorage, logger.NewLoggerStub())

	ctx := context.Background()

	result, err := rs.ReleaseGoods(ctx, &domain.Request{WarehouseID: domain.WarehouseID(-1)})
	require.Nil(t, result)
	require.Error(t, err)

	request := &domain.Request{
		WarehouseID: domain.WarehouseID(1),
		Codes:       []string{"uniqueGoodCode01", "uniqueGoodCode02"},
	}

	err = errors.New("some error")

	mockStorage.On("WarehouseAvailability", ctx, request.WarehouseID).Return(false, err)

	result, err = rs.ReleaseGoods(ctx, request)
	require.Nil(t, result)
	require.Error(t, err)

	request = &domain.Request{
		WarehouseID: domain.WarehouseID(1),
		Codes:       []string{"uniqueGoodCode01", "uniqueGoodCode02"},
	}

	mockStorage.On("WarehouseAvailability", ctx, request.WarehouseID).Return(true, nil)
	mockStorage.On("ReleaseGoods", ctx, request).Return(nil, err)

	result, err = rs.ReleaseGoods(ctx, request)
	require.Nil(t, result)
	require.Error(t, err)

	request = &domain.Request{
		WarehouseID: domain.WarehouseID(3),
		Codes:       []string{"uniqueGoodCode01", "uniqueGoodCode02"},
	}

	responce := &domain.Responce{
		Amount: map[string]int{"uniqueGoodCode01": 1, "uniqueGoodCode02": 1},
	}

	mockStorage.On("WarehouseAvailability", ctx, request.WarehouseID).Return(true, nil)
	mockStorage.On("ReleaseGoods", ctx, request).Return(responce, nil)

	actual, err := rs.ReleaseGoods(ctx, request)
	require.Equal(t, actual, responce)
	require.Nil(t, err)
}

func TestReserveGoods(t *testing.T) {
	mockStorage := &mock.Storage{}
	rs := NewReservationService(mockStorage, logger.NewLoggerStub())

	ctx := context.Background()

	invalidRequest := &domain.Request{WarehouseID: domain.WarehouseID(-1)}
	result, err := rs.ReserveGoods(ctx, invalidRequest)
	require.Nil(t, result)
	require.Error(t, err)

	err = errors.New("some error")

	request := &domain.Request{
		WarehouseID: domain.WarehouseID(4),
		Codes:       []string{"uniqueGoodCode03", "uniqueGoodCode04"},
	}

	mockStorage.On("WarehouseAvailability", ctx, request.WarehouseID).Return(false, err)

	result, err = rs.ReserveGoods(ctx, request)
	require.Nil(t, result)
	require.Error(t, err)

	request = &domain.Request{
		WarehouseID: domain.WarehouseID(5),
		Codes:       []string{"uniqueGoodCode03", "uniqueGoodCode04"},
	}

	mockStorage.On("WarehouseAvailability", ctx, request.WarehouseID).Return(true, nil)
	mockStorage.On("ReserveGoods", ctx, request).Return(nil, err)

	result, err = rs.ReserveGoods(ctx, request)
	require.Nil(t, result)
	require.Error(t, err)

	request = &domain.Request{
		WarehouseID: domain.WarehouseID(6),
		Codes:       []string{"uniqueGoodCode03", "uniqueGoodCode04"},
	}

	responce := &domain.Responce{
		Amount: map[string]int{"uniqueGoodCode01": 1, "uniqueGoodCode02": 1},
	}

	mockStorage.On("WarehouseAvailability", ctx, request.WarehouseID).Return(true, nil)
	mockStorage.On("ReserveGoods", ctx, request).Return(responce, nil)

	actual, err := rs.ReserveGoods(ctx, request)
	require.Equal(t, actual, responce)
	require.Nil(t, err)
}

func TestWarehouseAvailability(t *testing.T) {
	mockStorage := &mock.Storage{}
	rs := NewReservationService(mockStorage, logger.NewLoggerStub())

	ctx := context.Background()

	err := rs.WarehouseAvailability(ctx, -1)
	require.Error(t, err)

	err = errors.New("some error")

	mockStorage.On("WarehouseAvailability", ctx, domain.WarehouseID(1)).Return(false, err)
	mockStorage.On("WarehouseAvailability", ctx, domain.WarehouseID(2)).Return(false, nil)
	mockStorage.On("WarehouseAvailability", ctx, domain.WarehouseID(3)).Return(true, nil)

	err = rs.WarehouseAvailability(ctx, domain.WarehouseID(1))
	require.Error(t, err)

	var ce cerror.CustomError
	err = rs.WarehouseAvailability(ctx, domain.WarehouseID(2))
	require.ErrorAs(t, err, &ce)

	err = rs.WarehouseAvailability(ctx, domain.WarehouseID(3))
	require.NoError(t, err)
}

func TestGoodsAmount(t *testing.T) {
	mockStorage := &mock.Storage{}
	rs := NewReservationService(mockStorage, logger.NewLoggerStub())

	ctx := context.Background()

	responce, err := rs.GoodsAmount(ctx, -1)
	require.Nil(t, responce)
	require.Error(t, err)

	mockStorage.On("WarehouseAvailability", ctx, domain.WarehouseID(1)).Return(false, nil)
	mockStorage.On("WarehouseAvailability", ctx, domain.WarehouseID(2)).Return(true, nil)
	mockStorage.On("WarehouseAvailability", ctx, domain.WarehouseID(3)).Return(true, nil)

	responce = &domain.Responce{Amount: map[string]int{"uniqueGoodCode01": 5}}

	err = errors.New("some error")

	mockStorage.On("GoodsAmount", ctx, domain.WarehouseID(2)).Return(nil, err)
	mockStorage.On("GoodsAmount", ctx, domain.WarehouseID(3)).Return(responce, nil)

	result, err := rs.GoodsAmount(ctx, domain.WarehouseID(1))
	require.Nil(t, result)
	require.Error(t, err)

	result, err = rs.GoodsAmount(ctx, domain.WarehouseID(2))
	require.Nil(t, result)
	require.Error(t, err)

	actual, err := rs.GoodsAmount(ctx, domain.WarehouseID(3))
	require.Equal(t, responce, actual)
	require.NoError(t, err)

}
