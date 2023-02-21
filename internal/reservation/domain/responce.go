package domain

// Responce представляет собой количество
// зарезервированных/освобожденных товаров
// с уникальными кодами на складе
// (код товара -> количество)
type Responce struct {
	Amount map[string]int
}
