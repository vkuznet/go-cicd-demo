# Build Stage
FROM golang AS builder
WORKDIR /app
COPY go.mod go.sum ./
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o hello .

# Minimal Final Stage
FROM alpine:latest
COPY --from=builder /app/hello /usr/local/bin/hello
ENTRYPOINT ["hello"]
