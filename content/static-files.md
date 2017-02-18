+++
weight = 7
title = "Static Files"
description = "This example will show how to serve static files like CSS or JS from a specific directory using the StaticWeb in the Go programming language."
+++

# [Go Web Examples:](/) Static Files

This example will show how to serve static files like CSSs, JavaScripts or images from a specific directory.

```
// static-files.go
package main

import (
	"gopkg.in/kataras/iris.v6"
	"gopkg.in/kataras/iris.v6/adaptors/httprouter"
)

func main() {
	app := iris.New()
	app.Adapt(httprouter.New())
	// first parameter is the request path
	// second is the operating system directory
	app.StaticWeb("/static", "./assets")

	app.Listen(":8080")
}

```
```
$ tree assets/
assets/
└── css
    └── styles.css
```
```
$ go run static-files.go

$ curl -s http://localhost:8080/static/css/styles.css
body {
    background-color: black;
}
```
