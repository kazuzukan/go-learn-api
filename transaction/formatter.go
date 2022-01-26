package transaction

import "time"

type CampaignTransactionFormatter struct {
	Id        int       `json:"id"`
	Amount    int       `json:"amount"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
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
