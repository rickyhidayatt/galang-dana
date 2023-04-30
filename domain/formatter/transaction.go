package formatter

import (
	"time"

	"github.com/galang-dana/domain/model"
)

type CampaignTransactionFormater struct {
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Amount    int       `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
}

func FormatCampaignTransaction(transaction model.Transaction) CampaignTransactionFormater {
	formater := CampaignTransactionFormater{
		Id:        transaction.ID,
		Name:      transaction.User.Name,
		Amount:    transaction.Amount,
		CreatedAt: transaction.CreatedAt,
	}
	return formater
}

func FormatCampaignTransactions(transactions []model.Transaction) []CampaignTransactionFormater {

	if len(transactions) == 0 {
		return nil
	}

	var transactionsFormatter []CampaignTransactionFormater

	for _, data := range transactions {
		formatter := FormatCampaignTransaction(data)
		transactionsFormatter = append(transactionsFormatter, formatter)
	}

	return transactionsFormatter
}

type UserTransactionFormatter struct {
	Id        string         `json:"id"`
	Status    string         `json:"status"`
	Amount    int            `json:"amount"`
	CreatedAt time.Time      `json:"created_at"`
	Campaign  imageFormatter `json:"campaign"`
}

type imageFormatter struct {
	Name     string `json:"name"`
	ImageURL string `json:"image_url"`
}

func formatUserTransaction(transaction model.Transaction) UserTransactionFormatter {
	formater := UserTransactionFormatter{
		Id:        transaction.ID,
		Amount:    transaction.Amount,
		Status:    transaction.Status,
		CreatedAt: transaction.CreatedAt,
	}

	campaignFormatter := imageFormatter{}
	campaignFormatter.Name = transaction.Campaign.Name
	campaignFormatter.ImageURL = ""

	if len(transaction.Campaign.Images) > 0 {
		campaignFormatter.ImageURL = transaction.Campaign.Images[0].FileName
	}

	formater.Campaign = campaignFormatter
	return formater
}

func FormatUserTransactions(transactions []model.Transaction) []UserTransactionFormatter {
	if len(transactions) == 0 {
		return nil
	}

	var transactionFormat []UserTransactionFormatter

	for _, trx := range transactions {
		formatter := formatUserTransaction(trx)
		transactionFormat = append(transactionFormat, formatter)
	}
	return transactionFormat
}
