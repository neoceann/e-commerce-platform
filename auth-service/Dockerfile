FROM golang:latest AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o auth-service ./cmd/auth

FROM alpine:latest

RUN apk --no-cache add ca-certificates

COPY --from=builder /app/auth-service /auth-service

RUN GRPC_HEALTH_PROBE_VERSION=v0.4.24 && \
    wget -qO/bin/grpc_health_probe \
    https://github.com/grpc-ecosystem/grpc-health-probe/releases/download/${GRPC_HEALTH_PROBE_VERSION}/grpc_health_probe-linux-amd64 && \
    chmod +x /bin/grpc_health_probe

HEALTHCHECK --interval=10s --timeout=5s --start-period=5s --retries=3 \
    CMD /bin/grpc_health_probe -addr=localhost:50051 || exit 1

CMD ["/auth-service"]