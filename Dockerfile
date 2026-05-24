FROM golang:1.25 AS builder
WORKDIR /app

COPY go.mod .
RUN go mod download
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o ./photo-service-executor/photo ./cmd/main.go

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/photo-service-executor/photo .
CMD [ "./photo" ]