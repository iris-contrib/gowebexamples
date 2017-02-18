+++
weight = 6
title = "JSON"
description = "This example will show how to encode and decode JSON data in the Go programming language."
+++

# [Go Web Examples:](/) JSON

This example will show how to encode and decode JSON data.
```
// json.go
package main

import (
	"gopkg.in/kataras/iris.v6"
	"gopkg.in/kataras/iris.v6/adaptors/httprouter"
)

type User struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Age       int    `json:"age"`
}

func main() {
	app := iris.New()
	app.Adapt(httprouter.New())

	app.Post("/decode", func(ctx *iris.Context) {
		var user User
		ctx.ReadJSON(&user)

		ctx.Writef("%s %s is %d years old!", user.Firstname, user.Lastname, user.Age)
	})

	app.Get("/encode", func(ctx *iris.Context) {
		peter := User{
			Firstname: "John",
			Lastname:  "Doe",
			Age:       25,
		}

		ctx.JSON(iris.StatusOK, peter)
	})

	app.Listen(":8080")
}
```
```
$ go run json.go

$ curl -s -XPOST -d'{"firstname":"Donald","lastname":"Trump","age":70}' http://localhost:8080/decode
Donald Trump is 70 years old!

$ curl -s http://localhost:8080/encode
{"firstname":"John","lastname":"Doe","age":25}

```
