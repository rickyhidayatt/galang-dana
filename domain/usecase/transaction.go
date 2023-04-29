package usecase

import (
	"github.com/galang-dana/domain/input"
	"github.com/galang-dana/domain/model"
	"github.com/galang-dana/domain/repository"
)

type TransactionUseCase interface {
	GetTransactionByID(campaignID input.GetCampaignTransaction) ([]model.Transaction, error)
}

type transactionUseCase struct {
	TransactionRepo repository.TransactionRepository
}

func NewTransactionUseCase(r repository.TransactionRepository) TransactionUseCase {
	return &transactionUseCase{r}
}

func (u *transactionUseCase) GetTransactionByID(campaignID input.GetCampaignTransaction) ([]model.Transaction, error) {
	transaction, err := u.TransactionRepo.GetByCampaignID(campaignID.ID)
	if err != nil {
		return transaction, err
	}

	return transaction, nil

}
