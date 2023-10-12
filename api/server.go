package api

import (
	"fmt"
	"io"
	"log"
	"os"
	_ "time"

	db "github.com/dasotd/gocypher/db/sqlc"
	"github.com/dasotd/gocypher/token"
	"github.com/dasotd/gocypher/util"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	socketio "github.com/googollee/go-socket.io"
)

type Server struct {
	config     util.Config
	cypher db.Cypher
	tokenMaker token.Maker
	router *gin.Engine
}

func NewServer(config util.Config, cypher db.Cypher) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}
	server := &Server{
		config: config,
		cypher: cypher,
		tokenMaker: tokenMaker,
	}

	 // Logging to a file.
	 f, _ := os.Create("gin.log")
	 gin.DefaultWriter = io.MultiWriter(f)

	 
	

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", validCurrency)
	}

	server.setupRouter()
	return server, nil


	// router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {

	// 	// your custom format
	// 	return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
	// 		param.ClientIP,
	// 		param.TimeStamp.Format(time.RFC1123),
	// 		param.Method,
	// 		param.Path,
	// 		param.Request.Proto,
	// 		param.StatusCode,
	// 		param.Latency,
	// 		param.Request.UserAgent(),
	// 		param.ErrorMessage,
	// 	)
	//   }))

	


	


	
	}

	

func errorResponse(err error) gin.H {
		return gin.H{"error": err.Error()}
	}

	// // // Start runs the HTTP server on a specific address.
func (server *Server) Start(address string) error {
	log.Println(`Server started on :8080`)
	return server.router.Run(address)
}

func Socket(){
	fmt.Println("socket...")
	socket := socketio.NewServer(nil)
	socket.OnConnect("/", func(so socketio.Conn) error {
		fmt.Printf("Client connected: %s\n", so.ID())
		so.SetContext("")
		return nil
	})
}