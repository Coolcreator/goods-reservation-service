test:
	go test -v ./...

test_unit:
	go test -v -short ./...

run_postgres:
	docker run --name postgres -p 5432:5432 \
	-e POSTGRES_USER=${POSTGRES_USER} \
	-e POSTGRES_PASSWORD=${POSTGRES_PASSWORD} \
	-d postgres:15.2-alpine

rm_postgres:
	docker stop postgres && docker rm postgres

create_database:
	docker exec -it postgres createdb --username=admin --owner=admin ${POSTGRES_DB}

drop_database:
	docker exec -it postgres dropdb ${POSTGRES_DB}

migrate_up:
	migrate -path migration/sql -database ${POSTGRES_URI} -verbose up

migrate_down:
	migrate -path migration/sql -database ${POSTGRES_URI} -verbose down
