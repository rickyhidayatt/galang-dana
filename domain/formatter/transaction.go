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
