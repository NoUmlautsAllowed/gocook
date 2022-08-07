# build standalone docker container

#install certs
FROM ubuntu:latest as ubuntu
RUN apt-get update
RUN apt-get install ca-certificates -y
RUN update-ca-certificates

# Start from the latest golang base image
FROM golang:latest as golangbuilder

COPY . /go/src/gocook
WORKDIR /go/src/gocook
RUN go mod download
RUN go build  -a -tags netgo -v  -ldflags '-w -extldflags "-static"' -o /go/bin ./cmd/server

FROM scratch

# copy cert files
WORKDIR /etc/ssl/certs
COPY --from=ubuntu /etc/ssl/certs .

WORKDIR /
COPY --from=golangbuilder /go/bin .
COPY --from=golangbuilder /go/src/gocook/templates templates/

EXPOSE 8080
CMD ["./server"]