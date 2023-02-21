package method

import (
	"context"
	"encoding/json"
	"errors"
	"testing"

	"github.com/1nkh3art1/goods-reservation-service/internal/reservation/domain"
	"github.com/1nkh3art1/goods-reservation-service/internal/reservation/mock"
	"github.com/1nkh3art1/goods-reservation-service/pkg/logger"

	"github.com/osamingo/jsonrpc/v2"
	"github.com/stretchr/testify/require"
)

func TestReserveMethodFunctions(t *testing.T) {
	rh := NewReserveHandler(&mock.Service{}, logger.NewLoggerStub())

	require.Equal(t, "reserve", rh.Name())
	require.Equal(t, ReserveParams{}, rh.Params())
	require.Equal(t, ReserveResult{}, rh.Result())
}

func TestReserveServeJSONRPC(t *testing.T) {
	rs := &mock.Service{}

	rh := NewReserveHandler(rs, logger.NewLoggerStub())

	c := context.Background()

	invalidParams := json.RawMessage(`[]`)
	result, err := rh.ServeJSONRPC(c, &invalidParams)
	require.Nil(t, result)
	require.Equal(t, jsonrpc.ErrInvalidParams(), err)

	expected := &ReserveResult{
		WarehouseID: 2,
		Reserved: map[string]int{
			"uniqueGoodCode01": 1,
			"uniqueGoodCode02": 1,
		},
	}

	req := &domain.Request{
		WarehouseID: domain.WarehouseID(2),
		Codes:       []string{"uniqueGoodCode01", "uniqueGoodCode02"},
	}

	responce := &domain.Responce{
		Amount: map[string]int{"uniqueGoodCode01": 1, "uniqueGoodCode02": 1},
	}

	rs.On("ReserveGoods", c, req).Return(responce, nil)

	params := json.RawMessage(`{
		"warehouse_id": 2,
		"codes": ["uniqueGoodCode01", "uniqueGoodCode02"]
	}`)

	actual, err := rh.ServeJSONRPC(c, &params)
	require.Equal(t, expected, actual)
	require.Nil(t, err)

	e := errors.New("something went wrong")

	req = &domain.Request{WarehouseID: 10, Codes: []string{"uniqueGoodCode01", "uniqueGoodCode02"}}

	rs.On("ReserveGoods", c, req).Return(nil, e)

	params = json.RawMessage(`{
		"warehouse_id": "warehouseID10",
		"codes": ["uniqueGoodCode01", "uniqueGoodCode02"]
	}`)

	result, err = rh.ServeJSONRPC(c, &params)
	require.Nil(t, result)
	require.Error(t, err)

}
