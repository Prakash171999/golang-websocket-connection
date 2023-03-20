package main

import (
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	// Solve cross-domain problems
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
} // use default options
