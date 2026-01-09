# Multi-stage build: compile Go binary, then ship minimal image.
FROM golang:1.24-alpine AS builder

WORKDIR /src

COPY go.mod go.sum ./
RUN --mount=type=cache,target=/go/pkg/mod go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -trimpath -ldflags="-s -w" -o /out/todoapp .

FROM gcr.io/distroless/static:nonroot

WORKDIR /app
COPY --from=builder /out/todoapp /app/todoapp

USER nonroot:nonroot
ENTRYPOINT ["/app/todoapp"]
