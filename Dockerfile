FROM golang:1.21 AS builder

WORKDIR /app

COPY go.mod  ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o flight-calculator ./cmd/flight-calculator/main.go

FROM alpine:3.14

WORKDIR /root/

COPY --from=builder /app/flight-calculator .

EXPOSE 8080

CMD ["./flight-calculator"]
