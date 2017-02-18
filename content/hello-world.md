+++
weight = 1
title = "Hello World"
description = "This example will show how to start a webserver on port 8080 and print the classic \"hello world\" message in the Go programming language."
+++

# [Go Web Examples:](/) Hello World

This example will show how to start a webserver on port 8080 and print the classic "hello world" message.

```
// hello-world.go
package main

import (
	"gopkg.in/kataras/iris.v6"
	"gopkg.in/kataras/iris.v6/adaptors/httprouter"
)

func main() {
  app := iris.New()
	// Adapt the "httprouter", faster,
	// but it has limits on named path parameters' validation,
	// you can adapt "gorillamux" if you need regexp path validation!
	app.Adapt(httprouter.New())

	app.HandleFunc("GET", "/", func(ctx *iris.Context) {
		ctx.Writef("hello world\n")
	})

	app.Listen(":8080")
}
```
```
$ go run hello-world.go

$ curl -s http://localhost:8080/
hello world
```
