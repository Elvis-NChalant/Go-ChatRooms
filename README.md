# Go TCP Chat Application

A simple TCP-based chat application written in Go, supporting multiple rooms, user color-coding, and client-server architecture. Designed for learning and experimenting with socket programming, concurrency, and basic authentication concepts.

## Idea

- Uses persistent TCP to connect with the server
- Server keeps track of all the connections using a hashmap
- When a new client joins he is prompted to either create or join a room
- Every room id is mapped to a connection (user connection) in the server using the hashmap
- Whenever a client sends a message in the room, the server checks the room id of the client and looks at the hashmap to check all other members present in the same room
- The server sends the message received by the client to all the members mapped by the same room id except the sender

## Features

- Room-based chat support
- Unique color for each client
- Server handles multiple clients concurrently
- Server hosted on AWS EC2 and tested for real users across the globe
- Client executable can be compiled for different platforms (Windows, macOS, Linux)

## Architecture

- **Server:** Listens for client connections, handles room logic, and broadcasts messages.
- **Client:** Connects to the server, sends/receives messages, and displays chat in the terminal.



![Logo](https://github.com/Elvis-NChalant/Go-ChatRooms/blob/main/chatarch.png)



## Setup

### 1. Clone the Repository

```bash
git clone https://github.com/your-username/go-tcp-chat.git
cd go-tcp-chat
