# go_rest_api
> A simple REST API written with Go. Fetches and stores an article

## Build Setup

Download this code to any directory.
In a terminal, navigate to the selected directory.

``` bash

# build server
go build main.go

# run server
go run main.go

```

## Testing

``` bash

# when the server is running, tests may be run using
go test

```

## Usage

The server will be running on localhost:9090.

The functionality being served is as follows:

Get all article titles as a JSON array
- GET localhost:9090/articles 
	
Get an article as a JSON object
- GET localhost:9090/articles/(name) 

Add/Update an article
- PUT localhost:9090/articles/(name)

Using [Curl](https://github.com/curl/curl) add/update queries can be made to the API via the following command:

``` bash

curl -X PUT http://localhost:9090/articles/sample -d ‘This is an article content sample’

```

