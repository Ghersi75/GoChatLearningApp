package websocket

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

// I have no idea how any of this works honestly
// The tutorial explained nothing

var upgrader = websocket.Upgrader {
	ReadBufferSize: 1024,
	WriteBufferSize: 1024,
	// Function used to check origin of request
	// For now just return true, but can have logic in there
	CheckOrigin: func(r *http.Request) bool { return true },
}

// Capital letter function name to make it visible to the rest of the project
// Kinda weird ngl
func Upgrade(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return ws, err
	}
	return ws, nil
}

func Reader(conn *websocket.Conn) {
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

func Writer(conn *websocket.Conn) {
	for {
		fmt.Println("Sending")
		messageType, msg, err := conn.NextReader()
		if err != nil {
			fmt.Println(err)
			return
		}

		w, err := conn.NextWriter(messageType)
		if err != nil {
			fmt.Println(err)
			return
		}

		if _, err := io.Copy(w, msg); err != nil {
			fmt.Println(err)
			return
		}
		if err := w.Close(); err != nil {
			fmt.Println(err)
			return
		}
	}
}