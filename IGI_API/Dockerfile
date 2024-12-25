# Stage 1: Build Stage
FROM golang:1.23.4 AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy Go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go application with static linking
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/main.go

# Stage 2: Run Stage
FROM alpine:latest

# Install necessary runtime dependencies
RUN apk --no-cache add ca-certificates

# Set the working directory for runtime
WORKDIR /root/

# Copy the built application from the builder stage
COPY --from=builder /app/main .

# Environment variables
ENV BASE_URL=http://localhost:8080
ENV SWAPI_BASE_URL=https://www.swapi.tech/api/
ENV SWAGGER_URL=http://localhost:8080/docs/

# Expose the application port
EXPOSE 8080

# Command to run the application
CMD ["./main"]
