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
	ReserveHandler struct {
		service reservation.Service
		log     logger.Logger
	}

	ReserveParams struct {
		WarehouseID int      `json:"warehouse_id"`
		Codes       []string `json:"codes"`
	}

	ReserveResult struct {
		WarehouseID int            `json:"warehouse_id"`
		Reserved    map[string]int `json:"reserved"`
	}
)

func (h ReserveHandler) Name() string {
	return "reserve"
}

func (h ReserveHandler) Params() any {
	return ReserveParams{}
}

func (h ReserveHandler) Result() any {
	return ReserveResult{}
}

func NewReserveHandler(service reservation.Service, log logger.Logger) *ReserveHandler {
	return &ReserveHandler{service: service, log: log}
}

func (h ReserveHandler) ServeJSONRPC(c context.Context, params *json.RawMessage) (any, *jsonrpc.Error) {
	var p ReserveParams
	if err := jsonrpc.Unmarshal(params, &p); err != nil {
		return nil, err
	}

	reqID := request.GetReqID(c)

	h.log.Infof("[req: %s, method: %s]: reserve goods for warehouse %d\n", reqID, h.Name(), p.WarehouseID)

	result, err := h.service.ReserveGoods(c, &domain.Request{
		WarehouseID: domain.WarehouseID(p.WarehouseID),
		Codes:       p.Codes,
	})
	if err != nil {
		h.log.Errorf("[req: %s, method: %s]: error: %v\n", reqID, h.Name(), err)
		return nil, toJsonrpcError(err)
	}

	h.log.Infof("[req: %s, method: %s]: success", reqID, h.Name())

	return &ReserveResult{
		WarehouseID: p.WarehouseID,
		Reserved:    result.Amount,
	}, nil
}
