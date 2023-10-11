package main

import (
	_ "context"
	"net/http"

	db "github.com/dasotd/gocypher/db/sqlc"
	"github.com/gin-gonic/gin"
	socketio "github.com/googollee/go-socket.io"

	"github.com/dasotd/gocypher/api"
)

type CreateAccountRequests struct {
	Owner string `json:"owner" binding:"required,alphanum"`
	Balance int64 `json:"balance" binding:"required,min=6"`
	Currency string `json:"currency" binding:"required"`
	// Lastname    string `json:"lastname" binding:"required"`
}

var server *api.Server

type MyConn struct {
    socketio.Conn
}

func StoreUser (s socketio.Conn) {
	var req CreateAccountRequests
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
	

	// user, err := db.Cypher.CreateAccount( ctx, args)



}