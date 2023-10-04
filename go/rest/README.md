# Go REST Service

This is a simple REST service built in Go that returns a JSON response with a message.

## Getting Started

To run the REST service locally, you'll need to have Go and Docker installed on your machine.


1. Clone the repository:

```
git clone https://github.com/bcdunbar/mixedbag.git
```

2. Build the Docker image:

```
docker build -t go-rest-service .
```

3. Run the Docker container:

```
docker run -p 8080:8080 go-rest-service
```

4. Access the REST endpoint in your browser or using a tool like `curl`: http://localhost:8080/


## Running Tests

To run the unit tests for the REST service, you can use the following command:

```
go test -v
```


