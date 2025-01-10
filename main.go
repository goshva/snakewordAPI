package main

import (
    "log"
    "net/http"
    "sync"

    "github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
    ReadBufferSize:  1024,
    WriteBufferSize: 1024,
    CheckOrigin: func(r *http.Request) bool {
        return true
    },
}

type User struct {
    Tgid  int `json:"tgid"`
    Score int `json:"score"`
}

var users = []User{
    {Tgid: 190404167, Score: 10000},
    {Tgid: 190404169, Score: 1000},
}

var usersMutex = &sync.Mutex{}

func main() {
    http.HandleFunc("/ws", handleConnections)
    log.Fatal(http.ListenAndServe(":6969", nil))
}

func handleConnections(w http.ResponseWriter, r *http.Request) {
    conn, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        log.Fatal(err)
    }
    defer conn.Close()

    for {
        var msg map[string]interface{}
        err := conn.ReadJSON(&msg)
        if err != nil {
            log.Printf("error: %v", err)
            break
        }

        action := msg["action"].(string)
        userId := int(msg["userId"].(float64))

        switch action {
        case "searchUser":
            handleSearchUser(conn, userId)
        case "createUser":
            handleCreateUser(conn, userId)
        default:
            log.Printf("unknown action: %s", action)
        }
    }
}

func handleSearchUser(conn *websocket.Conn, userId int) {
    usersMutex.Lock()
    defer usersMutex.Unlock()

    for _, user := range users {
        if user.Tgid == userId {
            err := conn.WriteJSON(user)
            if err != nil {
                log.Printf("error: %v", err)
            }
            return
        }
    }

    err := conn.WriteJSON(map[string]string{"error": "User not found"})
    if err != nil {
        log.Printf("error: %v", err)
    }
}

func handleCreateUser(conn *websocket.Conn, userId int) {
    usersMutex.Lock()
    defer usersMutex.Unlock()

    newUser := User{Tgid: userId, Score: 0}
    users = append(users, newUser)

    err := conn.WriteJSON(newUser)
    if err != nil {
        log.Printf("error: %v", err)
    }
}
