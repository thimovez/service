db:
	docker run --name some-service -p 5432:5432 -e POSTGRES_USER=user -e POSTGRES_PASSWORD=password -e POSTGRES_DB=service -d postgres
#migrate_up:
#	goose -dir ./migrations postgres "postgresql://user-repo:password@localhost:5432/service?sslmode=disable" up
#migrate_down:
#	goose -dir ./migrations postgres "postgresql://user-repo:password@localhost:5432/service?sslmode=disable" down
