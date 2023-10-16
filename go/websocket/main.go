package main

import (
	"log"
	"net/http"
	"time"
	"encoding/json"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		// Allow connections from any origin
		return true
	},
}

type SecondsSinceConnection struct {
	Seconds int `json:"seconds"`
}

func main() {
	// Handle WebSocket connections on the /ws endpoint
	http.HandleFunc("/ws", handleWebSocket)

	// Start the HTTP server
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	// Upgrade the HTTP connection to a WebSocket connection
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	// Send a message to the client every second
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	startTime := time.Now()

	for {
		select {
		case <-ticker.C:
			// Calculate the seconds since the WebSocket connection was established
			seconds := int(time.Since(startTime) / time.Second)
	
			// Create a SecondsSinceConnection struct and set the seconds field
			ssc := SecondsSinceConnection{Seconds: seconds}
	
			// Convert the SecondsSinceConnection struct to a JSON byte slice
			jsonBytes, err := json.Marshal(ssc)
			if err != nil {
				log.Println(err)
				return
			}
	
			// Send the JSON byte slice to the client
			err = conn.WriteMessage(websocket.TextMessage, jsonBytes)
			if err != nil {
				log.Println(err)
				return
			}
		default:
			// Read messages from the WebSocket connection
			_, message, err := conn.ReadMessage()
			if err != nil {
				log.Println(err)
				return
			}
			log.Printf("Received message: %s\n", message)
	
			// Check if the message is a request for the seconds since connection
			if string(message) == "refresh" {
				continue
			}
	
			// Send a response message back to the client
			response := []byte(string(message))
			err = conn.WriteMessage(websocket.TextMessage, response)
			if err != nil {
				log.Println(err)
				return
			}
		}
	}
}