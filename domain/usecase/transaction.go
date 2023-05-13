package usecase

import (
	"errors"

	"github.com/galang-dana/domain/input"
	"github.com/galang-dana/domain/model"
	"github.com/galang-dana/domain/repository"
	"github.com/galang-dana/utils"
)

type TransactionUseCase interface {
	GetTransactionByID(campaignID input.GetCampaignTransaction) ([]model.Transaction, error)
	GetTransactionByUserID(userID string) ([]model.Transaction, error)
	CreateTransaction(input input.CreateTransactionInput) (model.Transaction, error)
}

type transactionUseCase struct {
	TransactionRepo repository.TransactionRepository
	CampaignRepo    repository.CampaignRepository
	PaymentUseCase  PaymentUseCase
}

func NewTransactionUseCase(r repository.TransactionRepository, c repository.CampaignRepository, p PaymentUseCase) TransactionUseCase {
	return &transactionUseCase{r, c, p}
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

func (u *transactionUseCase) CreateTransaction(input input.CreateTransactionInput) (model.Transaction, error) {

	codeFormat := "newtrx" + input.User.Name
	transaction := model.Transaction{
		ID:         utils.GenerateId(),
		CampaignID: input.CampaignID,
		Amount:     input.Amount,
		Status:     "pending",
		UserID:     input.User.Id,
		Code:       codeFormat,
	}

	newTransaction, err := u.TransactionRepo.Save(transaction)
	if err != nil {
		return newTransaction, err
	}

	paymentURL, err := u.PaymentUseCase.GetPaymentURL(newTransaction, input.User)
	if err != nil {
		return newTransaction, err
	}

	newTransaction.PaymentURL = paymentURL
	newTransaction, err = u.TransactionRepo.Update(newTransaction)
	if err != nil {
		return newTransaction, err
	}

	return newTransaction, nil
}
