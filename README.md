# go_rest_api
> A simple REST API written with Go. Fetches and stores an article

## Build Setup

Download this code to any desired directory.
In a terminal, navigate to the selected directory.

``` bash

# build server
go build main.go

# run server
go run main.go

```

## Testing

``` bash

# once server is running, another terminal may run tests
go test

```

## Usage

The server will be running on localhost:9090.

URLs being served are: 
> GET localhost:9090/articles - get all article titles as a JSON array
> GET localhost:9090/articles/(name) - get an article as a JSON object
> PUT localhost:9090/articles/(name) - add/edit an article

Using [Curl](https://github.com/curl/curl) add/edit queries can be made to the API via the following command:

``` bash

curl -X PUT http://localhost:9090/articles/sample -d ‘This is an article content sample’

```

