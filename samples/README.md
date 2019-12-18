<https://dev.to/ivan/go-build-a-minimal-docker-image-in-just-three-steps-514i>

```
FROM golang:1.13.0-stretch AS builder

ENV GO111MODULE=on \
    CGO_ENABLED=1

WORKDIR /build

# Let's cache modules retrieval - those don't change so often
COPY go.mod .
COPY go.sum .
RUN go mod download

# Copy the code necessary to build the application
# You may want to change this to copy only what you actually need.
COPY . .

# Build the application
RUN go build ./cmd/my-awesome-go-program

# Let's create a /dist folder containing just the files necessary for runtime.
# Later, it will be copied as the / (root) of the output image.
WORKDIR /dist
RUN cp /build/my-awesome-go-program ./my-awesome-go-program

# Optional: in case your application uses dynamic linking (often the case with CGO), 
# this will collect dependent libraries so they're later copied to the final image
# NOTE: make sure you honor the license terms of the libraries you copy and distribute
RUN ldd my-awesome-go-program | tr -s '[:blank:]' '\n' | grep '^/' | \
    xargs -I % sh -c 'mkdir -p $(dirname ./%); cp % ./%;'
RUN mkdir -p lib64 && cp /lib64/ld-linux-x86-64.so.2 lib64/

# Copy or create other directories/files your app needs during runtime.
# E.g. this example uses /data as a working directory that would probably
#      be bound to a perstistent dir when running the container normally
RUN mkdir /data

# Create the minimal runtime image
FROM scratch

COPY --chown=0:0 --from=builder /dist /

# Set up the app to run as a non-root user inside the /data folder
# User ID 65534 is usually user 'nobody'. 
# The executor of this image should still specify a user during setup.
COPY --chown=65534:0 --from=builder /data /data
USER 65534
WORKDIR /data

ENTRYPOINT ["/my-awesome-go-program"]
```

<https://www.cloudreach.com/en/insights/blog/containerize-this-how-to-build-golang-dockerfiles/>

```
FROM golang:alpine as builder
RUN mkdir /build 
ADD . /build/
WORKDIR /build 
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o main .
FROM scratch
COPY --from=builder /build/main /app/
WORKDIR /app
CMD ["./main"]
```

<https://blog.container-solutions.com/faster-builds-in-docker-with-go-1-11>
```

# Base build image
FROM golang:1.11-alpine AS build_base
# Install some dependencies needed to build the project
RUN apk add bash ca-certificates git gcc g++ libc-dev
WORKDIR /go/src/github.com/creativesoftwarefdn/weaviate

# Force the go compiler to use modules 
ENV GO111MODULE=on

# We want to populate the module cache based on the go.{mod,sum} files. 
COPY go.mod .
COPY go.sum .

#This is the ‘magic’ step that will download all the dependencies that are specified in 
# the go.mod and go.sum file.

# Because of how the layer caching system works in Docker, the go mod download 
# command will _ only_ be re-run when the go.mod or go.sum file change 
# (or when we add another docker instruction this line) 
RUN go mod download

# This image builds the weavaite server
FROM build_base AS server_builder
# Here we copy the rest of the source code
COPY . .
# And compile the project
RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go install -a -tags netgo -ldflags '-w -extldflags "-static"' ./cmd/weaviate-server

#In this last stage, we start from a fresh Alpine image, to reduce the image size and not ship the Go compiler in our production artifacts.
FROM alpine AS weaviate
# We add the certificates to be able to verify remote weaviate instances
RUN apk add ca-certificates
# Finally we copy the statically compiled Go binary.
COPY --from=server_builder /go/bin/weaviate-server /bin/weaviate
ENTRYPOINT ["/bin/weaviate"]
	```
  
  <https://blog.golang.org/docker>
  ```
 Start from a Debian image with the latest version of Go installed
 and a workspace (GOPATH) configured at /go.
FROM golang

 Copy the local package files to the container's workspace.
ADD . /go/src/github.com/golang/example/outyet

 Build the outyet command inside the container.
 (You may fetch or manage dependencies here,
 either manually or with a tool like "godep".)
RUN go install github.com/golang/example/outyet

 Run the outyet command by default when the container starts.
ENTRYPOINT /go/bin/outyet

 Document that the service listens on port 8080.
EXPOSE 8080
```

<https://ops.tips/blog/dockerfile-golang/>
```
 Retrieve the `golang:alpine` image to provide us the 
 necessary Golang tooling for building Go binaries.

 Here I retrieve the `alpine`-based just for the 
 convenience of using a tiny image.
FROM golang:alpine

 Add the `main` file that is really the only Golang 
 file under the root directory that matters for the 
 build 
ADD ./main.go /go/src/github.com/cirocosta/l7/main.go

 Add all the files from the packages that I own
ADD ./lib /go/src/github.com/cirocosta/l7/lib

 Add vendor dependencies (committed or not)
 I typically commit the vendor dependencies as it
 makes the final build more reproducible and less
 dependant on dependency managers.
ADD ./vendor /go/src/github.com/cirocosta/l7/vendor

 0.    Set some shell flags like `-e` to abort the 
       execution in case of any failure (useful if we 
       have many ';' commands) and also `-x` to print to 
       stderr each command already expanded.
 1.    Get into the directory with the golang source code
 2.    Perform the go build with some flags to make our
       build produce a static binary (CGO_ENABLED=0 and 
       the `netgo` tag).
 3.    copy the final binary to a suitable location that
       is easy to reference in the next stage
RUN set -ex && \
  cd /go/src/github.com/cirocosta/l7 && \       
  CGO_ENABLED=0 go build \
        -tags netgo \
        -v -a \
        -ldflags '-extldflags "-static"' && \
  mv ./l7 /usr/bin/l7

 Set the binary as the entrypoint of the container
ENTRYPOINT [ "l7" ]
```


<https://blog.hasura.io/the-ultimate-guide-to-writing-dockerfiles-for-go-web-apps-336efad7012c/>

```
FROM golang:1.8.5-jessie as builder
 install glide
RUN go get github.com/Masterminds/glide
 setup the working directory
WORKDIR /go/src/app
ADD glide.yaml glide.yaml
ADD glide.lock glide.lock
 install dependencies
RUN glide install
 add source code
ADD src src
 build the source
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main src/main.go

 use a minimal alpine image
FROM alpine:3.7
 add ca-certificates in case you need them
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
 set working directory
WORKDIR /root
 copy the binary from builder
COPY --from=builder /go/src/app/main .
 run the binary
CMD ["./main"]
```
