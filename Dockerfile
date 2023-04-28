# Start from the official Go image
FROM golang:1.17-alpine AS builder

# Set the working directory
WORKDIR /app

# Copy the Go modules and build cache
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Create a minimal image for deployment
FROM alpine:latest

# Set the working directory
WORKDIR /app

# Copy the binary from the builder image
COPY --from=builder /app/main .

# Expose the port on which the API will run
EXPOSE 9292

# Start the application
CMD ["./main"]

