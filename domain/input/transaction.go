package input

import "github.com/galang-dana/domain/model"

type GetCampaignTransaction struct {
	ID   string `uri:"id" binding:"required"`
	User model.User
}

type CreateTransactionInput struct {
	Amount     int        `json:"amount" binding:"required"`
	CampaignID string     `json:"campaign_id" binding:"required"`
	User       model.User `json:"user" binding:"required"`
}

// terima data dari midtrans
type TransactionNotificationInput struct {
	TransactionStatus string `json:"transaction_status"`
	OrderID           string `json:"order_id"`
	PaymentType       string `json:"payment_type"`
	FraudStatus       string `json:"fraud_status"`
}
