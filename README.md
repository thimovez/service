# service
- 1) run db - make db
- 2) run migration -dir ./migrations postgres "postgresql://user:password@localhost:5432/service?sslmode=disable" up
- 3) run server