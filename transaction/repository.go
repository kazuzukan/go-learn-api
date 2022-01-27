package transaction

import "gorm.io/gorm"

type Repository interface {
	GetByCampaignId(campaignId int) ([]Transaction, error)
	GetByUserId(userId int) ([]Transaction, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) repository {
	return repository{db}
}

func (r repository) GetByCampaignId(campaignId int) ([]Transaction, error) {
	var transaction []Transaction

	err := r.db.Preload("User").Where("campaign_id = ?", campaignId).Order("created_at DESC").Find(&transaction).Error
	if err != nil {
		return transaction, err
	}

	return transaction, nil
}

func (r repository) GetByUserId(userId int) ([]Transaction, error) {
	var transactions []Transaction
	// get nested key relations
	err := r.db.Preload("Campaign.CampaignImages", "campaign_images.is_primary = 1").Preload("User").Where("user_id = ?", userId).Order("created_at DESC").Find(&transactions).Error
	if err != nil {
		return transactions, err
	}

	return transactions, nil
}
