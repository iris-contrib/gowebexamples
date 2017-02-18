+++
weight = 1
title = "Hello World"
description = "This example will show how to start a webserver on port 8080 and print the classic \"hello world\" message in the Go programming language."
+++

# [Go Web Examples:](/) Hello World

This example will show how to start a webserver on port 8080 and print the classic "hello world" message.

For this we have to actually <a target="_blank" href="https://golang.org/doc/articles/go_command.html#tmp_3">go get</a> the popular <a target="_blank" href="https://github.com/kataras/iris">kataras/iris</a> library like so:
```
$ go get gopkg.in/kataras/iris.v6
```
From now on, every application we write will be able to make use of this library!

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
