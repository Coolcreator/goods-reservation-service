package request

import (
	"context"

	"github.com/google/uuid"
)

type ctxRequestIDKey string

const (
	HTTPHeaderNameRequestID                 = "X-Request-ID"
	ContextKeyReqID         ctxRequestIDKey = "requestID"
)

func GetReqID(ctx context.Context) string {
	reqID := ctx.Value(ContextKeyReqID)

	if rs, ok := reqID.(string); ok {
		return rs
	}

	return ""
}

func AttachReqID(ctx context.Context) context.Context {

	reqID := uuid.New()

	return context.WithValue(ctx, ContextKeyReqID, reqID.String())
}
