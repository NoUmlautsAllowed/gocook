[![Go](https://github.com/NoUmlautsAllowed/gocook/actions/workflows/go.yml/badge.svg)](https://github.com/NoUmlautsAllowed/gocook/actions/workflows/go.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/NoUmlautsAllowed/gocook)](https://goreportcard.com/report/github.com/NoUmlautsAllowed/gocook)
[![codecov](https://codecov.io/gh/NoUmlautsAllowed/gocook/branch/main/graph/badge.svg?token=OO2AKXBRKU)](https://codecov.io/gh/NoUmlautsAllowed/gocook)

# :cook: GoCook

An alternative frontend to Chefkoch with a focus on privacy. Static CSS and 
HTML only. Built with [Go](https://go.dev/)
and [Bulma](https://bulma.io/).

## What is this?

This is an alternative frontend to Chefkoch. The generated sites are static, 
i.e. containing no JavaScript that runs on the client side.

Currently, only the API traffic is proxied through the server. All CDN 
requests needed to get images are made locally by the browser.
The CDN requests to display images will be proxied through the server in 
the future to ensure a maximum of privacy.

## Build and Deployment

The easiest way to deploy this service is the standalone docker image. 
Alternatively, building the application from source is possible too.

### Docker

The official docker image is available at docker hub in the repository 
[`noumlautsallowed/gocook`](https://hub.docker.com/r/noumlautsallowed/gocook).

Publish the port `8080` of the container locally:

```shell
docker run -d --name gocook -p 127.0.0.1:8080:8080 noumlautsallowed/gocook:latest
```

### Build from source

If you want to build this project from source, checkout the repository 
locally and run the server.

Currently, only Go needs to be installed.

```shell
git clone git@github.com:NoUmlautsAllowed/gocook.git && cd gocook
npm i
npm run build
go build ./cmd/server
./server
```
