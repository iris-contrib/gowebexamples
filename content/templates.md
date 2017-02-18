+++
weight = 4
title = "Templates"
description = "This example will show how to render a simple list of TODO items into an html page using the html/template package in the Go programming language."
+++

# [Go Web Examples:](/) Templates

This example will show how to render a simple list of TODO items into an html page using the `html/template view adaptor`.

Iris supports 5 template engines out-of-the-box, you can still use any external golang template engine,
as `context.ResponseWriter` is an `io.Writer` and `http.ResponseWriter`.

All of these five template engines have common features with common API,
like Layout, Template Funcs, Party-specific layout, partial rendering and more.

      The standard html, based on github.com/kataras/go-template/tree/master/html
      its template parser is the golang.org/pkg/html/template/.

      Django, based ongithub.com/kataras/go-template/tree/master/django
      its template parser is the github.com/flosch/pongo2

      Pug(Jade), based on github.com/kataras/go-template/tree/master/pug
      its template parser is the github.com/Joker/jade

      Handlebars, based on github.com/kataras/go-template/tree/master/handlebars
      its template parser is the github.com/aymerick/raymond

      Amber, based on github.com/kataras/go-template/tree/master/amber
      its template parser is the github.com/eknkc/amber


View engine is the `adaptors/view` package.

- standard html  | view.HTML(...)
- django         | view.Django(...)
- pug(jade)      | view.Pug(...)
- handlebars     | view.Handlebars(...)
- amber          | view.Amber(...)


```
// todos.go
package main

import (
	"gopkg.in/kataras/iris.v6"
	"gopkg.in/kataras/iris.v6/adaptors/httprouter"
	"gopkg.in/kataras/iris.v6/adaptors/view"
)

type Todo struct {
	Task string
	Done bool
}

func main() {
  // Configuration is optional
  app := iris.New(iris.Configuration{Gzip: false, Charset: "UTF-8"})

  // Adapt a logger which will print all errors to os.Stdout
  app.Adapt(iris.DevLogger())

  // Adapt the httprouter (we will use that on all examples)
  app.Adapt(httprouter.New())

	// Parse all files inside `./mytemplates` directory ending with `.html`
	app.Adapt(view.HTML("./mytemplates", ".html"))

	todos := []Todo{
		{"Learn Go", true},
		{"Read Go Web Examples", true},
		{"Create a web app in Go", false},
	}

	app.Get("/", func(ctx *iris.Context) {
	  ctx.Render("todos.html", struct{ Todos []Todo }{todos})
	})

	app.Listen(":8080")
}
```
```
<!-- mytemplates/todos.html -->
<h1>Todos</h1>
<ul>
	{{range .Todos}}
		{{if .Done}}
			<li><s>{{.Task}}</s></li>
		{{else}}
			<li>{{.Task}}</li>
		{{end}}
	{{end}}
</ul>
```
```
$ go run todos.go
```
<div class="demo">
	<h1>Todos</h1>
	<ul>
		<li><s>Learn Go</s></li>
		<li><s>Read Go Web Examples</s></li>
		<li>Create a web app in Go</li>
	</ul>
</div>
