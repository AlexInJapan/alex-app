# syntax=docker/dockerfile:1

ARG GO_VERSION=1.23
ARG ALPINE_VERSION=3.22

FROM golang:${GO_VERSION}-alpine AS base
ARG VERSION="0.0.0+unknown"
WORKDIR /src

RUN apk --no-cache add ca-certificates

FROM base AS build
RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=bind,source=go.mod,target=go.mod \
    --mount=type=bind,source=go.sum,target=go.sum \
    go mod download

RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=bind,target=. \
    CGO_ENABLED=0 GOOS=linux go build \
    -v \
    -ldflags="-X main.Version=${VERSION} -w -s" \
    -o /bin/app \
    ./cmd/server

FROM alpine:${ALPINE_VERSION} AS image 

RUN apk --no-cache add ca-certificates tzdata && \
    addgroup -g 1001 -S appgroup && \
    adduser -u 1001 -S appuser -G appgroup

COPY --from=build /bin/app /bin/app
COPY app.env app.env
COPY db/migrations db/migrations

USER appuser
EXPOSE 8080
ENTRYPOINT ["/bin/app"]

