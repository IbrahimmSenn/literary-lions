FROM golang:1.23

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o forum cmd/main.go

EXPOSE 8080

# Create data directory for database
RUN mkdir -p /app/data

CMD ["./forum"]