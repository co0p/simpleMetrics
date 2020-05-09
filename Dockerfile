FROM golang:1.14-alpine AS builder
RUN apk add --no-cache git
COPY . tmp/simpleMetrics
WORKDIR tmp/simpleMetrics
RUN go build -o /bin/exampleServer cmd/exampleServer/main.go

FROM alpine:3.11 AS final
RUN apk add --no-cache
COPY --from=builder /bin/exampleServer /bin/exampleServer
CMD ["/bin/exampleServer"]
