package api

import (
	// "fmt"
	"testing"
	"time"

	mockDB "github.com/dasotd/gocypher/db/mock"
	db "github.com/dasotd/gocypher/db/sqlc"
	"github.com/dasotd/gocypher/util"
	"github.com/golang/mock/gomock"
)


func TestGetAccountAPI(t *testing.T){
	account := RandomAccount()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cypher := mockDB.NewMockCypher(ctrl)

	cypher.EXPECT().
		GetAccount(gomock.Any(), gomock.Eq(account.ID)).
		Times(1).
		Return(account, nil)

	// server = NewServer(t, cypher)
}

func RandomAccount() db.Account {
	return db.Account{
		ID: util.RandomInt(1, 2000),
		Owner: util.RandomOwner(),
		Balance: util.RandomMoney(),
		Currency: util.RandomCurrency(),
		CreatedAt: time.Now(),
	}
}