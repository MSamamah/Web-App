# Use the official Golang base image to build the Go application
FROM golang:1.21 as builder

# Set the working directory inside the container
WORKDIR /app

# Copy go module files and download dependencies
COPY go.mod ./
RUN go mod download

# Copy the source code
COPY . .

# Build the Go binary
RUN go build -o main .

#######################################################
# Use a minimal distroless image to reduce final image size
FROM gcr.io/distroless/base

# Copy the compiled binary and static files from the builder stage
COPY --from=builder /app/main .
COPY --from=builder /app/static ./static

# Expose the application port
EXPOSE 8080

# Run the application
CMD ["./main"]
