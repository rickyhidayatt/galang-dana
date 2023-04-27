package repository

import (
	"fmt"

	"github.com/galang-dana/domain/model"
	"gorm.io/gorm"
)

type CampaignRepository interface {
	FindAll() ([]model.Campaign, error)
	FindByUserId(userId string) ([]model.Campaign, error)
	FindCampaignById(userId string) (model.Campaign, error)
	Save(campaign model.Campaign) (model.Campaign, error)
	Update(campaign model.Campaign) (model.Campaign, error)
	CreateImage(campaignImage model.Image) (model.Image, error)
	MarkAllImagesAsNonPrimary(campaignID string) (bool, error)
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

//	func (c *campaignRepository) FindByUserId(userId string) ([]model.Campaign, error) {
//		var campaigns []model.Campaign
//		err := c.db.Where("user_id = ?", userId).Preload("Images", "images.is_primary = true").Find(&campaigns).Error
//		if err != nil {
//			return campaigns, err
//		}
//		return campaigns, nil
//	}

func (c *campaignRepository) FindByUserId(userID string) ([]model.Campaign, error) {
	var campaigns []model.Campaign
	err := c.db.Where("user_id = ?", userID).Find(&campaigns).Error
	if err != nil {
		return campaigns, err
	}

	// Ambil data images terpisah
	for i := range campaigns {
		var images []model.Image
		err := c.db.Where("campaign_id = ?", campaigns[i].Id).Find(&images).Error
		if err != nil {
			return campaigns, err
		}
		campaigns[i].Images = images
	}

	return campaigns, nil
}

// func (c *campaignRepository) FindCampaignById(Id string) (model.Campaign, error) {
// 	var campaign model.Campaign

// 	err := c.db.Preload("user").Preload("Images").Where("id = ?", Id).Find(&campaign).Error
// 	if err != nil {
// 		log.Fatal(err)
// 		return campaign, err
// 	}

// 	return campaign, nil
// }

func (c *campaignRepository) FindCampaignById(ID string) (model.Campaign, error) {
	var campaign model.Campaign

	err := c.db.Where("id = ?", ID).First(&campaign).Error
	if err != nil {
		return campaign, fmt.Errorf("failed to find campaign: %w", err)
	}

	user := model.User{}
	err = c.db.Where("id = ?", campaign.UserId).First(&user).Error
	if err != nil {
		return campaign, fmt.Errorf("failed to find user: %w", err)
	}
	campaign.User = user

	images := []model.Image{}
	err = c.db.Where("campaign_id = ?", campaign.Id).Find(&images).Error
	if err != nil {
		return campaign, fmt.Errorf("failed to find images: %w", err)
	}
	campaign.Images = images

	return campaign, nil
}

func (c *campaignRepository) Save(campaign model.Campaign) (model.Campaign, error) {
	err := c.db.Create(&campaign).Error
	if err != nil {
		return campaign, err
	}
	return campaign, nil
}

func (c *campaignRepository) Update(campaign model.Campaign) (model.Campaign, error) {
	err := c.db.Save(&campaign).Error
	if err != nil {
		return campaign, err
	}
	return campaign, nil
}

func (c *campaignRepository) CreateImage(campaignImage model.Image) (model.Image, error) {
	err := c.db.Create(&campaignImage).Error
	if err != nil {
		return campaignImage, err
	}
	return campaignImage, nil
}

func (c *campaignRepository) MarkAllImagesAsNonPrimary(campaignID string) (bool, error) {
	err := c.db.Model(&model.Image{}).Where("campaign_id = ?", campaignID).Update("is_primary", false).Error
	if err != nil {
		return false, err
	}
	return true, nil
}
