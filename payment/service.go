package payment

import (
	"bwa-project/campaign"
	"bwa-project/user"
	"os"
	"strconv"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
)

type Service interface {
	GetPaymentUrl(transaction Transaction, user user.User) (string, error)
}

type service struct {
	campaignRepository campaign.Repository
}

func NewService(campaignRepository campaign.Repository) service {
	return service{campaignRepository}
}

func (ser service) GetPaymentUrl(transaction Transaction, user user.User) (string, error) {
	// 1. Initiate Snap client
	var s = snap.Client{}
	s.New(os.Getenv("MIDTRANS_SERVER_KEY"), midtrans.Sandbox)

	// 2. Initiate Snap request
	req := &snap.Request{
		CustomerDetail: &midtrans.CustomerDetails{
			Email: user.Email,
			FName: user.Name,
		},
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  strconv.Itoa(transaction.Id),
			GrossAmt: int64(transaction.Amount),
		},
	}

	// 3. Request create Snap transaction to Midtrans
	snapResp, err := s.CreateTransactionUrl(req)
	if err != nil {
		return "", err.RawError
	}
	return snapResp, nil
}
