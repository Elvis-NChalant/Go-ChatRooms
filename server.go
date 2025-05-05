package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"sync"
	"math/rand"
)

var colours = []string{
    "\033[31m", // Red
    "\033[32m", // Green
    "\033[33m", // Yellow
    "\033[34m", // Blue
    "\033[35m", // Magenta
    "\033[36m", // Cyan
}

type Client struct {
		Conn     net.Conn
		Username string
		Colour string
}

var rooms = make(map[string][]Client)
var mu sync.Mutex

func handleConnection(conn net.Conn) {
	// Ensure the connection is closed when the function ends
	defer conn.Close()
	var roomID string
	var username string
	var client Client

	// Create a reader to read incoming messages from the client
	reader := bufio.NewReader(conn)
	message, _ := reader.ReadString('\n')
	message = strings.TrimSpace(message)

	if strings.Contains(message, "joined room") {
		parts := strings.Split(message, "joined room")
		if len(parts) == 2 {
			username = strings.TrimSpace(parts[0])
			roomID = strings.TrimSpace(parts[1])
			colour := colours[rand.Intn(len(colours))]
			
			client = Client{Conn: conn, Username: username, Colour: colour}
			mu.Lock()
			rooms[roomID] = append(rooms[roomID], client)
			mu.Unlock()
	
			fmt.Printf(username + " joined room %s\n", roomID)
		}
	}

	for {
		// Read the message sent by the client
		message, err := reader.ReadString('\n')
		message = strings.TrimSpace(message)
		
		if err != nil {
			log.Println("Error reading message:", err)
			return
		}

		if message == "show table"{
			conn.Write([]byte("Current Room Map:\n"))
			mu.Lock()
			for roomID, clients := range rooms {
				msg := fmt.Sprintf("Room %s:\n", roomID)
				conn.Write([]byte(msg))
				for i, client := range clients {
					//msg := fmt.Sprintf("  %d. %s (%s)\n", i+1, client.Username, client.Conn.RemoteAddr())
					msg := fmt.Sprintf("%d. \033[%s%s\033[0m:    (%s)\n", i+1, client.Colour, client.Username,  client.Conn.RemoteAddr())
					conn.Write([]byte(msg))
				}
			}
			mu.Unlock()
		}else{
			// Print the received message
			mu.Lock()
			broadcast := fmt.Sprintf("%s%s\033[0m: %s\n", client.Colour, client.Username, message)
			for _, client := range rooms[roomID] {
				if client.Conn != conn { // don't send to the sender
					_, err := client.Conn.Write([]byte(broadcast))
					if err != nil {
						log.Printf("Error sending to %s: %v\n", client.Username, err)
					}
				}
			}
			mu.Unlock()
		}
		

	}
}

func main() {
	// Listen on port 8080
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Println("Error starting server:", err)
		return
	}
	defer listener.Close()

	fmt.Println("Server is listening on :8080...")

	// Accept incoming connections
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Error accepting connection:", err)
			continue
		}

		// Handle each connection in a separate goroutine
		go handleConnection(conn)
	}
}

// NEED TO DELETE CONNECTION FROM ROOMS MAP WHEN CLIENT DISCONNECTS