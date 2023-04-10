package repository

import (
	"github.com/galang-dana/domain/model"
	"gorm.io/gorm"
)

type CampaignRepository interface {
	FindAll() ([]model.Campaign, error)
	FindById(userId string) ([]model.Campaign, error)
}

type campaignRepository struct {
	db *gorm.DB
}

func NewCampaignRepository(g *gorm.DB) CampaignRepository {
	return &campaignRepository{g}
}

func (c *campaignRepository) FindAll() ([]model.Campaign, error) {

	var campaigns []model.Campaign
	err := c.db.Preload("Images", "images.is_primary = true").Find(&campaigns).Error
	if err != nil {
		return campaigns, err
	}
	return campaigns, nil
}

func (c *campaignRepository) FindById(userId string) ([]model.Campaign, error) {
	var campaigns []model.Campaign
	err := c.db.Where("user_id = ?", userId).Preload("Images", "images.is_primary = true").Find(&campaigns).Error
	if err != nil {
		return campaigns, err
	}
	return campaigns, nil
}
