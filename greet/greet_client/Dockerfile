FROM golang
COPY ./ /end
WORKDIR /end
RUN CGO_ENABLED=0 GOOS=linux go build -o client greet/greet_client/*.go
ENTRYPOINT /end/client
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /end/
COPY --from=0 /end/client .
ENTRYPOINT ["./client"]
EXPOSE "50051"