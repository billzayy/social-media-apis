# Notification Service - Social Media Apis Golang

This service is created for Get and Send Notification APIs with RabbitMQ

Using: Supabase Postgres, Redis, Websocket, RabbitMq JWT and GoLang to build.

***P/S***: If you want to develop your own api with this project, please add file `.env` into [`./internal`](./internal/) and fill the file with [`Preparation for Develop`](#preparation-for-develop) below.

## Install & Download 
Before run the application, ensure that you already downloaded all necessary packages and apps:
* [`Go Package`](https://go.dev/doc/install)
* [`Docker Desktop`](https://docs.docker.com/desktop/)
* [`Postman` ](https://www.postman.com/downloads/)
* [`Protocol Buffer Compiler`](https://protobuf.dev/installation/)
* [`Swagger for Golang`](https://goswagger.io/go-swagger/install/install-source/)

You can check if some tools are already downloaded or not with : 
```bash
$ go version # check golang version
$ swag --version # check swagger version 
$ protoc --version # check protocol buffer compiler
```

Before you want to developer this service, please run `go get all` or `make -f ./scripts/Makefile install` to download all packages.
```bash
$ go get all
or
$ make -f ./scripts/Makefile install
```
P/S: If some packages are need to upgrade, you just run `go get <package need upgrade>` to upgrade.

```
## Preparation for develop
All you need is create `.env` file and replace all values `<>` of this properties:
```bash
POSTGRES_USERNAME="<Postgres Username>" # default: ""
POSTGRES_PASSWORD="<Postgres Password>" # default: ""
POSTGRES_HOST="<Postgres Host>" # default: ""
POSTGRES_PORT="<Postgres Port>" # default: 5432
POSTGRES_DATABASE="<Postgres Database>" # default: ""

REDIS_ADDR="<Redis Address>" # default: ""
REDIS_USERNAME="<Redis Username>" # default: default
REDIS_PASSWORD="<Redis Password>" # default: ""

RABBITMQ_USERNAME="<RabbitMQ Username>"  # default: "guest"
RABBITMQ_PASSWORD="<RabbitMQ Password>" # default: "guest"
RABBITMQ_HOST="<RabbitMQ Host>" # default: "localhost"
RABBITMQ_PORT="<RabbitMQ Port>" # default: 5672

REST_PORT="<Rest Port>" # default: 3000
GRPC_PORT="<Grpc Port>" # default: 50051
WS_PORT="<Websocket Port>" # default: 90000

ACCESS_TOKEN_KEY="<Access Token>" # default: ""
REFRESH_TOKEN_KEY="<Refresh Token>" # default: ""
```
*P/S:* 
* Postgres in this project is [`Supabase`](https://supabase.com/). So all Postgres configuration is based on **Supabase**.
* Redis in this project is from [`Redis Cloud`](https://app.redislabs.com/).

## Run
To run all step to build and push image project, run `make all` command:
```bash
$ make -f ./scripts/Makefile all
or
$ make all # if you are staying inside ./scripts folder
``` 

If you want to build image for this project step-by-step, run `make` command with some options:
* `docker` : To automatic run build and push docker image command.
    ```bash
    $ make -f ./scripts/Makefile docker
    ```
* `install` : To download all go packages.
    ```bash
    $ make -f ./scripts/Makefile install
    ```
* `allow` : To allow permission `./scripts/generate.sh` file to run protoc config.
    ```bash
    $ make -f ./scripts/Makefile allow
    ```
* `generate` : To run protoc configure `./scripts/generate.sh`.
    ```bash
    $ make -f ./scripts/Makefile generate
    ```
* `build` : To build docker image for `notification-service` service with [`Dockerfile`](./Dockerfile)
    ```bash
    $ make -f ./scripts/Makefile build
    $ docker images # check created images
    ```
* `push` : To push docker images into DockerHub
    ```bash
    $ make -f ./scripts/Makefile push
* `clean` : To clean all cache docker images and volumes.
    ```bash
    $ make -f ./scripts/Makefile clean
    ```
---
$${\color{lightgreen}HAPPY \space CODING \space 😉 \space !!!}$$	
