package repository

import (
	"context"
	"echo_simple_self_payroll_api/config"
	"echo_simple_self_payroll_api/model"
)

type transactionRepository struct {
	Cfg config.Config
}

func NewTransactionRepository(cfg config.Config) model.TransactionRepository {
	return &transactionRepository{Cfg: cfg}
}

func (t *transactionRepository) Fetch(ctx context.Context, limit, offset int) ([]*model.Transaction, error) {
	var data []*model.Transaction

	if err := t.Cfg.Database().WithContext(ctx).
		Limit(limit).Offset(offset).Find(&data).Error; err != nil {
		return nil, err
	}

	return data, nil
}
