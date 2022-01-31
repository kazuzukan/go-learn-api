package transaction

import "bwa-project/user"

type GetCampaignTransactionsInput struct {
	Id   int `uri:"id" binding:"required"`
	User user.User
}

type CreateTransactionInput struct {
	Amount     int       `json:"amount" binding:"required"`
	CampaignId int       `json:"campaign_id" binding:"required"`
	User       user.User `json:"user"`
}

type TransactionNotificationInput struct {
	TransactionStatus string `json:"transaction_status"`
	OrderId           string `json:"order_id"`
	PaymentType       string `json:"payment_type"`
	FraudStatus       string `json:"fraud_status"`
}
