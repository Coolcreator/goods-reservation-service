package storage

const (
	goodsAmountInWarehouse = `SELECT name, amount
		FROM goods_to_warehouses INNER JOIN goods
		ON goods_to_warehouses.good_code = goods.code
		WHERE warehouse_id = $1`

	updateAmountOfGoodInWarehouse = `UPDATE goods_to_warehouses
		SET amount = amount + $3
		WHERE warehouse_id = $1 AND good_code = $2`

	updateAmountOfGoodInReservations = `UPDATE reservations
		SET amount = reservations.amount + $3
		FROM goods_to_warehouses
		WHERE reservations.good_to_warehouse_id = goods_to_warehouses.id
		AND warehouse_id = $1 AND good_code = $2`

	warehouseAvailability = `SELECT available
		FROM warehouses
		WHERE id = $1`
)
