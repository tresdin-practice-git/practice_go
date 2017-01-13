package main

import (
  "fmt"
  "net/http"
  "github.com/gorilla/websocket"
)
var upgrader = websocket.Upgrader{
    ReadBufferSize:  1024,
    WriteBufferSize: 1024,
}
func handler (w http.ResponseWriter, r *http.Request) error {
    conn, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        fmt.Println(err)
        return err
    }
    messageType, p, err := conn.ReadMessage()
    fmt.Println(messageType, p, err)
    if err != nil {
        return err
    }
    if err = conn.WriteMessage(messageType, p); err != nil {
        return err
    }
    return nil
}
func main() {
    //db, err := sql.Open("mysql", "root:nhinCAIgi90$@/practice_go")
    // list, err := db.Query("SELECT * FROM users")
    // fmt.Println(list, err)
    http.Handle("/", http.FileServer(http.Dir("./")))
    http.HandleFunc("/ws", func (w http.ResponseWriter, r *http.Request) {
        handler(w, r)
    })
    http.ListenAndServe(":3000", nil)
}
