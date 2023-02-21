# Goods reservation service

## Инструкция по запуску

```
docker compose up
```

ИЛИ (запуск сервиса локально с postgres в контейнере)

```
export $(grep -v '^#' .env | xargs)
make run_postgres
make create_database
make migrate_up
export POSTGRES_HOST=localhost
go run cmd/app/main.go
```

Swagger не подходит для JSONRP, так как оперирует с глаголами и путями.
Добавлены фиктивные пути для демонстрации в файле docs/swagger.yaml,
можно посмотреть в Gitlab.

### Примеры запросов

Запросы направляются на эндпоинт http://localhost:8080/rpc
Код ответа 200 даже при ошибке

#### Amount

Request

```
{
	"jsonrpc": "2.0",
	"method": "amount",
	"params": {
		"warehouse_id": 1
	}
}
```

Responce

```
{
	"jsonrpc": "2.0",
	"result": {
		"warehouse_id": 1,
		"amount": {
			"вода": 10,
			"молоко": 5,
			"хлеб": 10,
			"яйца": 5
		}
	}
}
```

### Reserve

Request

```
{
	"jsonrpc": "2.0",
	"method": "reserve",
	"params": {
		"warehouse_id": 1,
		"codes": [
			"uniqueGoodCode01",
			"uniqueGoodCode02"
		]
	}
}
```

Responce

```
{
	"jsonrpc": "2.0",
	"result": {
		"warehouse_id": 1,
		"reserved": {
			"вода": 10,
			"молоко": 5,
			"хлеб": 10,
			"яйца": 5
		}
	}
}
```

### Release

Request

```
{
	"jsonrpc": "2.0",
	"method": "release",
	"params": {
		"warehouse_id": 1,
		"codes": [
			"uniqueGoodCode01"
		]
	}
}
```

Responce

```
{
	"jsonrpc": "2.0",
	"result": {
		"warehouse_id": 1,
		"released": {
			"uniqueGoodCode01": 1
		}
	}
}
```

Логика требует полного резервирования, те нельзя частично резервировать товары
Будет возвращена ошибка оut of stock/out of reserve с кодом первого незарезервированного/
неосвобожденного товара

```
{
	"jsonrpc": "2.0",
	"error": {
		"code": 102,
		"message": "good uniqueGoodCode01 is out of reservations in warehouse 1"
	}
}
```

На коды товаров наложено ограничение в 16 символов. Любой код в запросе, который
не равен 16 символам вызывает ошибку

```
{
	"jsonrpc": "2.0",
	"error": {
		"code": 102,
		"message": "invalid good code '12345678910': must be 16 chars long"
	}
}
```