# Build the binary
FROM golang:1.18.3 as builder
WORKDIR YOUR_WORKING_DIRECTORY
COPY go.mod go.mod
COPY go.sum go.sum
COPY plugins plugins
COPY main.go main.go
RUN go build -o kube-scheduler ./main.go

# Copy the binary into a thin image
FROM alpine:3.13.5
RUN apk add gcompat
WORKDIR /root
COPY --from=builder /YOUR_WORKING_DIRECTORY/kube-scheduler /usr/local/bin/
