FROM golang:onbuild

FROM golang:latest AS build
RUN mkdir /app
ADD . /app/
WORKDIR /app
RUN env GOOS=linux GOARCH=386 go build -o main .
CMD ["/app/main"]

FROM alpine:3.6
RUN apk add --update bash openssl ca-certificates libgcc libstdc++

COPY --from=build /app/main /
COPY gactions /bin

RUN chmod +x /main
CMD ["/main"]
