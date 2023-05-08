package usecase

import (
	"github.com/galang-dana/domain/model"
	"github.com/veritrans/go-midtrans"
)

type PaymentUseCase interface {
	GetPaymentURL(transaction model.Transaction, user model.User) (string, error)
}
type paymentUsecase struct {
}

func NewPaymentUseCase() *paymentUsecase {
	return &paymentUsecase{}
}

func (u *paymentUsecase) GetPaymentURL(transaction model.Transaction, user model.User) (string, error) {
	midclient := midtrans.NewClient()
	midclient.ServerKey = "ambil di midtrans key nya" //server key di acount midtrans
	midclient.ClientKey = "ambil di midtrans key nya" //clirny key di acount midtrans
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
