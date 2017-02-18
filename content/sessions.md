+++
weight = 8
title = "Sessions"
description = "This example will show how to store data in session cookies using the popular gorilla/sessions package in the Go programming language."
+++

# [Go Web Examples:](/) Sessions

This example will show how to store data from a session.

You don't need any third-party library except Iris, but if you want you can use anything, remember Iris is fully compatible with the standard library. You can find a more detailed example by pressing [here](https://github.com/kataras/iris/blob/v6/adaptors/sessions/_example/main.go).

In this example we will only allow authenticated users to view our secret message on the `/secret` page. To get access to it, the will first have to visit `/login` to get a valid session cookie, which logs him in. Additionally he can visit `/logout` to revoke his access to our secret message.

```
// sessions.go
package main

import (
	"gopkg.in/kataras/iris.v6"
	"gopkg.in/kataras/iris.v6/adaptors/httprouter"
	"gopkg.in/kataras/iris.v6/adaptors/sessions"
)

var (
	key = "my_sessionid"
)

func secret(ctx *iris.Context) {

	// Check if user is authenticated
	if auth, _ := ctx.Session().GetBoolean("authenticated"); !auth {
		ctx.EmitError(iris.StatusForbidden)
		return
	}

	// Print secret message
	ctx.WriteString("The cake is a lie!")
}

func login(ctx *iris.Context) {
	session := ctx.Session()

	// Authentication goes here
	// ...

	// Set user as authenticated
	session.Set("authenticated", true)
}

func logout(ctx *iris.Context) {
	session := ctx.Session()

	// Revoke users authentication
	session.Set("authenticated", false)
}

func main() {
	app := iris.New()
	app.Adapt(httprouter.New())

	sess := sessions.New(sessions.Config{Cookie: key})
	app.Adapt(sess)

	app.Get("/secret", secret)
	app.Get("/login", login)
	app.Get("/logout", logout)

	app.Listen(":8080")
}

```
```
$ go run sessions.go

$ curl -s http://localhost:8080/secret
Forbidden

$ curl -s -I http://localhost:8080/login
Set-Cookie: mysessionid=MTQ4NzE5Mz...

$ curl -s --cookie "mysessionid=MTQ4NzE5Mz..." http://localhost:8080/secret
The cake is a lie!
```
