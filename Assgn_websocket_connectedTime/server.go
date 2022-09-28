package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

// Global
var clients = make(map[*websocket.Conn]bool)
var broadcast = make(chan Message)
var sender = make(chan string)

var upgrader = websocket.Upgrader{}

type Clientstruct struct {
	ipaddress     string
	localAddress  string
	connectedTime string
}

var allClients = make([]Clientstruct, 0, 0)

type Message struct {
	Message string `json:"message"`
}

func HandleClients(w http.ResponseWriter, r *http.Request) {
	go broadcastMessagesToClients()

	websocket, err := upgrader.Upgrade(w, r, nil) // returning *websocket.Conn
	if err != nil {
		log.Fatal("error upgrading GET request to a websocket :: ", err)
	}

	defer websocket.Close()
	//Save all  Clients data base
	clients[websocket] = true

	cl := Clientstruct{
		websocket.RemoteAddr().String(),
		websocket.LocalAddr().String(),
		fmt.Sprintf("%s   ", time.Now()),
	}

	allClients = append(allClients, cl)

	// for i := range clients {
	// 	fmt.Println(i.RemoteAddr(), " ", i.LocalAddr(), "  ", websocket.RemoteAddr(), " ")
	// }
	for _, v := range allClients {
		fmt.Println(v.ipaddress, " ", v.localAddress, "  ", v.connectedTime, " ")
	}

	fmt.Println("******")

	for {
		var message Message

		err := websocket.ReadJSON(&message)
		if err != nil {
			log.Printf("error occurred while reading message : %v", err)
			delete(clients, websocket)
			break
		}
		broadcast <- message
		sender <- websocket.RemoteAddr().String()
	}
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})
	http.HandleFunc("/echo", HandleClients)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("error starting http server :: ", err)
		return
	}
}

func broadcastMessagesToClients() {
	for {
		message := <-broadcast
		senderAddress := <-sender
		for client := range clients {
			if client.RemoteAddr().String() == senderAddress {
				err := client.WriteJSON(message)
				if err != nil {
					log.Printf("error occurred while writing message to client: %v", err)
					client.Close()
					delete(clients, client)
				}
			}

		}
	}
}
