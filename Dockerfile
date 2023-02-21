FROM golang:1.20.1-alpine3.17 AS builder

ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /build

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go build -o app cmd/app/main.go

WORKDIR /dist

RUN cp /build/app .

FROM alpine:latest

COPY . .

COPY --from=builder /dist/app /

ENTRYPOINT ["/app"]