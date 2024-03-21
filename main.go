package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"text/template"
	"time"

	"github.com/olahol/melody"
)

const (
	port = ":8080"

	rootRoute     = "/"
	chatRoomRoute = "/chat-room"

	pathToTemplates = "./templates"
	indexHTML       = pathToTemplates + "/index.html"
	chatRoomHTML    = pathToTemplates + "/chat-room.html"
	messageHTML     = pathToTemplates + "/message.html"
)

var melodyInstance = melody.New()

type OutboundMessage struct {
	Timestamp string
	Message   string
}

type InboundMessage struct {
	ChatMessage string `json:"chat_message"`
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	log.Printf("Recieved request to %s", r.URL.Path)
	tmpl, err := template.ParseFiles(indexHTML, chatRoomHTML)
	if err != nil {
		fmt.Fprint(w, "Internal Error")
		log.Printf("Encountered error while handling %s\n" + r.URL.Path)
		log.Printf("Error: %s\n", err.Error())
		w.WriteHeader(500)
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		fmt.Fprint(w, "Internal Error")
		log.Printf("Encountered error while handling %s\n" + r.URL.Path)
		log.Printf("Error: %s\n", err.Error())
		w.WriteHeader(500)
	}
}

func handleChatRoom(w http.ResponseWriter, r *http.Request) {
	log.Printf("Recieved request to %s", r.URL.Path)
	err := melodyInstance.HandleRequest(w, r)
	if err != nil {
		fmt.Fprint(w, "Internal Error")
		log.Printf("Encountered error while handling %s\n" + r.URL.Path)
		log.Printf("Error: %s\n", err.Error())
		w.WriteHeader(500)
		return
	}
}

func handleMessage(s *melody.Session, msg []byte) {
	log.Print("Handling message\n")
	tmpl, err := template.ParseFiles(messageHTML)
	if err != nil {
		log.Print("Encountered error while handling message broadcast\n")
		log.Printf("Error: %s\n", err.Error())
		return
	}

	var inbound InboundMessage
	err = json.Unmarshal(msg, &inbound)
	if err != nil {
		log.Print("Encountered error while handling message broadcast\n")
		log.Printf("Error: %s\n", err.Error())
		return
	}

	var outbound OutboundMessage
	var html bytes.Buffer
	outbound.Message = inbound.ChatMessage
	outbound.Timestamp = time.Now().Format(time.RFC850)

	err = tmpl.Execute(&html, outbound)
	if err != nil {
		log.Print("Encountered error while handling message broadcast\n")
		log.Printf("Error: %s\n", err.Error())
		return
	}

	err = melodyInstance.Broadcast(html.Bytes())
	if err != nil {
		log.Print("Encountered error while handling message broadcast\n")
		log.Printf("Error: %s\n", err.Error())
		return
	}
}

func main() {
	http.HandleFunc(rootRoute, handleRoot)
	http.HandleFunc(chatRoomRoute, handleChatRoom)

	melodyInstance.HandleMessage(handleMessage)

	log.Printf("Listening on port %s\n", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatalln("ListenAndServe failed")
	}
}
