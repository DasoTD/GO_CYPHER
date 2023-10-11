package api

import (
	"net/http"

	db "github.com/dasotd/gocypher/db/sqlc"
	_ "github.com/dasotd/gocypher/util"
	"github.com/gin-gonic/gin"
	// "github.com/dasotd/gocypher/api"
)

type CreateAccountRequest struct {
	Owner string `json:"owner" binding:"required,alphanum"`
	Balance int64 `json:"balance" binding:"required,min=6"`
	Currency string `json:"currency" binding:"required"`
	// Lastname    string `json:"lastname" binding:"required"`
}

func(server *Server) createAccount(ctx *gin.Context){

	var req CreateAccountRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	// hashedPassword, err := util.HashPassword(req.Password)

	args := db.CreateAccountParams{
		Owner:    req.Owner,
		Balance:  req.Balance,
		Currency: req.Currency,
	}

	account, err := server.cypher.CreateAccount(ctx, args)
	if err !=nil {
		errCode := db.ErrorCode(err)
		if errCode == db.ForeignKeyViolation || errCode == db.UniqueViolation {
			ctx.JSON(http.StatusForbidden, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, account)
}