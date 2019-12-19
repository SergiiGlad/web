# Dockerfile References: https://docs.docker.com/engine/reference/builder/
######## from scratch #######
FROM scratch

# Copy the Pre-built binary file from the previous stage
# Set up the app to run as a non-user
# User ID 65534 is usually user 'nobody'
COPY --chown=65534:0 ./main ./main
USER 65534

# Expose port 3000 to the outside world
EXPOSE 3030

# Command to run the executable
ENTRYPOINT ["./main"]
