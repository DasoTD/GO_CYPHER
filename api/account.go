package api

import (
	"errors"
	"net/http"
	db "github.com/dasotd/gocypher/db/sqlc"
	"github.com/dasotd/gocypher/token"
	_ "github.com/dasotd/gocypher/util"
	"github.com/gin-gonic/gin"
	// "github.com/dasotd/gocypher/api"
)

type CreateAccountRequest struct {
	// Owner string `json:"owner" binding:"required,alphanum"`
	Currency string `json:"currency" binding:"required"`
}

type AddAccountBalanceRequest struct {
	Amount int64 `json:"amount"`
	ID     int64 `json:"id"`
}
type DeleteAccounRequest struct {
	ID     int64 `json:"id"`
}

func(server *Server) createAccount(ctx *gin.Context){

	var req CreateAccountRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	AuthPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)

	args := db.CreateAccountParams{
		Owner:    AuthPayload.Username,
		Balance:  0,
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
	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	if account.Owner != authPayload.Username {
		err := errors.New("account doesn't belong to the authenticated user")
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
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


type listAccountRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

func (server *Server) listAccounts(ctx *gin.Context) {
	var req listAccountRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	arg := db.ListAccountsParams{
		Owner:  authPayload.Username,
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	accounts, err := server.cypher.ListAccounts(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, accounts)
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