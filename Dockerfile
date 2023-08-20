FROM golang:1.18

ENV GOPATH=/

# Copy the local package files to the container's workspace
COPY . .

# Build the Go application
RUN go mod download
RUN go build -o service ./cmd/server/main.go

# Command to run the application
CMD ["./service"]
