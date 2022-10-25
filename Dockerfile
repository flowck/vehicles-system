FROM golang as builder
WORKDIR /app
COPY . ./
RUN go build -o vehicles_system_server ./cmd/server

FROM alpine
WORKDIR /app
COPY --from=builder /app/vehicles_system_server ./
ENTRYPOINT ["./vehicles_system_server"]
