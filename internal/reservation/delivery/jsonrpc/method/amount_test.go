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

func TestMethodFunctions(t *testing.T) {
	ah := NewAmountHandler(&mock.Service{}, logger.NewLoggerStub())

	require.Equal(t, "amount", ah.Name())
	require.Equal(t, AmountParams{}, ah.Params())
	require.Equal(t, AmountResult{}, ah.Result())
}

func TestAmountServeJSONRPC(t *testing.T) {
	rs := &mock.Service{}

	ah := NewAmountHandler(rs, logger.NewLoggerStub())

	c := context.Background()

	invalidParams := json.RawMessage(`[]`)
	result, err := ah.ServeJSONRPC(c, &invalidParams)
	require.Nil(t, result)
	require.Equal(t, jsonrpc.ErrInvalidParams(), err)

	amount := map[string]int{"uniqueGoodCode01": 10}

	mockResult := &domain.Responce{Amount: amount}

	rs.On("GoodsAmount", c, domain.WarehouseID(1)).Return(mockResult, nil)

	params := json.RawMessage(`{"warehouse_id": 1}`)

	result = &AmountResult{WarehouseID: 1, Amount: amount}

	actual, err := ah.ServeJSONRPC(c, &params)
	require.Equal(t, result, actual)
	require.Nil(t, err)

	e := errors.New("something went wrong")

	rs.On("GoodsAmount", c, domain.WarehouseID(10)).Return(nil, e)

	params = json.RawMessage(`{"warehouse_id": 10}`)

	result, err = ah.ServeJSONRPC(c, &params)
	require.Nil(t, result)
	require.Error(t, err)

}
