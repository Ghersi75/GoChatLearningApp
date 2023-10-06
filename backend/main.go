package main

import (
	"fmt"
	"net/http"

	"backend/pkg/websocket"
)

func serveWebsocket(w http.ResponseWriter, r *http.Request) {
	ws, err := websocket.Upgrade(w, r)
	if err != nil {
		fmt.Fprintf(w, "%+V\n", err)
	}
	go websocket.Writer(ws)
	websocket.Reader(ws)
}

func setupRoutes() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Simple Server")
	})

	http.HandleFunc("/ws", serveWebsocket)
}

func main() {
	setupRoutes()
	http.ListenAndServe(":8080", nil)
}