# Objects Server in Go
[![Build Status](https://travis-ci.org/gobjserver/gobjserver.svg?branch=master)](https://travis-ci.org/gobjserver/gobjserver)
[![codecov](https://codecov.io/gh/gobjserver/gobjserver/branch/master/graph/badge.svg)](https://codecov.io/gh/gobjserver/gobjserver)
[![Go Report Card](https://goreportcard.com/badge/github.com/gobjserver/gobjserver)](https://goreportcard.com/report/github.com/gobjserver/gobjserver)
[![GoDoc](https://godoc.org/github.com/gobjserver/gobjserver?status.svg)](https://godoc.org/github.com/gobjserver/gobjserver)
[![](https://images.microbadger.com/badges/image/gobjserver/gobjserver.svg)](https://microbadger.com/images/gobjserver/gobjserver)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/gobjserver/gobjserver/blob/master/LICENSE)

ObjServer is the server that store any objects in effective way.

## Features
- REST API and Swagger UI
- Middlewares (cors, gzip and static)
- RethinkDB
- Docker Build and CI with Travis

## Installation

Get the gobjserver repository

```
go get github.com/gobjserver/gobjserver

cd echo $GOPATH/src/github.com/gobjserver/gobjserver
```

And install dependencies

```
go get -u github.com/golang/dep/cmd/dep

dep ensure
```

## Running the tests

Run all tests

```
go test ./...
```

Or run all tests with coverage

```
bash script/coverage.sh
```

## Build and Run

Run main.go
``` bash
go run main.go
# serve at localhost:9000
```

Build and run native binary

``` bash
bash script/Build.sh

./gobjserver.out
```
Build native binary for multiple platforms (darwin, windows and linux)

```
bash script/BuildMulti.sh
```

## Environment variables

```bash
    # enable production mode, default is false
    env GBP_PROMODE=true

    # choose RethinkDB host, default is :28015
    env GBP_RETHINK_HOST=rethinkdbhost:28015

    # choose database name, default is objectdb
    env GBP_DB_NAME=dbname
```
## Docker support 

Build docker image

```
bash script/Dockerbuild.sh
```

Run docker container

```
docker run -d --name gobjserver -p 9000:9000 gobjserver/gobjserver
```

Run docker-compose
```bash
# GobjServer serves at localhost:8080
# RethinkDB Web UI serves at localhost:8000
docker-compose up -d
```
## Contributing

Please read [CONTRIBUTING.md](CONTRIBUTING.md) for details on our code of conduct, and the process for submitting pull requests to us.

## Versioning

We use [SemVer](http://semver.org/) for versioning. For the versions available, see the [tags on this repository](https://github.com/gobjserver/gobjserver/tags). 

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details

