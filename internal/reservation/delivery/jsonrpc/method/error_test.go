package method

import (
	"errors"
	"testing"

	"github.com/1nkh3art1/goods-reservation-service/internal/cerror"

	"github.com/osamingo/jsonrpc/v2"
	"github.com/stretchr/testify/require"
)

func TestError(t *testing.T) {
	err := toJsonrpcError(errors.New("some error"))
	require.Equal(t, err, jsonrpc.ErrInternal())

	msg := "something not found"
	cerr := cerror.NotFound.New(msg)
	expected := &jsonrpc.Error{Code: jsonrpc.ErrorCode(cerror.NotFound), Message: msg}

	err = toJsonrpcError(cerr)
	require.Equal(t, expected, err)

}
