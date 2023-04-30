package usecase

import (
	"errors"

	"github.com/galang-dana/domain/input"
	"github.com/galang-dana/domain/model"
	"github.com/galang-dana/domain/repository"
)

type TransactionUseCase interface {
	GetTransactionByID(campaignID input.GetCampaignTransaction) ([]model.Transaction, error)
	GetTransactionByUserID(userID string) ([]model.Transaction, error)
}

type transactionUseCase struct {
	TransactionRepo repository.TransactionRepository
	CampaignRepo    repository.CampaignRepository
}

func NewTransactionUseCase(r repository.TransactionRepository, c repository.CampaignRepository) TransactionUseCase {
	return &transactionUseCase{r, c}
}

func (u *transactionUseCase) GetTransactionByID(campaignID input.GetCampaignTransaction) ([]model.Transaction, error) {

	campaign, err := u.CampaignRepo.FindCampaignById(campaignID.ID)
	if err != nil {
		return []model.Transaction{}, err
	}

	if campaign.UserId != campaignID.User.Id {
		return []model.Transaction{}, errors.New("not an owner of the campaign")
	}

	transaction, err := u.TransactionRepo.GetByCampaignID(campaignID.ID)
	if err != nil {
		return transaction, err
	}

	return transaction, nil
}

func (u *transactionUseCase) GetTransactionByUserID(userID string) ([]model.Transaction, error) {
	transactions, err := u.TransactionRepo.GetByUserID(userID)
	if err != nil {
		return transactions, err
	}
	return transactions, nil
}
