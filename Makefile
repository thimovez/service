db:
	docker run --name some-service -p 5433:5432 -e POSTGRES_USER=user -e POSTGRES_PASSWORD=password -e POSTGRES_DB=service -d postgres
app: 
	go run cmd/app/main.go
start:
	docker compose up --build service