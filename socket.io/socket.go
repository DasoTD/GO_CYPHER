package main

import (
	"fmt"
	"log"
	"net/http"

	// "strings"

	db "github.com/dasotd/gocypher/db/sqlc"
	"github.com/gin-gonic/gin"
	socketio "github.com/googollee/go-socket.io"
)

func Socket() *socketio.Server{
	server := socketio.NewServer(nil)
	return server
}



type CreateAccountRequest struct {
	Owner string `json:"owner" binding:"required,alphanum"`
	Balance int64 `json:"balance" binding:"required,min=6"`
	Currency string `json:"currency" binding:"required"`
	// Lastname    string `json:"lastname" binding:"required"`
}

func main() {

	server := Socket()
	
	

	// server.OnConnect("/", func(so socketio.Conn) error {
	// 	log.Println("on connection")

    //     so.Join("chat")
	// 	log.Printf("Client connected: %s\n", so.ID())
	// 	so.SetContext("")
	// 	return nil
	// })





	server.OnConnect("/", func(s socketio.Conn) error {
		s.SetContext("")
		fmt.Println("connected:", s.ID())
		s.Emit("chat", "oloba hava")
		server.OnEvent("/", "chat", func(s socketio.Conn, msg string){
			log.Printf("Message from client [%s]: %s\n", s.ID(), msg)
		})
		return nil
	})

	
	

	// server.OnEvent("/", "join", func(so socketio.Conn, room string) {
	// 	log.Printf("Client %s joined room %s\n", so.ID(), room)
	// 	so.Join(room)
	// 	so.SetContext(room)
	// })

	server.OnEvent("/", "chat", func(so socketio.Conn, msg string) {
		so.Join("chat")
		room := so.Context().(string)
		log.Printf("Received chat message '%s' from client %s in room %s\n", msg, so.ID(), room)
		server.BroadcastToRoom(room, "chat", so.ID()+": "+msg)
		server.BroadcastToNamespace("/", "chat", msg)
	})

	server.OnEvent("/", "storeUser", func(s socketio.Conn, data CreateAccountRequest){
		var req CreateAccountRequest
	var ctx *gin.Context
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	args := db.CreateAccountParams{
		Owner:    req.Owner,
		Balance:  req.Balance,
		Currency: req.Currency,
	}

	s.Emit("storeUser", args)
	})

	server.OnEvent("/", "chatmessage", func(s socketio.Conn, msg string) {
		log.Printf("Message from client [%s]: %s\n", s.ID(), msg)
		s.Emit("chatmessage", msg)
		// server.BroadcastToRoom("/", "chatroom", "chat message", msg)
	})

	// server.OnDisconnect("/", func(so socketio.Conn, reason string) {
	// 	room := so.Context().(string)
	// 	log.Printf("Client %s disconnected from room %s: %s\n", so.ID(), room, reason)
	// 	so.Leave(room)
	// })
	go server.Serve()

	defer server.Close()

	http.Handle("/socket.io/", server)
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)
	// http.Handle("/", http.FileServer(http.Dir("./static/index.html")))

	log.Println("Server started on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

	////CLIENT CONNECTION

	uri := "http://127.0.0.1:8080"

	client, err := socketio.NewClient(uri, nil)
	if err !=nil {
		log.Fatal(err)
	}

	// Handle an incoming event
	client.OnEvent("reply", func(s socketio.Conn, msg string) {
		log.Println("Receive Message /reply: ", "reply", msg)
	})

	client.Connect()
	client.Emit("notice", "hello")
	client.Close()
}
