FROM golang:1.25

WORKDIR /app

# Create user with proper UID/GID and set up directories
RUN useradd -m -u 1000 appuser && \
    chown -R appuser:appuser /app && \
    chown -R appuser:appuser /go

# Copy go files and download dependencies as root first
COPY go.mod go.sum ./
RUN go mod download

EXPOSE 8080

# Switch to non-root user
USER appuser