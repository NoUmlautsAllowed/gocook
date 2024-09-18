[![status-badge](https://ci.codeberg.org/api/badges/13698/status.svg)](https://ci.codeberg.org/repos/13698)
[![Go](https://github.com/NoUmlautsAllowed/gocook/actions/workflows/go.yml/badge.svg)](https://github.com/NoUmlautsAllowed/gocook/actions/workflows/go.yml)
[![Go Report Card](https://goreportcard.com/badge/codeberg.org/NoUmlautsAllowed/gocook)](https://goreportcard.com/report/codeberg.org/NoUmlautsAllowed/gocook)
[![codecov](https://codecov.io/gh/NoUmlautsAllowed/gocook/branch/main/graph/badge.svg?token=OO2AKXBRKU)](https://codecov.io/gh/NoUmlautsAllowed/gocook)

# :cook: GoCook

An alternative frontend to Chefkoch with a focus on privacy. Static CSS and 
HTML only. Built with [Go](https://go.dev/)
and [Bulma](https://bulma.io/).

## What is this?

This is an alternative frontend to Chefkoch. The generated sites are static, 
i.e. containing no JavaScript that runs on the client side.

All API and CDN traffic is proxied through GoCook, there is no communication
with any other hosts than GoCook itself.

The repository is hosted on [Codeberg](https://codeberg.org/NoUmlautsAllowed/gocook), 
a mirror is maintained on [GitHub](https://github.com/NoUmlautsAllowed/gocook).
Please open any new issues or pull requests on Codeberg :pray:

## Hosted instances

- <https://cook.adminforge.de/>
- <https://gocook.lunar.icu/>

## Build and Deployment

The easiest way to deploy this service is the standalone docker image. 
Alternatively, building the application from source is possible too.

API and CDN requests are not cached for now. This may be a feature to be
added in the future.

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

Currently, only Go and NPM needs to be installed.

```shell
git clone https://codeberg.org/NoUmlautsAllowed/gocook.git && cd gocook
npm i
npm run build
go build ./cmd/server
./server
```
