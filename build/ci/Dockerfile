ARG DOCKER_PROXY_HOST_1=docker.io
FROM $DOCKER_PROXY_HOST_1/golang:1.19 as builder
WORKDIR /app
COPY go.* ./
RUN go mod download && go mod verify
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o booking ./cmd/be/main.go


FROM $DOCKER_PROXY_HOST_1/alpine:3
WORKDIR /app
COPY --from=builder /app/booking .
COPY --from=builder /app/cmd/be/migrations/* /be/migrations/
CMD ["./cmd/be", "run"]
