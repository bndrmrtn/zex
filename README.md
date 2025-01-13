# Zex

An easy-to-use http router for Go.

## Installation

Create a new go project and install the package with the following command:

```
go get github.com/bndrmrtn/zex@latest
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
	userID := zx.Param(r, "id")
	w.Write([]byte("User ID is " + userID))
})
```

In this example, the `{id}` part of the URL is a placeholder that will match any string. When a request like `/user/123` is made, the value `123` will be captured as a parameter and can be accessed using `zx.Param(r, "id")`.

```go
app.Get("/optional/{id}?", func(...) {...})
```
Optional parameters are also supported via `?` mark.
In this case, both `/optional` and `/optional/any` will work.

### Route Parameter Validation

Zex also supports route parameter validation with the `@` symbol.
A route should look like this:
```go
app.Get("/optional/{id@uuid}", func(...) {...})
```

In this example, the `id` parameter is validated, and if it’s not a valid `UUIDv4`, the default NotFound error handler is triggered.

Zex also lets you define custom route parameter validators.

```go
app.RegisterParamValidator("png", func(value string) (string, error) {
	strs := strings.SplitN(value, ".", 2)
	if strs[len(strs)-1] != "png" {
		return "", errors.New("file extension is not png")
	}

	return strs[0], nil
})
```

## Error Handling

### Handlers With Error Return

Zex provides a flexible error-handling system.
Use `HandlerFuncWithErr` to write handlers that return errors, and wrap them with `NewWithErrorConverter` to ensure consistent error handling.
By default, `DefaultErrHandler` logs errors and sends appropriate HTTP responses, but you can define custom handlers to suit your needs.

```go
e := zex.NewWithErrorConverter()

app.Get("/error", e(func(w http.ResponseWriter, r *http.Request) error {
	return zex.NewError(400, "Bad request").SetInternal(errors.New("internal error"))
}))
```

Internal errors can be used to provide additional context to the error message.
They are not sent to the client but are logged for debugging purposes.
There are also built-in error types like `ErrNotFoun`, `ErrBadRequest`, etc.

### Built-in Error Type

Zex includes a customizable error type, `Error`, to simplify error management:
- Fields: Stores HTTP status, error message, and an optional internal error for debugging.
- Methods:
  - `Error()`: Returns the error message.
  - `Status()`: Returns the HTTP status code.
  - `SetInternal(err)`: Sets an internal error for additional context.
  - `Internal()`: Retrieves the internal error.

Predefined errors like `ErrBadRequest` help standardize client-side error responses.
Use `NewError` to create additional custom errors as needed.

## Built-in Utilities

Zex comes with `zx`, a handy utility library to make repetitive tasks faster and easier. (currently in development)

```go
func handler(w http.ResponseWriter, r *http.Request) {
	// get url param
	id := zx.Param(r, "id")
	// get url param as int
	num, err := zx.ParamInt(r, "num")
	// get query parameter
	query := zx.Query(r, "key")
	// bind body
	err := zx.Bind(r, &struct{}{})
	// get header
	header := zx.Header(r, "key")
}
```
