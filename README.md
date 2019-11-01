# go-microservices

## Requirements:

- Setup a database in mySQL.
- Install redis: https://redis.io/topics/quickstart
- Install gRPC: https://grpc.io/blog/installation/
- Install dep: https://golang.github.io/dep/docs/installation.html

## System architecture:


## How to run:

- front-controller:
    + git clone https://github.com/trongdth/go_microservices.git
    + cd front-controller
    + mv config/.env_sample.sh config/.env.sh
    + source config/.env.sh
    + dep ensure
    + go run server.go

- entry-cache:
    + git clone https://github.com/trongdth/go_microservices.git
    + redis-server
    + cd entry-cache
    + mv config/.env_sample.sh config/.env.sh
    + source config/.env.sh
    + dep ensure
    + go run server.go

- entry-store:
    + git clone https://github.com/trongdth/go_microservices.git
    + cd entry-store
    + mv config/.env_sample.sh config/.env.sh
    + source config/.env.sh
    + dep ensure
    + go run server.go