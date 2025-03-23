package service

import (
	"context"

	"github.com/rasatmaja/phoenix-med/banking-system/model"
)

// CreateTransaction ---
func (s *Service) CreateTransaction(ctx context.Context, data model.Transaction) error {

	err := s.repository.LockAccount(ctx, data.UserID)
	if err != nil {
		return err
	}

	defer s.repository.UnlockAccount(ctx, data.UserID)

	trxCTX, err := s.repository.BeginTx(ctx)
	if err != nil {
		return err
	}
	defer func(cause error) {
		if cause != nil {
			err = s.repository.RollbackTx(trxCTX)
		}
	}(err)

	user, err := s.repository.GetUserByID(trxCTX, data.UserID)
	if err != nil {
		return err
	}

	if data.Type == model.TransactionTypeWithdraw {
		if user.Balance < data.Amount {
			return model.ErrInsufficientBalance
		}
		user.Balance -= data.Amount
	} else if data.Type == model.TransactionTypeDeposit {
		if data.Amount <= 0 {
			return model.ErrInvalidTransactionAmount
		}
		user.Balance += data.Amount
	} else {
		return model.ErrInvalidTransactionType
	}

	err = s.repository.UpdateUserBalance(trxCTX, user)
	if err != nil {
		return err
	}

	err = s.repository.CreateTransaction(trxCTX, data)
	if err != nil {
		return err
	}
	err = s.repository.CommitTx(trxCTX)
	if err != nil {
		return err
	}
	return nil
}
