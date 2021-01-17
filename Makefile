## golang-migrate
.PHONY: create_migrations
create_migrations:
	docker run -v $(PWD)/db/migrations:/db/migrations --network host migrate/migrate create -ext sql -dir /db/migrations -seq $(FILE_NAME)

.PHONY: migrateup
migrateup:
	docker run -v $(PWD)/db/migrations:/db/migrations --network host migrate/migrate 	migrate -source file://db/mysql/migrations/ -database '$(PATHPORT_DATABASE_URL)' up
