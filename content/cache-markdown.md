+++
weight = 12
title = "Render Markdown and Cache Handler"
description = "This example will show how to send Markdown contents and use the Cache handler to cache the handler's response for your page using the Go programming language."
+++

# [Go Web Examples:](/) Render Markdown and Cache Handler

This example will show how to send Markdown contents and use the Cache handler to cache the handler's response.

```
// file markdown-cache.go
package main

import (
	"time"

	"gopkg.in/kataras/iris.v6"
	"gopkg.in/kataras/iris.v6/adaptors/httprouter"
)

var testMarkdownContents = `## Hello Markdown

This is a sample of Markdown contents



Features
--------

All features of Sundown are supported, including:

*   **Compatibility**. The Markdown v1.0.3 test suite passes with
    the --tidy option.  Without --tidy, the differences are
    mostly in whitespace and entity escaping, where blackfriday is
    more consistent and cleaner.

*   **Common extensions**, including table support, fenced code
    blocks, autolinks, strikethroughs, non-strict emphasis, etc.

*   **Safety**. Blackfriday is paranoid when parsing, making it safe
    to feed untrusted user input without fear of bad things
    happening. The test suite stress tests this and there are no
    known inputs that make it crash.  If you find one, please let me
    know and send me the input that does it.

    NOTE: "safety" in this context means *runtime safety only*. In order to
    protect yourself against JavaScript injection in untrusted content, see
    [this example](https://github.com/russross/blackfriday#sanitize-untrusted-content).

*   **Fast processing**. It is fast enough to render on-demand in
    most web applications without having to cache the output.

*   **Thread safety**. You can run multiple parsers in different
    goroutines without ill effect. There is no dependence on global
    shared state.

*   **Minimal dependencies**. Blackfriday only depends on standard
    library packages in Go. The source code is pretty
    self-contained, so it is easy to add to any project, including
    Google App Engine projects.

*   **Standards compliant**. Output successfully validates using the
    W3C validation tool for HTML 4.01 and XHTML 1.0 Transitional.

	[this is a link](https://github.com/kataras/iris) `

func main() {
	app := iris.New()
	// output startup banner and error logs on os.Stdout
	app.Adapt(iris.DevLogger())
	// set the router, you can choose gorillamux too
	app.Adapt(httprouter.New())

	app.Get("/hi", app.Cache(func(c *iris.Context) {
		c.WriteString("Hi this is a big content, do not try cache on small content it will not make any significant difference!")
	}, time.Duration(10)*time.Second))

	bodyHandler := func(ctx *iris.Context) {
		ctx.Markdown(iris.StatusOK, testMarkdownContents)
	}

	expiration := time.Duration(5 * time.Second)

	app.Get("/", app.Cache(bodyHandler, expiration))

	// if expiration is <=time.Second then the cache tries to set the expiration from the "cache-control" maxage header's value(in seconds)
	// // if this header doesn't founds then the default is 5 minutes
	app.Get("/cache_control", app.Cache(func(ctx *iris.Context) {
		ctx.HTML(iris.StatusOK, "<h1>Hello!</h1>")
	}, -1))

	app.Listen(":8080")
}

```

```
$ go run markdown-cache.go
```

<h2>Hello Markdown</h2>

<p>This is a sample of Markdown contents</p>

<h2>Features</h2>

<p>All features of Sundown are supported, including:</p>

<ul>
<li><p><strong>Compatibility</strong>. The Markdown v1.0.3 test suite passes with
the &ndash;tidy option.  Without &ndash;tidy, the differences are
mostly in whitespace and entity escaping, where blackfriday is
more consistent and cleaner.</p></li>

<li><p><strong>Common extensions</strong>, including table support, fenced code
blocks, autolinks, strikethroughs, non-strict emphasis, etc.</p></li>

<li><p><strong>Safety</strong>. Blackfriday is paranoid when parsing, making it safe
to feed untrusted user input without fear of bad things
happening. The test suite stress tests this and there are no
known inputs that make it crash.  If you find one, please let me
know and send me the input that does it.</p>

<p>NOTE: &ldquo;safety&rdquo; in this context means <em>runtime safety only</em>. In order to
protect yourself against JavaScript injection in untrusted content, see
<a href="https://github.com/russross/blackfriday#sanitize-untrusted-content">this example</a>.</p></li>

<li><p><strong>Fast processing</strong>. It is fast enough to render on-demand in
most web applications without having to cache the output.</p></li>

<li><p><strong>Thread safety</strong>. You can run multiple parsers in different
goroutines without ill effect. There is no dependence on global
shared state.</p></li>

<li><p><strong>Minimal dependencies</strong>. Blackfriday only depends on standard
library packages in Go. The source code is pretty
self-contained, so it is easy to add to any project, including
Google App Engine projects.</p></li>

<li><p><strong>Standards compliant</strong>. Output successfully validates using the
W3C validation tool for HTML 4.01 and XHTML 1.0 Transitional.</p>

<p><a href="https://github.com/kataras/iris">this is a link</a></p></li>
</ul>
