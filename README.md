# Social Media Apis Golang

This project is created for develop common RESTful APIs of Social Media Applications.

Services: 
* Authenticate Service: created for API Login, SignUp and Refresh Tokens.
* User Service: created for Users's Account Management
* Post Service: created for CRUD Post APIs and Interact APIs
* Notification Service: created for Get and Send Notification APIs with RabbitMQ
* Chat Service: created for Get and Send Chat APIs with Websocket server

Deployment:
[![Deploy to Koyeb](https://www.koyeb.com/static/images/deploy/button.svg)](https://app.koyeb.com/deploy?name=social-media-apis&type=git&repository=billzayy%2Fsocial-media-apis&branch=main&builder=dockerfile&dockerfile=Dockerfile.koyeb&privileged=true&instance_type=free&regions=was&instances_min=0&autoscaling_sleep_idle_delay=300&env%5BAPP_ENV%5D=production)

***P/S***: If you want to develop your own api with this project, please add file `.env` into `./api-gateway/internal/` and fill the file with [`Preparation for Develop`](#preparation-for-develop) below

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

## Preparation for develop
All you need is create `.env` file and replace all values `< >` of this properties:
```bash
AUTH_PORT="<Auth Service Port>" #default: 50051
POST_PORT="<Post Service Port>" # default: 50052
USER_PORT="<User Service Port>" # default: 50053
NOTIFICATION_PORT="<Notification Service Port>" # default: 50054
CHAT_PORT="<Chat Service Port>" # default: 50055

REST_PORT="<Gateway Service Port>" # default: 8000
CHAT_MESSAGE="<Websocket Port>" # default: 9000

RABBITMQ_PORT="<RabbitMQ Service Port>" # default: 5672
RABBITMQ_USER="<RabbitMQ UserName>" # default: guest
RABBITMQ_PASS="<RabbitMQ Password>" # default: guest
RABBITMQ_MANAGEMENT="<RabbitMQ UI Port>" # default: 15672

ACCESS_TOKEN_KEY="<Access Key>" # default: ""
REFRESH_TOKEN_KEY="<Refresh Key>" # default: ""
```
## Run
To run all project, run `make` command:
```bash
$ make all
``` 
P/S: After you run all successful, using postman or enter [`localhost:<REST_PORT>/swagger/index`]()

If you want to run this project step-by-step, run `make` command with some options:
* `swagger` : To format and update docs swagger on `./api-gateway/`
    ```bash
    $ make swagger
    ```
* `build-image` : To build docker image for `api-gateway` service with [`Dockerfile`](./Dockerfile)
    ```bash
    $ make build
    $ docker images # check created images
    ```
* `run-docker` : To initiate application with `docker-compose` on file [`docker-compose.yml`](./docker-compose.yml). 
    ```bash
    $ make run-docker
    ```
* `run-docker-with-env` : To initiate application with `docker-compose` on file [`docker-compose.yml`](./docker-compose.yml) with custom `./env`.
    ```bash
    $ make run-docker-with-env
    ```
* `push` : To push docker images into DockerHub
    ```bash
    $ make push
* `stop` : To stop docker containers.
    ```bash
    $ make stop
* `down` : To stop and remove all docker containers.
    ```bash
    $ make down
* `clean` : To clean all cache docker images and volumes.
    ```bash
    $ make clean
    ```
---
$${\color{lightgreen}HAPPY \space CODING \space 😉 \space !!!}$$	
