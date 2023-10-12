package api

import (
	"github.com/gin-gonic/gin"
	socketio "github.com/googollee/go-socket.io"
)

func (server *Server) setupRouter() {
	router := gin.Default()

	router.POST("/users", server.createUser)
	router.POST("/users/login", server.loginUser)
	// router.POST("/tokens/renew_access", server.renewAccessToken)

	// authRoutes := router.Group("/").Use(authMiddleware(server.tokenMaker))
	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts", server.listAccounts)

	router.POST("/transfers", server.createTransfer)
	socket := socketio.NewServer(nil)
	router.GET("/socket.io/", gin.WrapH(socket))

	server.router = router
}


// func (server *Server) setupRouter() {
	// 		router  := gin.Default()
	// 		socket := socketio.NewServer(nil)
	// 		router.GET("/socket.io/", gin.WrapH(socket))
	// 		socket.OnConnect("/", func(so socketio.Conn) error {
	// 			log.Printf("Client connected: %s\n", so.ID())
	// 			so.SetContext("")
	// 			return nil
	// 		})
	// 		socket.OnEvent("/", "chat message", func(s socketio.Conn, msg string) {
	// 			log.Printf("Message from client [%s]: %s\n", s.ID(), msg)
	// 			socket.BroadcastToRoom("/", "chatroom", "chat message", msg)
	// 		})
		
	// 		socket.OnDisconnect("/", func(s socketio.Conn, reason string) {
	// 			log.Println("Client disconnected:", s.ID(), reason)
	// 		})
		


	// 		server.setupRouter()
	// }


