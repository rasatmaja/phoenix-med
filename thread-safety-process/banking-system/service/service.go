package service

import (
	"context"

	"github.com/rasatmaja/phoenix-med/banking-system/model"
)

// Repository ---
type Repository interface {
	BeginTx(ctx context.Context) (context.Context, error)
	CommitTx(ctx context.Context) error
	RollbackTx(ctx context.Context) error
	VerifyTX(ctx context.Context) error
	CreateUser(ctx context.Context, user model.User) error
	GetUserByID(ctx context.Context, id string) (model.User, error)
	UpdateUserBalance(ctx context.Context, user model.User) error
	CreateTransaction(ctx context.Context, transactions ...model.Transaction) error
	GetTransaction(ctx context.Context) ([]model.Transaction, error)
	LockAccount(ctx context.Context, userID string) error
	UnlockAccount(ctx context.Context, userID string) error
}

// Service ---
type Service struct{ repository Repository }

// NewService ---
func NewService(repository Repository) *Service {
	return &Service{repository: repository}
}
