package http

import (
	"net/http"

	"github.com/1nkh3art1/goods-reservation-service/internal/reservation/delivery/http/request"
)

func RequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := request.AttachReqID(r.Context())
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
		h := w.Header()
		h.Add(request.HTTPHeaderNameRequestID, request.GetReqID(ctx))
	})
}
