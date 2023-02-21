package domain

import (
	"github.com/1nkh3art1/goods-reservation-service/internal/cerror"
)

type (
	WarehouseID int

	// Request представляет собой запрос на
	// резервирование/освобождение товаров
	// с уникальными кодами на складе
	Request struct {
		WarehouseID WarehouseID
		Codes       []string
	}
)

const (
	warehouseIDLength = 13
	goodCodeLength    = 16
)

func (w WarehouseID) Validate() error {
	if w <= 0 {
		return cerror.BadRequest.Newf("invalid warehouse ID %d: must be positive", w)
	}

	return nil
}

// Validate валидирует параметры запроса
func (r Request) Validate() error {
	if err := r.WarehouseID.Validate(); err != nil {
		return err
	}

	if len(r.Codes) == 0 {
		return cerror.BadRequest.New("empty codes list")
	}

	for _, code := range r.Codes {
		if len(code) != goodCodeLength {
			return cerror.BadRequest.Newf(
				"invalid good code '%s': must be %d chars long",
				code,
				goodCodeLength,
			)
		}
	}

	return nil
}
