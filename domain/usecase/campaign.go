package usecase

import (
	"github.com/galang-dana/domain/input"
	"github.com/galang-dana/domain/model"
	"github.com/galang-dana/domain/repository"
)

type CampaignUseCase interface {
	GetCampaigns(userId string) ([]model.Campaign, error)
	GetCampaignById(input input.GetCampaignDetailInput) (model.Campaign, error)
}

type campaignUseCase struct {
	CampaignRepo repository.CampaignRepository
}

func NewCampaignUseCase(r repository.CampaignRepository) CampaignUseCase {
	return &campaignUseCase{r}
}

func (c *campaignUseCase) GetCampaigns(userId string) ([]model.Campaign, error) {
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

func (c *campaignUseCase) GetCampaignById(input input.GetCampaignDetailInput) (model.Campaign, error) {
	campaign, err := c.CampaignRepo.FindCampaignById(input.ID)
	if err != nil {
		return campaign, err
	}
	return campaign, nil
}
