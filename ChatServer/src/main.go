package main

import (
	"log"
	"net/http"
	"github.com/gorilla/websocket"
)

//global variables
var clients = make(map[*websocket.Conn]bool) //boolean function for connected clients
var broadcast = make(chan Message) //broadcast channel 

// Configure the upgrader
// This is an object with methods for taking a normal HTTP connection and upgrading it to WebSocket 
var upgrader = websocket.Upgrader{}

//object to hold messages
type Message struct {
	Email string `json:"email"`
	Username string `json:"username"`
	Message string `json:"message"`
}
func handleConnections(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	// ensure connection close when function returns
	defer ws.Close()
	clients[ws] = true

	for {
		var msg Message
		// Read in a new message as JSON and map it to a Message object
		err := ws.ReadJSON(&msg)
		if err != nil {
			log.Printf("error: %v", err)
			delete(clients, ws)
			break
		}
		// send the new message to the broadcast channel
		broadcast <- msg
	}
}
/*
func handleConnections(w http.ResponseWriter, r *http.Request) {
        // Upgrade initial GET request to a websocket
        ws, err := upgrader.Upgrade(w, r, nil)
        if err != nil {
                log.Fatal(err)
        }
        // Make sure we close the connection when the function returns
        defer ws.Close()

	clients[ws] = true //add to global clients map

	for {
                var msg Message
                // Read in a new message as JSON and map it to a Message object
                err := ws.ReadJSON(&msg)
                if err != nil {
                        log.Printf("error: %v", err)
                        delete(clients, ws)
                        break
                }
                // Send the newly received message to the broadcast channel
                broadcast <- msg

		// Grab the next message from the broadcast channel
               // msg := <-broadcast
        }
}
*/
func handleMessages() {
        for {
                // Grab the next message from the broadcast channel
                msg := <-broadcast
                // Send it out to every client that is currently connected
                for client := range clients {
                        err := client.WriteJSON(msg)
                        if err != nil {
                                log.Printf("error: %v", err)
                                client.Close()
                                delete(clients, client)
                        }
                }
        }
}

func main() {
	//make static fileserver to let users access public assets
	fs := http.FileServer(http.Dir("../public"))
	http.Handle("/", fs)

	//configure websocket route to handle WebSocket requests 
	http.HandleFunc("/ws", handleConnections)
	//run goroutine
	//concurrent process that only takes messages from broadcast channel
	go handleMessages()

	// Start the server on localhost port 8000 and log any errors
        log.Println("http server started on :8000")
        err := http.ListenAndServe(":8000", nil)
        if err != nil {
                log.Fatal("ListenAndServe: ", err)
        }

}

