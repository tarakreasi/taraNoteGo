# Build Stage
FROM golang:1.23-alpine AS builder

WORKDIR /app

# Install build dependencies
RUN apk add --no-cache gcc musl-dev

# Copy modules manifests
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build binary
RUN go build -ldflags="-s -w" -o server cmd/server/main.go
RUN go build -ldflags="-s -w" -o migrate cmd/migrate/main.go

# Final Stage
FROM alpine:latest

WORKDIR /app

# Copy binaries
COPY --from=builder /app/server .
COPY --from=builder /app/migrate .

# Copy resources (templates, assets)
COPY --from=builder /app/views ./views
COPY --from=builder /app/resources ./resources
COPY --from=builder /app/public ./public
# Copy default database if exists (or rely on volume in prod)
# COPY --from=builder /app/database/database.sqlite ./database/

# Expose port
EXPOSE 3000

# Environment variables
ENV PORT=3000
ENV DATABASE_URL=database/database.sqlite

# Run server
CMD ["./server"]
