# Use an official Golang image
FROM golang:1.22.4

# Set the working directory
WORKDIR /app

# Copy all files to the container
COPY . .

# Download dependencies
RUN go mod download

# Build the application
RUN go build -o main .

# Expose the port
EXPOSE 8001

# Run the application
CMD ["./main"]
