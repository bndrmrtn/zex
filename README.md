# Zex

An easy-to-use http router for Go.

## Installation

Create a new go project and install the package with the following command:

```
go get github.com/bndrmrtn/go-gale@latest
```

## Usage

```go
app := zex.New() // with default configuration
app := zex.New(&zex.Config{
    Development: true,
    NotFoundHandler: http.NotFound,
}) // with custom configuration
```

## Routing

Zex provides a straightforward way to define and manage HTTP routing. Routing allows you to define specific paths and associate them with handler functions that process incoming HTTP requests. Zex supports dynamic routes, where parts of the path are treated as variables, allowing you to easily capture parameters from the URL.

### Defining Routes

With Zex, you can define routes using methods like `Get()`, `Post()`, and other HTTP method handlers. Each route is associated with a URL pattern, and when a request matches this pattern, the corresponding handler function is executed.

```go
// simple http get request
app.Get("/hello", func(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Hello, world!"))
})
```

In this example, any `GET` request to the `/hello` path will invoke the handler that responds with `"Hello, world!"`.

### Dynamic Routes

Zex allows you to define dynamic routes that can capture variables from the URL path. These variables are then accessible in your handler function.

```go
app.Get("/user/{id}", func(w http.ResponseWriter, r *http.Request) {
    userID := zex.Param(r, "id")
    w.Write([]byte("User ID is " + userID))
})
```

In this example, the `{id}` part of the URL is a placeholder that will match any string. When a request like `/user/123` is made, the value `123` will be captured as a parameter and can be accessed using `zex.Param(r, "id")`.
