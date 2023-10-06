package websocket

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

// I have no idea how any of this works honestly
// The tutorial explained nothing

var upgrader = websocket.Upgrader {
	ReadBufferSize: 1024,
	WriteBufferSize: 1024,
}

// Capital letter function name to make it visible to the rest of the project
// Kinda weird ngl
func Upgrade(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return conn, nil
}