# Start from golang v1.13.4 base image to have access to go modules
FROM golang:1.16.6

# create a working directory
RUN mkdir -p /app
WORKDIR /app

# copy source from the host to the working directory inside
# the container
COPY . /app

# Fetch dependencies on separate layer as they are less likely to
# change on every build and will therefore be cached for speeding
# up the next build
RUN go mod download

# Build api binary within container
RUN go build ./...

# Set binary as entrypoint
ENTRYPOINT /go/bin/api

# This container exposes port 7777 to the outside world
EXPOSE 7777