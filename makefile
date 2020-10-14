
migrate_status:
	cd ./migrations; \
	goose postgres "$(KITCHEN_PG)" status

migrate_up:
	cd ./migrations; \
	goose postgres "$(KITCHEN_PG)" up

migrate_down:
	cd ./migrations; \
	goose postgres "$(KITCHEN_PG)" down

test:
	go test ./...
