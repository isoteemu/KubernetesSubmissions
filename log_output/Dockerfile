## Multi-stage build: compile Go binary, then ship minimal image.
FROM golang:1.24-alpine AS builder

WORKDIR /src

COPY go.mod ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
	go build -trimpath -ldflags="-s -w" -o /out/app ./main.go

FROM gcr.io/distroless/static:nonroot

WORKDIR /app
COPY --from=builder /out/app /app/app

USER nonroot:nonroot
ENTRYPOINT ["/app/app"]
