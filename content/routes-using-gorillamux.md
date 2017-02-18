+++
weight = 3
title = "Routes (using gorillamux)"
description = "This example will show how to register routes using the popular gorillamux router in the Go programming language."
+++

# [Go Web Examples:](/) Routes (using gorillamux)

This example will show how to register routes using the popular <a target="_blank" href="https://github.com/gorilla/mux">gorillamux</a> router.




```
// routes.go
package main

import (
	"gopkg.in/kataras/iris.v6"
	"gopkg.in/kataras/iris.v6/adaptors/gorillamux"
)

func main() {
	app := iris.New()
	// Adapt the "httprouter", you can use "gorillamux" too.
	app.Adapt(gorillamux.New())

	userAges := map[string]int{
		"Alice":  25,
		"Bob":    30,
		"Claire": 29,
	}

	// Equivalent with app.HandleFunc("GET", ...)
	app.Get("/users/{name}", func(ctx *iris.Context) {
		name := ctx.Param("name")
		age := userAges[name]

		ctx.Writef("%s is %d years old!", name, age)
	})

	app.Listen(":8080")
}

```
```
$ go run routes.go

$ curl -s http://localhost:8080/users/Bob
Bob is 30 years old!
```
