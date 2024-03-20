package main

import "log"

func handleMessages() {
	for {
		msg := <-broadcast
		var outbound OutboundMessage
		outbound.Content = msg.ChatMessage
		log.Println("Recieved message")

		for client := range clients {
			err := client.WriteJSON(outbound)
			if err != nil {
				log.Printf("Error writing JSON: %s\n", err.Error())
				client.Close()
				delete(clients, client)
			}
		}
	}
}
