<h1 align="center"><samp>Hermes 💬</samp></h1>

<p align="center">
  <img src="https://img.shields.io/badge/Go-1.27-00ADD8?logo=go&logoColor=white" alt="Go">
  <img src="https://img.shields.io/badge/TCP-:8090-black" alt="TCP 8090">
</p>

<p align="center"><samp>A minimalist TCP chat server in Go 🐹</samp></p>

<br>

<p><samp>Requirements</samp></p>

- [Go](https://go.dev/dl/) 1.27 or newer
- `nc` (netcat) to connect, or any TCP client

<br>

<p><samp>Start</samp></p>

```sh
go run ./cmd/server
```

<br>

<p><samp>Connect</samp></p>

```sh
nc localhost 8090
```

<br>

<p><samp>Usage</samp></p>

- Choose a username
- Type a message
- Everyone receives it
- `/leave` to quit

<br>

<p><samp>Structure</samp></p>

```
cmd/
└── server/
    └── main.go      starts the server, accepts connections
internal/
├── design/
│   └── design.go    ASCII banner shown on connection
├── identity/
│   ├── identity.go  asks for a username and creates the user
│   └── user.go      user type and list of connected users
└── handlers/
    └── handle.go    reads messages and sends them to others
```
