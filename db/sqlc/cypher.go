package db

import (
	// "context"
	// "database/sql"

	"github.com/jackc/pgx/v5/pgxpool"
	// "fmt"
)

// Cypher defines all functions to execute db queries and transactions
type Cypher interface {
	Querier
	// TransferTx(ctx context.Context, arg TransferTxParams) (TransferTxResult, error)
	// CreateUserTx(ctx context.Context, arg CreateUserTxParams) (CreateUserTxResult, error)
	// VerifyEmailTx(ctx context.Context, arg VerifyEmailTxParams) (VerifyEmailTxResult, error)
}

// SQLCypher provides all functions to execute SQL queries and transactions
type SQLCypher struct {
	connPool *pgxpool.Pool
	*Queries
}

// NewCypher creates a new Cypher
func NewCypher(pool *pgxpool.Pool) Cypher {
	return &SQLCypher{
		connPool:      pool,
		Queries: New(pool),
	}
}

// ExecTx executes a function within a database transaction
// func (Cypher *SQLCypher) execTx(ctx context.Context, fn func(*Queries) error) error {
// 	tx, err := Cypher.db.BeginTx(ctx, nil)
// 	if err != nil {
// 		return err
// 	}

// 	q := New(tx)
// 	err = fn(q)
// 	if err != nil {
// 		if rbErr := tx.Rollback(); rbErr != nil {
// 			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
// 		}
// 		return err
// 	}

// 	return tx.Commit()
// }