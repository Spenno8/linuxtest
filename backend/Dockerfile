# ---------- Build Stage ----------
FROM golang:1.25-alpine AS builder

WORKDIR /app

# Copy go mod files first (better caching)
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app

# ---------- Run Stage ----------
FROM alpine:latest

WORKDIR /root/

# Copy compiled binary from builder stage
COPY --from=builder /app/app .

# Expose port
EXPOSE 8080

# Run the app
CMD ["./app"]
