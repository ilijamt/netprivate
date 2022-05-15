FROM golang:1.18-alpine as builder

ENV CGO_ENABLED 0

RUN mkdir -p /app
WORKDIR /app

COPY . /app

RUN go mod tidy
RUN go generate ./...
RUN go mod download

RUN go build -o /app/is-private /app/cmd/is-private/main.go

FROM scratch

COPY --from=builder /app/is-private /app/is-private
ENTRYPOINT ["/app/is-private"]
