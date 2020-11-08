package main

import (
  "fmt"
  "log"
  "net/http"

  "github.com/gorilla/websocket"
)

// Define upgrader, which require a Read and Write buffer size
var upgrader = websocket.Upgrader{
  ReadBufferSize: 1024,
  WriteBufferSize: 1024,

  // Check connections'origin
  // Allowing to make requests from React development server to here.

  // Allow any connections
  CheckOrigin: func(r *http.Request) bool { return true },
}

// define reader which will listen for new messages being sent to our WebSocket endpoint
func reader(conn *websocket.Conn) {
  for {
    // read in a message
    messageType, p, err := conn.ReadMessage()
    if err != nil {
      log.Println(err)
      return
    }

    // print out that message for clarity
    fmt.Println(string(p))

    if err := conn.WriteMessage(messageType, p); err != nil {
      log.Println(err)
      return
    }
  }
}

// define our WebSocket endpoint
func serveWs(w http.ResponseWriter, r *http.Request) {
  fmt.Println(r.Host)
  // upgrade this connection to a Websocket connection
  ws, err := upgrader.Upgrade(w, r, nil)
  if err != nil {
    log.Println(err)
  }
  // listen indefinetly for new messages coming through on our Websocket connection
  reader(ws)
}

func setupRoutes() {
  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Simple Server")
  })
  // mape `/ws` endpoint to the `serveWs` function
  http.HandleFunc("/ws", serveWs)
}

func main() {
  setupRoutes()
  http.ListenAndServe(":8080", nil)
}
