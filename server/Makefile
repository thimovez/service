db:
	docker run --name some-service -p 5432:5432 -e POSTGRES_USER=user -e POSTGRES_PASSWORD=password -e POSTGRES_DB=service -d postgres
start:
	docker compose up --build service