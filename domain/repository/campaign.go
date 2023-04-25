package repository

import (
	"log"

	"github.com/galang-dana/domain/model"
	"gorm.io/gorm"
)

type CampaignRepository interface {
	FindAll() ([]model.Campaign, error)
	FindById(userId string) ([]model.Campaign, error)
	FindCampaignById(userId string) (model.Campaign, error)
	Save(campaign model.Campaign) (model.Campaign, error)
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

func (c *campaignRepository) FindCampaignById(userId string) (model.Campaign, error) {
	var campaign model.Campaign

	err := c.db.Where("user_id = ?", userId).Preload("User").Preload("Images").Find(&campaign).Error
	if err != nil {
		log.Fatal(err)
		return campaign, err
	}

	return campaign, nil
}

func (c *campaignRepository) Save(campaign model.Campaign) (model.Campaign, error) {
	err := c.db.Create(&campaign).Error
	if err != nil {
		return campaign, err
	}
	return campaign, nil
}
