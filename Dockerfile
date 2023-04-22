# Download golang image
FROM golang:latest

# Set necessary environmen variables needed for out image
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Create to working directory /app
RUN mkdir /app

# Move to working directory /app
WORKDIR /app

# Create directory /build for compiler
RUN mkdir /build

# Copy the code into container
COPY . .

# Download required modules
RUN go mod download

# Create directory logs
RUN mkdir /etc/push

# Build the application
RUN go build -o ./dist/main .

# Expose in external port 8891
EXPOSE 8891

# Command to run
ENTRYPOINT ["/app/dist/main"]