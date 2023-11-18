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

type AddAccountBalanceRequest struct {
	Amount int64 `json:"amount"`
	ID     int64 `json:"id"`
}
type DeleteAccounRequest struct {
	// Amount int64 `json:"amount"`
	ID     int64 `json:"id"`
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


func (server *Server) AddAccountBalance(ctx *gin.Context){
	var req AddAccountBalanceRequest
	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	args  := db.AddAccountBalanceParams {
		Amount: req.Amount,
		ID: req.ID,
	}
	_, err := server.cypher.GetAccount(ctx, req.ID);
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	account, err := server.cypher.AddAccountBalance(ctx, args);
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, account)
}


func (server *Server) DeleteAccount(ctx *gin.Context){
	var req DeleteAccounRequest
	if err := ctx.BindJSON(&req); err !=nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	err := server.cypher.DeleteAccount(ctx, req.ID);
	if err != nil { 
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, "Deleted successfully")
}