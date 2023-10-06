package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader {
	ReadBufferSize: 1024,
	WriteBufferSize: 1024,
	// Function used to check origin of request
	// For now just return true, but can have logic in there
	CheckOrigin: func(r *http.Request) bool { return true },
}

func reader(conn *websocket.Conn) {
	// I guess for is both a for and while loop
	// In this case a while loop
	for { 
		messageType, msg, err := conn.ReadMessage()
		// If there's an error log it and break out of loop
		if err != nil {
			log.Println(err)
			return
		}

		fmt.Println(string(msg))

		// I have no idea what this is doing honestly
		if err := conn.WriteMessage(messageType, msg); err != nil {
			log.Println(err)
			return
		}
	}
}

func serveWebsocket(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Host)

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}

	reader(ws)
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