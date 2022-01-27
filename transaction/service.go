package transaction

import (
	"bwa-project/campaign"
	"errors"
)

type Service interface {
	GetTransactionByCampaignId(input GetCampaignTransactionsInput) ([]Transaction, error)
	GetTransactionByUserId(userId int) ([]Transaction, error)
}

type service struct {
	repository         Repository
	campaignRepository campaign.Repository
}

func NewServices(repository Repository, campaignRepository campaign.Repository) service {
	return service{repository, campaignRepository}
}

func (s service) GetTransactionByCampaignId(input GetCampaignTransactionsInput) ([]Transaction, error) {
	campaign, err := s.campaignRepository.FindById(input.Id)
	if err != nil {
		return []Transaction{}, err
	}

	if campaign.UserId != input.User.ID {
		return []Transaction{}, errors.New("not an owner of the campaign")
	}

	transactions, err := s.repository.GetByCampaignId(input.Id)

	if err != nil {
		return transactions, err
	}

	return transactions, nil
}

func (s service) GetTransactionByUserId(userId int) ([]Transaction, error) {
	transactions, err := s.repository.GetByUserId(userId)
	if err != nil {
		return transactions, err
	}

	return transactions, nil
}
