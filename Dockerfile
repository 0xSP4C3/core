FROM golang:1.17-alpine AS builder

LABEL maintainer="p3nj <p3nj@bumpto.space> (https://benji.ltd)"

# Move to working directory (/build).
WORKDIR /build

# Copy and download dependency using go mod.
COPY go.mod go.sum ./
RUN go mod download

# Copy the code into the container.
COPY . .

# Set necessary environment variables needed for our image and build the API server.
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
RUN go build -ldflags="-s -w" -o coreapi .

FROM scratch

# Copy binary and config files from /build to root folder of scratch container.
COPY --from=builder ["/build/coreapi", "/build/.env", "/build/certs/ca.crt", "/"]


# Command to run when starting the container.
ENTRYPOINT ["/coreapi"]
