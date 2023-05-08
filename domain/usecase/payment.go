package usecase

import (
	"github.com/galang-dana/domain/input"
	"github.com/galang-dana/domain/model"
	"github.com/galang-dana/domain/repository"
	"github.com/veritrans/go-midtrans"
)

type PaymentUseCase interface {
	GetPaymentURL(transaction model.Transaction, user model.User) (string, error)
	ProcessPayment(input input.TransactionNotificationInput) error
}
type paymentUsecase struct {
	transactionRepo repository.TransactionRepository
	campaignRepo    repository.CampaignRepository
}

func NewPaymentUseCase(t repository.TransactionRepository, c repository.CampaignRepository) *paymentUsecase {
	return &paymentUsecase{t, c}
}

func (u *paymentUsecase) GetPaymentURL(transaction model.Transaction, user model.User) (string, error) {
	midclient := midtrans.NewClient()
	midclient.ServerKey = "AMBIL DARI MIDTRANS" //server key di acount midtrans
	midclient.ClientKey = "AMBIL DARI MIDTRANS" //clirny key di acount midtrans
	midclient.APIEnvType = midtrans.Sandbox

	var snapGateway midtrans.SnapGateway
	snapGateway = midtrans.SnapGateway{
		Client: midclient,
	}

	snapReq := &midtrans.SnapReq{
		CustomerDetail: &midtrans.CustDetail{
			Email: user.Email,
			FName: user.Name,
		},
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  transaction.ID,
			GrossAmt: int64(transaction.Amount),
		},
	}
	snapTokenResp, err := snapGateway.GetToken(snapReq)
	if err != nil {
		return "", err
	}
	return snapTokenResp.RedirectURL, nil

}

func (u *paymentUsecase) ProcessPayment(input input.TransactionNotificationInput) error {
	transaction, err := u.transactionRepo.GetByID(input.OrderID)
	if err != nil {
		return err
	}

	if input.PaymentType == "credit_card" && input.TransactionStatus == "capture" && input.FraudStatus == "accept" {
		transaction.Status = "paid"
	} else if input.TransactionStatus == "settelment" {
		transaction.Status = "paid"
	} else if input.TransactionStatus == "deny" || input.TransactionStatus == "expire" || input.TransactionStatus == "cancel" {
		transaction.Status = "cancelled"
	}

	updatedTransaction, err := u.transactionRepo.Update(transaction)
	if err != nil {
		return err
	}

	campaign, err := u.campaignRepo.FindCampaignById(updatedTransaction.CampaignID)
	if err != nil {
		return err
	}

	if updatedTransaction.Status == "paid" {
		campaign.BackerCount = campaign.BackerCount + 1
		campaign.CurrentAmount = campaign.CurrentAmount + updatedTransaction.Amount
		_, err := u.campaignRepo.Update(campaign)
		if err != nil {
			return err
		}
	}
	return nil
}
