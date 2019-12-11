# Dockerfile References: https://docs.docker.com/engine/reference/builder/

# Start from the latest golang base image
FROM golang:latest AS builder

# Add Maintainer Info
LABEL maintainer="Sergii Gladchenko <gladseo@gmail.com>"

# Set the Current Working Directory inside the container
WORKDIR /build

# Copy the source from the current directory to the Working Directory inside the container
COPY ./main.go .

# Build the Go app
RUN CGO_ENABLED=0 go build -a -installsuffix cgo -o main .

######## Start a new stage from scratch #######
FROM scratch

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /build/main .

# Expose port 3000 to the outside world
EXPOSE 3000

# Command to run the executable
ENTRYPOINT ["./main"]
