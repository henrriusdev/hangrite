ARG GO_VERSION=1
FROM golang:${GO_VERSION}-bookworm as builder

WORKDIR /usr/src/app
COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY . .
RUN go build -v -o /run-app .

FROM debian:bookworm

# Create necessary directories
RUN mkdir -p /app/data

# Copy the binary and assets
COPY --from=builder /run-app /usr/local/bin/
COPY assets /app/assets

# Set working directory
WORKDIR /app

# Expose the port
EXPOSE 8080

CMD ["run-app"]
