FROM golang:1.20-alpine AS builder

WORKDIR /usr/local/src

RUN apk --no-cache add bash git make gcc gettext musl-dev

COPY src/go.mod src/go.sum ./
RUN go mod download

# Build the application
COPY src ./
RUN go build -o ./bin/service ./cmd/app/main.go

FROM alpine AS runner

COPY --from=builder /usr/local/src/bin/service /
COPY --from=builder /usr/local/src/config/config.yml /
COPY --from=builder /usr/local/src/migrations /migrations

EXPOSE 8080

CMD [ "/service" ]
