# Build stage
FROM golang:latest AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN chmod +x entrypoint.sh

RUN CGO_ENABLED=0 GOOS=linux go build -o fireback cmd/fireback/main.go


# Runtime stage
FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/fireback .
COPY --from=builder /app/entrypoint.sh .

RUN chmod +x entrypoint.sh

EXPOSE 4500

ENTRYPOINT ["./entrypoint.sh"]