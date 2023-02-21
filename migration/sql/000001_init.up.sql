CREATE TABLE warehouses (
    id bigserial PRIMARY KEY,
    name varchar NOT NULL UNIQUE,
    available boolean NOT NULL
);

CREATE TABLE goods(
    code varchar(16) PRIMARY KEY,
    name varchar(255) NOT NULL,
    size varchar(255)
);

CREATE TABLE goods_to_warehouses(
    id bigserial PRIMARY KEY,
    good_code varchar(16) NOT NULL,
    warehouse_id bigint NOT NULL,
    amount bigint NOT NULL CHECK (amount >= 0)
);

CREATE TABLE reservations(
    id bigserial PRIMARY KEY,
    good_to_warehouse_id bigint NOT NULL,
    amount bigint NOT NULL CHECK (amount >= 0)
);

ALTER TABLE goods_to_warehouses ADD FOREIGN KEY (good_code) REFERENCES goods(code);
ALTER TABLE goods_to_warehouses ADD FOREIGN KEY (warehouse_id) REFERENCES warehouses(id);
ALTER TABLE reservations ADD FOREIGN KEY (good_to_warehouse_id) REFERENCES goods_to_warehouses(id);

CREATE UNIQUE INDEX good_to_warehouse_index ON goods_to_warehouses(good_code, warehouse_id);


CREATE OR REPLACE FUNCTION reservations_insert_func()
    RETURNS trigger AS
$$
BEGIN
    INSERT INTO reservations(good_to_warehouse_id, amount) VALUES(NEW.id, 0);
RETURN NEW;
END;
$$
LANGUAGE 'plpgsql';
CREATE TRIGGER reservations_insert_trigger
    AFTER INSERT
    ON "goods_to_warehouses"
    FOR EACH ROW
    EXECUTE PROCEDURE reservations_insert_func();


-- CREATE OR REPLACE FUNCTION reservations_delete_func()
--     RETURNS trigger AS
-- $$
-- BEGIN
--     INSERT INTO reservations(good_to_warehouse_id, amount) VALUES(NEW.id, 0);
-- RETURN NEW;
-- END;
-- $$
-- LANGUAGE 'plpgsql';
-- CREATE TRIGGER reservations_insert_trigger
--     AFTER INSERT
--     ON "goods_to_warehouses"
--     FOR EACH ROW
--     EXECUTE PROCEDURE reservations_insert_func();