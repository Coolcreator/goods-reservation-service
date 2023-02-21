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

func TestReleaseMethodFunctions(t *testing.T) {
	rh := NewReleaseHandler(&mock.Service{}, logger.NewLoggerStub())

	require.Equal(t, "release", rh.Name())
	require.Equal(t, ReleaseParams{}, rh.Params())
	require.Equal(t, ReleaseResult{}, rh.Result())
}

func TestReleaseServeJSONRPC(t *testing.T) {
	rs := &mock.Service{}

	rh := NewReleaseHandler(rs, logger.NewLoggerStub())

	c := context.Background()

	invalidParams := json.RawMessage(`[]`)
	result, err := rh.ServeJSONRPC(c, &invalidParams)
	require.Nil(t, result)
	require.Equal(t, jsonrpc.ErrInvalidParams(), err)

	expected := &ReleaseResult{
		WarehouseID: 3,
		Released: map[string]int{
			"uniqueGoodCode01": 1,
			"uniqueGoodCode02": 1,
		},
	}

	req := &domain.Request{
		WarehouseID: domain.WarehouseID(3),
		Codes:       []string{"uniqueGoodCode01", "uniqueGoodCode02"},
	}

	responce := &domain.Responce{
		Amount: map[string]int{"uniqueGoodCode01": 1, "uniqueGoodCode02": 1},
	}

	rs.On("ReleaseGoods", c, req).Return(responce, nil)

	params := json.RawMessage(`{
		"warehouse_id": 3,
		"codes": ["uniqueGoodCode01", "uniqueGoodCode02"]
	}`)

	actual, err := rh.ServeJSONRPC(c, &params)
	require.Equal(t, expected, actual)
	require.Nil(t, err)

	e := errors.New("something went wrong")

	req = &domain.Request{WarehouseID: 10, Codes: []string{"uniqueGoodCode01", "uniqueGoodCode02"}}

	rs.On("ReleaseGoods", c, req).Return(nil, e)

	params = json.RawMessage(`{
		"warehouse_id": 10,
		"codes": ["uniqueGoodCode01", "uniqueGoodCode02"]
	}`)

	result, err = rh.ServeJSONRPC(c, &params)
	require.Nil(t, result)
	require.Error(t, err)

}
