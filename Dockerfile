FROM alpine:latest

RUN apk --no-cache add ca-certificates
WORKDIR /var/lib/app

CMD ["./app"]