+++
weight = 10
title = "Middleware (Basic)"
description = "This example will show how to create basic logging middleware in the Go programming language."
+++

# [Go Web Examples:](/) Middleware (Basic)

This example will show how to create basic logging middleware in Go.

An `Iris` Middleware is just a chain of `iris.HandlerFunc` runs explicity by calling `context.Next()`. On each request route's handlers share the same `*Context`, so these handlers can share information or stop execution the execution of the next in the chain.(When `context.Next` is missing then the next handler doesn't run).


```
// basic-middleware.go
package main

import (
	"gopkg.in/kataras/iris.v6"
	"gopkg.in/kataras/iris.v6/adaptors/httprouter"
)

func logging(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL.Path)
		f(w, r)
	}
}

func foo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "foo")
}

func bar(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "bar")
}

func main() {
	app := iris.New()
	app.Adapt(iris.DevLogger())
	app.Adapt(httprouter.New())

	app.Get("/foo", logging(foo))
	app.Get("/bar", logging(bar))

	http.ListenAndServe(":8080", nil)
}
```
```
$ go run basic-middleware.go
2017/02/10 23:59:34 /foo
2017/02/10 23:59:35 /bar
2017/02/10 23:59:36 /foo?bar

$ curl -s http://localhost:8080/foo
$ curl -s http://localhost:8080/bar
$ curl -s http://localhost:8080/foo?bar
```
