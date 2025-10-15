FROM alpine:latest

COPY build/ main/

WORKDIR /main
CMD ["./app"]
