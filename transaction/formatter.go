package transaction

import "time"

type CampaignTransactionFormatter struct {
	Id        int       `json:"id"`
	Amount    int       `json:"amount"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

type UserTransactionFormatter struct {
	Id        int               `json:"id"`
	Amount    int               `json:"amount"`
	Status    string            `json:"status"`
	CreatedAt time.Time         `json:"created-at"`
	Campaign  CampaignFormatter `json:"campaign"`
}

type CampaignFormatter struct {
	Name     string `json:"name"`
	ImageUrl string `json:"image_url"`
}

func FormatCampaignTransction(transation Transaction) CampaignTransactionFormatter {
	formatter := CampaignTransactionFormatter{}
	formatter.Id = transation.Id
	formatter.Name = transation.User.Name
	formatter.Amount = transation.Amount
	formatter.CreatedAt = transation.CreatedAt

	return formatter
}

func FormatCampaignTransctions(transactions []Transaction) []CampaignTransactionFormatter {
	if len(transactions) == 0 {
		return []CampaignTransactionFormatter{}
	}

	var transactionFormatter []CampaignTransactionFormatter
	for _, transaction := range transactions {
		formatter := FormatCampaignTransction(transaction)
		transactionFormatter = append(transactionFormatter, formatter)
	}

	return transactionFormatter
}

func FormatUserTransaction(transaction Transaction) UserTransactionFormatter {
	formatter := UserTransactionFormatter{}
	formatter.Id = transaction.Id
	formatter.Amount = transaction.Amount
	formatter.Status = transaction.Status
	formatter.CreatedAt = transaction.CreatedAt

	campaignFormatter := CampaignFormatter{}
	campaignFormatter.Name = transaction.Campaign.Name
	campaignFormatter.ImageUrl = ""
	if len(transaction.Campaign.CampaignImages) > 0 {

		campaignFormatter.ImageUrl = transaction.Campaign.CampaignImages[0].FileName
	}

	formatter.Campaign = campaignFormatter

	return formatter
}

func FormatUserTransctions(transactions []Transaction) []UserTransactionFormatter {
	if len(transactions) == 0 {
		return []UserTransactionFormatter{}
	}

	var userTransactions []UserTransactionFormatter
	for _, transaction := range transactions {
		formatter := FormatUserTransaction(transaction)
		userTransactions = append(userTransactions, formatter)
	}

	return userTransactions
}
