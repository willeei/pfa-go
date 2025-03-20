FROM golang:1.24

# Set the Current Working Directory inside the container
WORKDIR /app

ENTRYPOINT [ "tail", "-f", "/dev/null" ]