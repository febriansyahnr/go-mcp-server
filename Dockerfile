# syntax = docker/dockerfile:1
FROM golang:1.23.3-alpine3.20 AS builder

ARG GITHUB_TOKEN
ENV GOCACHE=/go/pkg/mod

# Set workdir
WORKDIR /app

RUN apk update --no-cache && \
    apk add --no-cache git openssh

# Configure Git to use GITHUB_TOKEN for any GitHub repository
RUN echo "machine github.com login DevHarsya password ${GITHUB_TOKEN}" > ~/.netrc && \
    #git config --global url."https://${GITHUB_TOKEN}@github.com/".insteadOf "https://github.com/" && \
    export GOPRIVATE=github.com/paper-indonesia/*

# RUN go env -w GOMODCACHE=/root/.cache/go-build

COPY go.mod go.sum ./

# Cache Go modules
# RUN --mount=type=cache,target=/go/pkg/mod \
#     --mount=type=cache,target=/root/.cache/go-build \
RUN go mod tidy

# COPY all files from repo
COPY . .

# Update and install git
# RUN --mount=type=cache,target=/go/pkg/mod \
#     --mount=type=cache,target=/root/.cache/go-build \
#     GO111MODULE=on CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -installsuffix cgo -o /app/snap-core .
RUN GO111MODULE=on CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -installsuffix cgo -o /app/mcp-server .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -gcflags "all=-N -l" -a -installsuffix cgo -o /app/mcp-server_debug .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go install -ldflags "-s -w -extldflags '-static'" github.com/go-delve/delve/cmd/dlv@latest

# Use a smaller base image for the final image
FROM alpine:3.20
# Set workdir
WORKDIR /app
# Copy from builder
COPY --from=builder /app/mcp-server /app/mcp-server
COPY --from=builder /app/docs /app/docs
COPY --from=builder /app/mcp-server_debug /app/mcp-server_debug
COPY --from=builder /go/bin/dlv /app/dlv

# Install tzdata and update alpine
RUN apk update --no-cache && \
    apk add --no-cache bash libstdc++ libx11 libxrender libxext ca-certificates

# Set permission
RUN chmod +x /app/mcp-server

# Expose application
EXPOSE 3000

# Set CMD
CMD [ "/bin/sh", "-c", "/app/mcp-server $mode --config /app/env/.config.yaml --secret /app/env/.secret.yaml" ]