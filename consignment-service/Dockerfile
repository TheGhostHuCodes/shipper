# We use the official golang image, which contains all the correct tools and
# libraries.
FROM golang:1.9.0 as builder

# Set our workdir to our current service in the gopath.
WORKDIR /go/src/github.com/TheGhostHuCodes/shipper/consignment-service

# Copy the current code into our workdir.
COPY . .

# Create a dep project, and run `ensure`, which will pull in all of the
# dependencies within this directory.
RUN go get -u github.com/golang/dep/cmd/dep
RUN dep init && dep ensure

# Build the binary with flags to allow it to run in Alpine.
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo .

FROM alpine:latest

RUN apk --no-cache add ca-certificates

RUN mkdir /app
WORKDIR /app

# Instead of copying the binary from our host machine, we pull the binary from
# the container named `builder`, within this build context. This reaches into
# our previous image, finds the binary we built, and pulls it into this
# container.
COPY --from=builder /go/src/github.com/TheGhostHuCodes/shipper/consignment-service .

CMD ["./consignment-service"]