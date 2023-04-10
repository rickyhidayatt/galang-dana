package usecase

import (
	"github.com/galang-dana/domain/model"
	"github.com/galang-dana/domain/repository"
)

type CampaignUseCase interface {
	FindCampaigns(userId string) ([]model.Campaign, error)
}

type campaignUseCase struct {
	CampaignRepo repository.CampaignRepository
}

func NewCampaignUseCase(r repository.CampaignRepository) CampaignUseCase {
	return &campaignUseCase{r}
}

func (c *campaignUseCase) FindCampaigns(userId string) ([]model.Campaign, error) {
	if userId == "" {
		campaigns, err := c.CampaignRepo.FindAll()
		if err != nil {
			return campaigns, err
		}
		return campaigns, nil
	}
	campaigns, err := c.CampaignRepo.FindById(userId)
	if err != nil {
		return campaigns, err
	}
	return campaigns, nil
}
