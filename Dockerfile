# build standalone docker container

#install certs
FROM ubuntu:latest AS ubuntu
RUN apt-get update
RUN apt-get install ca-certificates -y
RUN update-ca-certificates

# Build node modules
FROM node:22 AS nodebuilder
COPY . /home/node/app
WORKDIR /home/node/app
RUN npm i
RUN npm run build

# Start from the latest golang base image
FROM golang:1.24 AS golangbuilder

COPY . /go/src/gocook
WORKDIR /go/src/gocook


# copy static files
COPY --from=nodebuilder /home/node/app/web/static/css/ web/static/css/
COPY --from=nodebuilder /home/node/app/web/static/fonts/ web/static/fonts/

RUN go mod download
RUN go build  -a -tags netgo -v  -ldflags '-w -extldflags "-static"' -o /go/bin ./cmd/server

FROM scratch

# copy cert files
WORKDIR /etc/ssl/certs
COPY --from=ubuntu /etc/ssl/certs .

# target gocook directory in image
WORKDIR /gocook

# copy go binary and templates
COPY --from=golangbuilder /go/bin .

ENV GIN_MODE=release
EXPOSE 8080
CMD ["./server"]
