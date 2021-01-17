DATABASE_URL := mysql://${MYSQL_USER}:${MYSQL_PASS}@tcp(${MYSQL_HOST}:${MYSQL_PORT})/${MYSQL_DATABASE}

## golang-migrate
.PHONY: create_migrations
create_migrations:
	docker run -v $(PWD)/db/migrations:/db/migrations --network host migrate/migrate create -ext sql -dir /db/migrations -seq $(FILE_NAME)

.PHONY: migrateup
migrateup:
	docker run -v $(PWD)/db/migrations:/db/migrations --network host migrate/migrate -source file://db/migrations -database "${DATABASE_URL}" up

.PHONY: migratedown
migratedown:
	docker run -v $(PWD)/db/migrations:/db/migrations --network host migrate/migrate migrate -source file://db/migrations/ -database "${DATABASE_URL}" down

.PHONY: gen
gen:
	rm -rf ./src/models/*.xo.go
	xo mysql://${MYSQL_USER}:${MYSQL_PASS}@${MYSQL_HOST}:${MYSQL_PORT}/${MYSQL_DATABASE} -o ./src/models
