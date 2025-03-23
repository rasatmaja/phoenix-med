package service

import (
	"context"

	"github.com/rasatmaja/phoenix-med/banking-system/model"
)

// CreateUser ---
func (s *Service) CreateUser(ctx context.Context, user model.User, trx ...model.Transaction) error {

	trxCTX, err := s.repository.BeginTx(ctx)
	if err != nil {
		return err
	}

	defer func(cause error) {
		if cause != nil {
			err = s.repository.RollbackTx(trxCTX)
		}
	}(err)

	if err = s.repository.CreateUser(trxCTX, user); err != nil {
		return err
	}

	if err = s.repository.CreateTransaction(trxCTX, trx...); err != nil {
		return err
	}

	if err = s.repository.CommitTx(trxCTX); err != nil {
		return err
	}

	return nil
}

// GetUserByID ---
func (s *Service) GetUserByID(ctx context.Context, id string) (model.User, error) {
	return s.repository.GetUserByID(ctx, id)
}
