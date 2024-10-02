FROM golang:1.23

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY cmd/**/*.go ./

# Use multi-stage build for smaller final image
RUN CGO_ENABLED=0 GOOS=linux go build -v -o /go-template

# Start a new stage with a minimal base image
FROM alpine:latest  

# Copy the binary from the build stage
COPY --from=0 /go-template /go-template

EXPOSE 8080

# Run as non-root user for better security
RUN adduser -D appuser
USER appuser

CMD ["/go-template"]