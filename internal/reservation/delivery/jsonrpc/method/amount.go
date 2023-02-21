package method

import (
	"context"
	"encoding/json"

	"github.com/1nkh3art1/goods-reservation-service/internal/reservation"
	"github.com/1nkh3art1/goods-reservation-service/internal/reservation/delivery/http/request"
	"github.com/1nkh3art1/goods-reservation-service/internal/reservation/domain"
	"github.com/1nkh3art1/goods-reservation-service/pkg/logger"

	"github.com/osamingo/jsonrpc/v2"
)

type (
	AmountHandler struct {
		service reservation.Service
		log     logger.Logger
	}

	AmountParams struct {
		WarehouseID int `json:"warehouse_id"`
	}

	AmountResult struct {
		WarehouseID int            `json:"warehouse_id"`
		Amount      map[string]int `json:"amount"`
	}
)

func (h AmountHandler) Name() string {
	return "amount"
}

func (h AmountHandler) Params() any {
	return AmountParams{}
}

func (h AmountHandler) Result() any {
	return AmountResult{}
}

func NewAmountHandler(service reservation.Service, log logger.Logger) *AmountHandler {
	return &AmountHandler{service: service, log: log}
}

func (h AmountHandler) ServeJSONRPC(c context.Context, params *json.RawMessage) (any, *jsonrpc.Error) {
	var p AmountParams
	if err := jsonrpc.Unmarshal(params, &p); err != nil {
		return nil, err
	}

	reqID := request.GetReqID(c)

	h.log.Infof("[req: %s, method: %s]: get goods amounts for warehouse '%d'\n", reqID, h.Name(), p.WarehouseID)

	result, err := h.service.GoodsAmount(c, domain.WarehouseID(p.WarehouseID))
	if err != nil {
		h.log.Errorf("[req: %s, method: %s]: error: %v\n", reqID, h.Name(), err)
		return nil, toJsonrpcError(err)
	}

	h.log.Infof("[req: %s, method: %s]: success", reqID, h.Name())

	return &AmountResult{
		WarehouseID: p.WarehouseID,
		Amount:      result.Amount,
	}, nil
}
