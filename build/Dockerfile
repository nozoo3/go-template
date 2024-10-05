FROM golang:1.23

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

EXPOSE 8080

# Run as non-root user for better security
RUN useradd -m appuser
USER appuser