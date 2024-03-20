package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

const (
	baseHTML     = "./templates/base.html"
	homePageHTML = "./templates/homePage.html"

	rootRoute    = "/"
	connectRoute = "/connect"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var clients = make(map[*websocket.Conn]bool)
var broadcast = make(chan InboundMessage)

type InboundMessage struct {
	ChatMessage string `json:"chat_message"`
}

type OutboundMessage struct {
	Content string `json:"content"`
}

func main() {
	http.HandleFunc(rootRoute, homePageHandler)
	http.HandleFunc(connectRoute, handleConnections)

	go handleMessages()

	log.Println("Server started on 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Printf("Error starting error: %s\n", err.Error())
	}
}
