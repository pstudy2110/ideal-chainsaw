package handlers

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

var Videodata []byte

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// Store all connected clients
var clients = make(map[*websocket.Conn]bool)
var clientsMutex sync.Mutex

func WebSocketSend(w http.ResponseWriter, r *http.Request, a string) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()
	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		fmt.Printf("recv: %s", message)
		err = c.WriteMessage(mt, ([]byte(a)))
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}

func WebSocketSendEcho(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()
	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		fmt.Printf("\n recv: %s", string(message))
		a := fmt.Sprintf("return:  %v", string(message))

		err = c.WriteMessage(mt, ([]byte(a)))
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}

func WebSocketSendPing(w http.ResponseWriter, r *http.Request) { //echo message for production test
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()
	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		fmt.Printf("\n recv: %s", string(message))
		//a := fmt.Sprintf("return:  %v", string(message)) //a := fmt.Sprintf("return:  %v", string(message))

		executeFunction := func() {
			a := fmt.Sprintf("ping:  %v", time.Now())
			err = c.WriteMessage(mt, ([]byte(a))) // swap a for message for now
			if err != nil {
				log.Println("write:", err)
			}

		}
		// Set the interval to 3 seconds
		interval := 3 * time.Second
		// Create a ticker that fires every 10 minutes
		ticker := time.NewTicker(interval)
		// Run the function initially
		executeFunction()
		// Use a for loop to execute the function every time the ticker fires
		for range ticker.C {
			executeFunction()
		}

	}
}

func RegularWebsocket(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()
	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		log.Printf("recv: %s", message)
		a := fmt.Sprintln("the messaGE")
		err = c.WriteMessage(mt, ([]byte(a)))
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}

func StoreAllWebsocket(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()
	// Add the new client to the list of connected clients
	clientsMutex.Lock()
	clients[c] = true
	clientsMutex.Unlock()

	defer func() {
		clientsMutex.Lock()
		delete(clients, c)
		clientsMutex.Unlock()
	}()

	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		log.Printf("recv: %s", string(message))

		a := fmt.Sprintf("%v", string(message))
		clientsMutex.Lock()
		for client := range clients {
			err := client.WriteMessage(mt, []byte(a))
			if err != nil {
				log.Println("write:", err)
				client.Close()
				delete(clients, client)
			}
		}
		clientsMutex.Unlock()
	}
}
