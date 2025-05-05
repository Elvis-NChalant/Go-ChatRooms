# Go TCP Chat Application

A simple TCP-based chat application written in Go, supporting multiple rooms, user color-coding, and client-server architecture. Designed for learning and experimenting with socket programming, concurrency, and basic authentication concepts.

## Features

- Room-based chat support
- Unique color for each client
- Server handles multiple clients concurrently
- Server hosted on AWS EC2 and tested for real users across the globe
- Client executable can be compiled for different platforms (Windows, macOS, Linux)

## Architecture

- **Server:** Listens for client connections, handles room logic, and broadcasts messages.
- **Client:** Connects to the server, sends/receives messages, and displays chat in the terminal.


## Setup

### 1. Clone the Repository

```bash
git clone https://github.com/your-username/go-tcp-chat.git
cd go-tcp-chat
