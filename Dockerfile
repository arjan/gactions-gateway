FROM golang:onbuild AS build

FROM golang:latest
RUN mkdir /app
ADD . /app/
WORKDIR /app
RUN go build -o main .
CMD ["/app/main"]

FROM scratch
COPY gactions /
COPY --from=build /app/main /

ENTRYPOINT ["/main"]
