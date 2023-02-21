INSERT INTO goods(code, name) VALUES
('uniqueGoodCode01', 'хлеб'),
('uniqueGoodCode02', 'вода'),
('uniqueGoodCode03', 'молоко'),
('uniqueGoodCode04', 'водка'),
('uniqueGoodCode05', 'макароны'),
('uniqueGoodCode06', 'филе индейки'),
('uniqueGoodCode07', 'сендвич'),
('uniqueGoodCode08', 'яйца'),
('uniqueGoodCode09', 'чипсы'),
('uniqueGoodCode10', 'яблоки');

INSERT INTO warehouses(id, name, available) VALUES
(1, 'Warehouse1', true),
(2, 'Warehouse2', true),
(3, 'Warehouse3', true),
(4, 'Warehouse4', false),
(5, 'Warehouse5', true);

INSERT INTO goods_to_warehouses(good_code, warehouse_id, amount) VALUES
('uniqueGoodCode01', 1, 10),
('uniqueGoodCode01', 2, 10),
('uniqueGoodCode01', 3, 10),
('uniqueGoodCode01', 4, 10),
('uniqueGoodCode01', 5, 10),
('uniqueGoodCode02', 1, 10),
('uniqueGoodCode02', 2, 10),
('uniqueGoodCode02', 3, 10),
('uniqueGoodCode02', 4, 10),
('uniqueGoodCode02', 5, 10),
('uniqueGoodCode03', 1, 5),
('uniqueGoodCode04', 2, 5),
('uniqueGoodCode05', 3, 5),
('uniqueGoodCode06', 4, 5),
('uniqueGoodCode07', 5, 5),
('uniqueGoodCode08', 1, 5),
('uniqueGoodCode09', 2, 5),
('uniqueGoodCode03', 3, 10),
('uniqueGoodCode04', 4, 10),
('uniqueGoodCode05', 5, 10);