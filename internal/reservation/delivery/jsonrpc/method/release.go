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
	ReleaseHandler struct {
		service reservation.Service
		log     logger.Logger
	}

	ReleaseParams struct {
		WarehouseID int      `json:"warehouse_id"`
		Codes       []string `json:"codes"`
	}

	ReleaseResult struct {
		WarehouseID int            `json:"warehouse_id"`
		Released    map[string]int `json:"released"`
	}
)

func (h ReleaseHandler) Name() string {
	return "release"
}

func (h ReleaseHandler) Params() any {
	return ReleaseParams{}
}

func (h ReleaseHandler) Result() any {
	return ReleaseResult{}
}

func NewReleaseHandler(service reservation.Service, log logger.Logger) *ReleaseHandler {
	return &ReleaseHandler{service: service, log: log}
}

func (h ReleaseHandler) ServeJSONRPC(c context.Context, params *json.RawMessage) (any, *jsonrpc.Error) {
	var p ReleaseParams
	if err := jsonrpc.Unmarshal(params, &p); err != nil {
		return nil, err
	}

	if len(p.Codes) == 0 {
		return &ReleaseResult{
			WarehouseID: p.WarehouseID,
			Released:    map[string]int{},
		}, nil
	}

	reqID := request.GetReqID(c)

	h.log.Infof("[req: %s, method: %s]: release goods for warehouse %d\n", reqID, h.Name(), p.WarehouseID)

	result, err := h.service.ReleaseGoods(c, &domain.Request{
		WarehouseID: domain.WarehouseID(p.WarehouseID),
		Codes:       p.Codes,
	})
	if err != nil {
		h.log.Errorf("[req: %s, method: %s]: error: %v\n", reqID, h.Name(), err)
		return nil, toJsonrpcError(err)
	}

	h.log.Infof("[req: %s, method: %s]: success", reqID, h.Name())

	return &ReleaseResult{
		WarehouseID: p.WarehouseID,
		Released:    result.Amount,
	}, nil
}
