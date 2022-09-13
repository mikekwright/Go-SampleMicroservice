Go simple Microservice
====================================

This is a demonstration of what it could be like to create a microservice
in golang.  Unlike the C# example where ASP.NET core is standard and part
of the dotnet ecosystem, in go the common tool is [gin](https://gin-gonic.com/)
which can be installed as a dependency.

    go get github.com/gin-gonic/gin

### Install Swag for swagger generation

    go install github.com/swaggo/swag/cmd/swag@latest

### Install Wire for dependency injection

    go install github.com/google/wire/cmd/wire@latest

Running
------------------------------------

This service can be started by running the following command

    go run ./src

It will start a service you can reference as demonstrated below

    curl http://localhost:8080/user

Running (Docker)
-------------------------------------

There is also an example of how you can run the service using
docker

    make build-docker
    make run-docker

Swagger
--------------------------------------

Swagger is configured using `gin-swagger` and can be found by opening
a browser to [http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html)
