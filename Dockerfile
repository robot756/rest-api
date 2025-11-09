FROM golang:1.24-alpine AS builder

WORKDIR /usr/local/src

RUN apk --no-cache add bash git make gcc gettext musl-dev

# dependencies
COPY ["go.mod", "go.sum", "./"]
RUN go mod download

# build
COPY . .
RUN go build -o ./bin/url cmd/url-shortener/main.go

FROM alpine AS runner

COPY --from=builder /usr/local/src/bin/url /
COPY config/local.yaml /config/local.yaml
CMD ["/url"]