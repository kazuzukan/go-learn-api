package transaction

import "gorm.io/gorm"

type Repository interface {
	GetByCampaignId(campaignId int) ([]Transaction, error)
	GetByUserId(userId int) ([]Transaction, error)
	Save(transaction Transaction) (Transaction, error)
	Update(transaction Transaction) (Transaction, error)
	GetById(Id int) (Transaction, error)
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

func (r repository) Save(transaction Transaction) (Transaction, error) {
	err := r.db.Create(&transaction).Error
	if err != nil {
		return transaction, err
	}

	return transaction, nil
}

func (r repository) Update(transaction Transaction) (Transaction, error) {
	err := r.db.Save(&transaction).Error
	if err != nil {
		return transaction, err
	}

	return transaction, nil
}

func (r repository) GetById(Id int) (Transaction, error) {
	var transaction Transaction
	// get nested key relations
	err := r.db.Where("id = ?", Id).Find(&transaction).Error
	if err != nil {
		return transaction, err
	}

	return transaction, nil
}
