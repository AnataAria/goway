# syntax=docker/dockerfile:1

FROM golang:1.21-alpine AS builder

LABEL org.opencontainers.image.source="https://github.com/AnataAria/goway"
LABEL org.opencontainers.image.licenses="MIT"

RUN apk add --no-cache ca-certificates git build-base \
    && adduser -D -H -u 1000 goway

WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download

COPY . .
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
RUN --mount=type=cache,target=/root/.cache/go-build \
    --mount=type=cache,target=/go/pkg/mod \
    go build -trimpath -ldflags="-s -w" -o /goway ./cmd/worker
RUN chmod 755 /goway && chown 1000:1000 /goway

FROM gcr.io/distroless/static:nonroot
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /goway /goway
EXPOSE 8080
USER 1000
ENTRYPOINT ["/goway"]
