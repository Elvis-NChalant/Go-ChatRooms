package main

import (
    "fmt"
    "net"
    "bufio"
    "os"
    "strings"
    "math/rand"
)

func generateRoomID() string {
    //rand.Seed(time.Now().UnixNano())
    roomID := fmt.Sprintf("%d", rand.Intn(1000000)) // Generate a random room ID
    return roomID
}

func gracefulExit(conn net.Conn, username string) {
    fmt.Println("Exiting chat...")
    conn.Write([]byte(username + " has left the chat.\n"))
}

func main() {
    serverAddr := os.Getenv("SERVER_ADDR")
    conn, err := net.Dial("tcp", serverAddr)
    
    if err != nil {
        fmt.Println("Server is not online")
        return
    }
    defer conn.Close()

    fmt.Println("Welcome to the chat!")
    
    // Ask the user for a username
    fmt.Print("Enter your username: ")
    reader := bufio.NewReader(os.Stdin)
    username, _ := reader.ReadString('\n')
    username = strings.TrimSpace(username)

    // Ask the user if they want to create or join a room
    fmt.Print("Do you want to create a new room or join an existing room? (create/join): ")
    action, _ := reader.ReadString('\n')
    action = strings.TrimSpace(action)

    var roomID string

    if action == "create" || action == "c" {
        // Generate a unique room ID
        roomID = generateRoomID()
        fmt.Printf("Room created! Your room ID is: %s\n", roomID)
        
    } else if action == "join" || action == "j" {
        // Ask the user for the room ID to join
        fmt.Print("Enter the room ID to join: ")
        roomID, _ = reader.ReadString('\n')
        roomID = strings.TrimSpace(roomID)
    } else {
        fmt.Println("Invalid option, exiting.")
        return
    }

    // Send the username and room ID to the server
    conn.Write([]byte(username + " joined room " + roomID + "\n"))

    // Send messages to the server and receive responses
    go func() {
        serverReader := bufio.NewReader(conn)
        for {
            message, err := serverReader.ReadString('\n')
            if err != nil {
                fmt.Println("Disconnected from server.")
                return
            }
            fmt.Print("\r\033[2K")

            fmt.Print(message)

            fmt.Print("> ") 
        }
    }()

    // Loop to send messages to the server
    for {
        
        fmt.Print("> ")
        message, _ := reader.ReadString('\n')
        message = strings.TrimSpace(message)
        if message == "exit" {
            gracefulExit(conn, username)
            return
        }else if message == "show table" {
            conn.Write([]byte("show table\n"))
        }else{
            conn.Write([]byte(message + "\n"))
        }
    }
}
