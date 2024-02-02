package repository

import (
	"context"

	"github.com/mesxx/Echo_Simple_Self_Payroll_API/config"
	"github.com/mesxx/Echo_Simple_Self_Payroll_API/model"
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
