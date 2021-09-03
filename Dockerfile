FROM golang:1.17 AS build

WORKDIR /workdir

COPY . .

RUN CGO_ENABLED=0 \
    go build -mod=vendor -o app .

ENTRYPOINT [ "./app" ]