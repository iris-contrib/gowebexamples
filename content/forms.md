+++
weight = 5
title = "Forms"
description = "This example will show how to simulate a contact form and parse the message into a struct using the Go programming language."
+++

# [Go Web Examples:](/) Forms

This example will show how to simulate a contact form and parse the message into a struct.

```
// forms.go
package main

import (
	"gopkg.in/kataras/iris.v6"
	"gopkg.in/kataras/iris.v6/adaptors/httprouter"
	"gopkg.in/kataras/iris.v6/adaptors/view"
)

// ContactDetails the information from user
type ContactDetails struct {
	Email   string
	Subject string
	Message string
}

func main() {
	app := iris.New()
	app.Adapt(httprouter.New())

	// Parse all files inside `./mytemplates` directory ending with `.html`
	app.Adapt(view.HTML("./mytemplates", ".html"))

	app.Get("/", func(ctx *iris.Context) {
		ctx.Render("forms.html", nil)
	})

	// Equivalent with app.HandleFunc("POST", ...)
	app.Post("/", func(ctx *iris.Context) {

		// details := ContactDetails{
		// 	Email:   ctx.FormValue("email"),
		// 	Subject: ctx.FormValue("subject"),
		// 	Message: ctx.FormValue("message"),
		// }

		// or simply:
		var details ContactDetails
		ctx.ReadForm(&details)
		
		// do something with details
		_ = details

		ctx.Render("forms.html", struct{ Success bool }{true})
	})

	app.Listen(":8080")
}

```
```
<!-- mytemplates/forms.html -->
{{if .Success}}
	<h1>Thanks for your message!</h1>
{{else}}
	<h1>Contact</h1>
	<form method="POST">
		<label>Email:</label><br />
		<input type="text" name="email"><br />
		<label>Subject:</label><br />
		<input type="text" name="subject"><br />
		<label>Message:</label><br />
		<textarea name="message"></textarea><br />
		<input type="submit">
	</form>
{{end}}
```
```
$ go run forms.go
```
<div class="demo">
	<h1>Contact</h1>
	<form method="POST">
		<label>Email:</label><br />
		<input type="text" name="email"><br />
		<label>Subject:</label><br />
		<input type="text" name="subject"><br />
		<label>Message:</label><br />
		<textarea name="message"></textarea><br />
		<input type="submit">
	</form>
</div>
