# Go To-Do

An example web app that enables to-do list functionality following common patterns established by the Go community and detailed by [Mat Ryer](https://grafana.com/author/mat_ryer/) in this [blog post](https://grafana.com/blog/2024/02/09/how-i-write-http-services-in-go-after-13-years/).

The project aims to use stdlib packages where possible to minimize dependencies (i.e. no frameworks).

## API

### `GET /api/v1/todos`

List all to-dos

### `POST /api/v1/todos`

Create a new to-do

### `GET /api/v1/todos/{id}`

Retrieve a single to-do

### `PUT /api/v1/todos/{id}`

Update a todo's status

### `DELETE /api/v1/todos/{id}`

Delete a todo


## Local Dev

### Running the WebServer

```sh
go run cmd/rest/main.go
```