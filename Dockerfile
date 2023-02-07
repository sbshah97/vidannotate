FROM golang:latest

# Set working directory
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN go build -o main .

# Set environment variables
ENV API_KEY=valid_api_key
ENV JWT_TOKEN=valid_jwt

# Expose port 8000
EXPOSE 8000

# Run the application
CMD ["./main"]
