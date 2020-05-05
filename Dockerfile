FROM golang:1.14-alpine AS builder
RUN apk add --no-cache git
COPY . tmp/simpleMetricsServiceInGo
WORKDIR tmp/simpleMetricsServiceInGo
RUN go build -o /bin/metricsServer cmd/metricsServer/main.go

FROM alpine:3.11 AS final
RUN apk add --no-cache
COPY --from=builder /bin/metricsServer /bin/metricsServer
CMD ["/bin/metricsServer"]
