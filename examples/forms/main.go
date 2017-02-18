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

		details := ContactDetails{
			Email:   ctx.FormValue("email"),
			Subject: ctx.FormValue("subject"),
			Message: ctx.FormValue("message"),
		}

		// do something with details
		_ = details

		ctx.Render("forms.html", struct{ Success bool }{true})
	})

	app.Listen(":8080")
}
