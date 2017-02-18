+++
weight = 9
title = "Websockets"
description = "This example will show how to work with websockets in Go. We will build a simple server which echoes back everything we send to it."
+++

# [Go Web Examples:](/) Websockets

This example will show how to work with websockets in Go. We will build a simple server which echoes back everything we send to it.
You don't need any third-party library except `adaptors/websocket`, but if you want you can use anything. Remember: Iris is fully compatible with the standard library.

You can find more websocket examples by pressing [here](https://github.com/kataras/iris/tree/v6/adaptors/websocket/_examples).

```
// websockets.go
package main

import (
	"fmt"

	"gopkg.in/kataras/iris.v6"
	"gopkg.in/kataras/iris.v6/adaptors/httprouter"
	"gopkg.in/kataras/iris.v6/adaptors/websocket"
)

const (
	ReadBufferSize =  1024
	WriteBufferSize = 1024
)

func handleConnection(c websocket.Connection){

	// Read events from browser
	c.On("chat", func(msg string) {

		// Print the message to the console
		fmt.Printf("%s sent: %s\n", c.Context().RemoteAddr(), msg)

		// Write message back to browser
		c.Emit("chat", msg)
	})

}

func main() {
	app := iris.New()
	app.Adapt(httprouter.New())

	// create our echo websocket server
	ws := websocket.New(websocket.Config{
	  ReadBufferSize: 1024,
		WriteBufferSize: 1024,
		Endpoint: "/echo",
	})

	ws.OnConnection(handleConnection)

	// Adapt the websocket server.
	// you can adapt more than one of course.
	app.Adapt(ws)

	app.Get("/", func(ctx *iris.Context) {
		ctx.ServeFile("websockets.html")
	})

	app.Listen(":8080")
}
```
```
<!-- websockets.html -->
<input id="input" type="text" />
<button onclick="send()">Send</button>
<pre id="output"></pre>
<script src="/iris-ws.js"></script>
<script>
	var input = document.getElementById("input");
	var output = document.getElementById("output");

	// Ws comes from the auto-served '/iris-ws.js'
	var socket = new Ws("ws://localhost:8080/echo");
	socket.OnConnect(function () {
		output.innerHTML += "Status: Connected\n";
	});

	socket.OnDisconnect(function () {
		output.innerHTML += "Status: Disconnected\n";
	});

	// read events from the server
	socket.On("chat", function (msg) {
		output.innerHTML += "Server: " + msg + "\n";
	});

	function send() {
		// send chat event data to the server
		socket.Emit("chat", input.value);
		input.value = "";
	}
</script>
```
```
$ go run websockets.go
[127.0.0.1]:53403 sent: Hello Go Web Examples, you're doing great!
```
<div class="demo">
	<input type="text">
	<button>Send</button>
	<pre>Status: Connected
Server: Hello Go Web Examples, you're doing great!</pre>
</div>
