package usecase

import (
	"fmt"

	"github.com/galang-dana/domain/input"
	"github.com/galang-dana/domain/model"
	"github.com/galang-dana/domain/repository"
	"github.com/galang-dana/utils"
	"github.com/gosimple/slug"
)

type CampaignUseCase interface {
	GetCampaigns(userId string) ([]model.Campaign, error)
	GetCampaignById(input input.GetCampaignDetailInput) (model.Campaign, error)
	CreateCampaign(inpt input.CreateCampaign) (model.Campaign, error)
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

func (c *campaignUseCase) CreateCampaign(inpt input.CreateCampaign) (model.Campaign, error) {

	src := fmt.Sprintf("%s %s", inpt.Name, inpt.User.Id)
	slugData := slug.Make(src)

	campaign := model.Campaign{
		Id:               utils.GenerateId(),
		Name:             inpt.Name,
		Description:      inpt.Description,
		ShortDescription: inpt.ShortDescription,
		Perks:            inpt.Perks,
		GoalAmount:       inpt.GoalAmount,
		UserId:           inpt.User.Id,
		Slug:             slugData,
	}

	newCampaign, err := c.CampaignRepo.Save(campaign)
	if err != nil {
		fmt.Println("ADA EROR DI SERVICE")
		return newCampaign, err
	}

	return newCampaign, nil

}
