package repository

import (
	"context"
	"sync"
	"time"

	pgxtxpool "github.com/rasatmaja/pgx-txpool"
	"github.com/rasatmaja/phoenix-med/banking-system/model"
	"github.com/rasatmaja/phoenix-med/banking-system/utils"
)

// Repository --
type Repository struct {
	db   *pgxtxpool.Pool
	lock sync.Map
}

// NewRepository ---
func NewRepository(db *pgxtxpool.Pool) *Repository {
	return &Repository{db: db}
}

// BeginTx ---
func (r *Repository) BeginTx(ctx context.Context) (context.Context, error) {
	return r.db.BeginTX(ctx)
}

// CommitTx ---
func (r *Repository) CommitTx(ctx context.Context) error {
	time.Sleep(utils.RandomDuration(20, 200, time.Millisecond))
	return r.db.CommitTX(ctx)
}

// RollbackTx ---
func (r *Repository) RollbackTx(ctx context.Context) error {
	// this sleep is for simulate long running query/process
	time.Sleep(utils.RandomDuration(20, 200, time.Millisecond))
	return r.db.RollbackTX(ctx)
}

// VerifyTX --
func (r *Repository) VerifyTX(ctx context.Context) error {
	time.Sleep(utils.RandomDuration(20, 200, time.Millisecond))
	return r.db.VerifyTX(ctx)
}

// CreateUser ---
func (r *Repository) CreateUser(ctx context.Context, user model.User) error {

	time.Sleep(utils.RandomDuration(20, 200, time.Millisecond))

	query := `INSERT INTO users (id, name, balance) VALUES ($1, $2, $3)`
	_, err := r.db.Exec(ctx, query, user.ID, user.Name, user.Balance)
	if err != nil {
		return err
	}
	return nil
}

// GetUserByID --
func (r *Repository) GetUserByID(ctx context.Context, id string) (model.User, error) {
	time.Sleep(utils.RandomDuration(20, 200, time.Millisecond))

	query := `SELECT id, name, balance FROM users WHERE id = $1`
	row := r.db.QueryRow(ctx, query, id)

	var user model.User
	err := row.Scan(&user.ID, &user.Name, &user.Balance)
	if err != nil {
		return user, err
	}
	return user, nil
}

// UpdateUserBalance ---
func (r *Repository) UpdateUserBalance(ctx context.Context, user model.User) error {
	time.Sleep(utils.RandomDuration(20, 200, time.Millisecond))

	query := `UPDATE users SET balance = $1 WHERE id = $2`
	_, err := r.db.Exec(ctx, query, user.Balance, user.ID)
	if err != nil {
		return err
	}
	return nil
}

// CreateTransaction ---
func (r *Repository) CreateTransaction(ctx context.Context, transactions ...model.Transaction) error {

	time.Sleep(utils.RandomDuration(20, 200, time.Millisecond))

	query := `INSERT INTO transactions (id, user_id, type, amount) VALUES ($1, $2, $3, $4)`
	for _, transaction := range transactions {
		_, err := r.db.Exec(ctx, query, transaction.ID, transaction.UserID, transaction.Type, transaction.Amount)
		if err != nil {
			return err
		}
	}
	return nil
}

// GetTransaction ---
func (r *Repository) GetTransaction(ctx context.Context) ([]model.Transaction, error) {
	var transactions []model.Transaction
	query := `SELECT id, user_id, type, amount FROM transactions`
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return transactions, err
	}
	for rows.Next() {
		var transaction model.Transaction
		err := rows.Scan(&transaction.ID, &transaction.UserID, &transaction.Type, &transaction.Amount)
		if err != nil {
			return transactions, err
		}
		transactions = append(transactions, transaction)
	}
	return transactions, nil
}

// LockAccount --
func (r *Repository) LockAccount(ctx context.Context, userID string) error {

	_, ok := r.lock.Load(r.generateLockKey(userID))
	if ok {
		return model.ErrAccountLocked
	}
	r.lock.Store(r.generateLockKey(userID), true)
	return nil
}

// UnlockAccount --
func (r *Repository) UnlockAccount(ctx context.Context, userID string) error {
	r.lock.Delete(r.generateLockKey(userID))
	return nil
}

func (r *Repository) generateLockKey(userID string) string {
	return "lock_account:" + userID
}
