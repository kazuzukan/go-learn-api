package transaction

import "bwa-project/user"

type GetCampaignTransactionsInput struct {
	Id   int `uri:"id" binding:"required"`
	User user.User
}
