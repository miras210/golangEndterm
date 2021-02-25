FROM golang
COPY ./ /end
WORKDIR /end
RUN CGO_ENABLED=0 GOOS=linux go build -o server greet/greet_server/*.go
ENTRYPOINT /end/server
# Build final image
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /end/
COPY --from=0 /end/server .
ENTRYPOINT ["./server"]
EXPOSE "50051"