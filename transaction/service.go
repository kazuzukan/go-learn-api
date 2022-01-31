package transaction

import (
	"bwa-project/campaign"
	"bwa-project/payment"
	"strconv"

	"errors"
)

type Service interface {
	GetTransactionByCampaignId(input GetCampaignTransactionsInput) ([]Transaction, error)
	GetTransactionByUserId(userId int) ([]Transaction, error)
	CreateTransaction(input CreateTransactionInput) (Transaction, error)
	ProcessPayment(input TransactionNotificationInput) error
}

type service struct {
	repository         Repository
	campaignRepository campaign.Repository
	paymentService     payment.Service
}

func NewServices(repository Repository, campaignRepository campaign.Repository, paymentService payment.Service) service {
	return service{repository, campaignRepository, paymentService}
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

func (s service) CreateTransaction(input CreateTransactionInput) (Transaction, error) {
	// mapping input to transaction struct
	transaction := Transaction{}
	transaction.CampaignId = input.CampaignId
	transaction.Amount = input.Amount
	transaction.UserId = input.User.ID
	transaction.Status = "Pending"

	newTransaction, err := s.repository.Save(transaction)
	if err != nil {
		return newTransaction, err
	}

	paymentTransaction := payment.Transaction{
		Id:     newTransaction.Id,
		Amount: newTransaction.Amount,
	}

	paymentUrl, err := s.paymentService.GetPaymentUrl(paymentTransaction, input.User)
	if err != nil {
		return newTransaction, err
	}

	newTransaction.PaymentUrl = paymentUrl
	newTransaction, err = s.repository.Update(newTransaction)
	if err != nil {
		return newTransaction, err
	}

	return newTransaction, nil
}

func (s service) ProcessPayment(input TransactionNotificationInput) error {
	transaction_id, _ := strconv.Atoi(input.OrderId)

	transaction, err := s.repository.GetById(transaction_id)
	if err != nil {
		return err
	}

	if input.PaymentType == "credit_card" && input.TransactionStatus == "captured" && input.FraudStatus == "accept" {
		transaction.Status = "paid"
	} else if input.TransactionStatus == "settlement" {
		transaction.Status = "paid"
	} else if input.TransactionStatus == "deny" || input.TransactionStatus == "expire" || input.TransactionStatus == "canceled" {
		transaction.Status = "cancelled"
	}

	updatedTransaction, err := s.repository.Update(transaction)
	if err != nil {
		return err
	}

	campaign, err := s.campaignRepository.FindById(updatedTransaction.CampaignId)
	if err != nil {
		return err
	}

	if updatedTransaction.Status == "paid" {
		campaign.CurrentAmount += updatedTransaction.Amount
		campaign.BackerCount += 1

		_, err := s.campaignRepository.Update(campaign)
		if err != nil {
			return err
		}
	}

	return nil
}
