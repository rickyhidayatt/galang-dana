package repository

import (
	"github.com/galang-dana/domain/model"
	"gorm.io/gorm"
)

type TransactionRepository interface {
	GetByCampaignID(campaignID string) ([]model.Transaction, error)
	GetByUserID(userID string) ([]model.Transaction, error)
	GetByID(TransactionID string) (model.Transaction, error)
	Save(transaction model.Transaction) (model.Transaction, error)
	Update(transaction model.Transaction) (model.Transaction, error)
}

type transactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(g *gorm.DB) TransactionRepository {
	return &transactionRepository{g}
}

func (r *transactionRepository) GetByCampaignID(campaignID string) ([]model.Transaction, error) {
	var transaction []model.Transaction

	err := r.db.Where("campaign_id = ?", campaignID).Order("id desc").Find(&transaction).Error
	if err != nil {
		return transaction, err
	}

	for i := range transaction {
		var user model.User
		err := r.db.Where("id = ?", transaction[i].UserID).Find(&user).Error
		if err != nil {
			return transaction, err
		}
		transaction[i].User = user
	}
	return transaction, nil
}

func (r *transactionRepository) GetByUserID(userID string) ([]model.Transaction, error) {
	var transaction []model.Transaction

	err := r.db.Preload("Campaign.Images", "images.is_primary = true").Where("user_id = ?", userID).Order("created_at desc").Find(&transaction).Error

	if err != nil {
		return transaction, err
	}

	return transaction, nil
}

func (r *transactionRepository) Save(transaction model.Transaction) (model.Transaction, error) {
	err := r.db.Create(&transaction).Error
	if err != nil {
		return transaction, err
	}
	return transaction, nil
}

func (r *transactionRepository) Update(transaction model.Transaction) (model.Transaction, error) {
	err := r.db.Save(&transaction).Error
	if err != nil {
		return transaction, err
	}
	return transaction, nil
}

func (r *transactionRepository) GetByID(TransactionID string) (model.Transaction, error) {
	var transaction model.Transaction
	err := r.db.Where("id = ?", TransactionID).Find(&transaction).Error
	if err != nil {
		return transaction, err
	}
	return transaction, nil
}
