package main

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

func wsHandler(w http.ResponseWriter, r *http.Request) {
	upgrader := websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	for {
		_, p, err := conn.ReadMessage()
		if err != nil {
			conn.Close()
			return
		}
		log.Printf("read: %s", p)
	}
}

func main() {
	http.HandleFunc("/ws", wsHandler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
