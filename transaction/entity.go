package transaction

import (
	"bwa-project/campaign"
	"bwa-project/user"
	"time"
)

type Transaction struct {
	Id         int
	CampaignId int
	UserId     int
	Amount     int
	Status     string
	Code       string
	PaymentUrl string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	User       user.User
	Campaign   campaign.Campaign
}
