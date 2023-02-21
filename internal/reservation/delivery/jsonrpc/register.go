package jsonrpc

import (
	"github.com/osamingo/jsonrpc/v2"
	"github.com/pkg/errors"
)

// RegisterMethods регистрирует методы хендлеров в репозитории методов
func RegisterMethods(mr *jsonrpc.MethodRepository, methods ...Method) error {
	for _, m := range methods {
		err := mr.RegisterMethod(m.Name(), m, m.Params(), m.Result())
		if err != nil {
			return errors.WithMessagef(err, "register method %s", m.Name())
		}
	}

	return nil
}
