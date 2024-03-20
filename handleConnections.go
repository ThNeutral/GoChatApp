package main

import (
	"log"
	"net/http"
)

func handleConnections(w http.ResponseWriter, r *http.Request) {
	log.Println("Attempt to establish WebSocket connection registered")
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Error upgrading request to WebSocket: %s\n", err.Error())
		return
	}
	defer conn.Close()

	clients[conn] = true

	for {
		var msg InboundMessage
		err := conn.ReadJSON(&msg)
		if err != nil {
			log.Printf("Error parsing JSON: %s\n", err.Error())
			delete(clients, conn)
			return
		}

		broadcast <- msg
	}
}
