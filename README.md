API using Golang 1.23 and gin-gonic/gin
Startup: 

To start the Go API, run the following command in the terminal:

```shell
go mod tidy && go run cmd/go-api/main.go
```

Docker:

To start the Go API using Docker, make sure you have Docker installed and follow these steps:

1. Create a Dockerfile in the root directory of your project (`/c:/Users/minu/code/study/go-api`) with the provided content.

2. Create a `docker-compose.yml` file in the same directory with the provided content.

3. Open a terminal and navigate to the project directory (`/c:/Users/minu/code/study/go-api`).

4. Run the following command to start the Go API using Docker Compose:

```shell
docker-compose up
```

This will build the Docker image and start the container, exposing the API on port 8080.
