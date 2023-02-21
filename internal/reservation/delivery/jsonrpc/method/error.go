package method

import (
	"github.com/1nkh3art1/goods-reservation-service/internal/cerror"

	"github.com/osamingo/jsonrpc/v2"
	"github.com/pkg/errors"
)

// toJsonrpcError преобразовывает кастомную ошибку в ошибку JSONRPC
// если может или отдает внутреннюю ошибку сервера
func toJsonrpcError(err error) *jsonrpc.Error {
	var ce cerror.CustomError
	if errors.As(err, &ce) {
		return &jsonrpc.Error{Code: jsonrpc.ErrorCode(ce.Code()), Message: ce.Error()}
	} else {
		return jsonrpc.ErrInternal()
	}
}
