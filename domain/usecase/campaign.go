package usecase

import (
	"errors"
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
	UpdateCampaign(idCampaign input.GetCampaignDetailInput, inputData input.CreateCampaign) (model.Campaign, error)
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
	campaigns, err := c.CampaignRepo.FindByUserId(userId)
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
		return newCampaign, err
	}
	return newCampaign, nil
}

func (c *campaignUseCase) UpdateCampaign(campaignID input.GetCampaignDetailInput, campaignData input.CreateCampaign) (model.Campaign, error) {
	campaign, err := c.CampaignRepo.FindCampaignById(campaignID.ID)
	if err != nil {
		return campaign, errors.New("unauthorized id nya")
	}

	if campaign.UserId != campaignData.User.Id {
		return campaign, errors.New("not an owner of the campaign")
	}

	campaign.Name = campaignData.Name
	campaign.Description = campaignData.Description
	campaign.ShortDescription = campaignData.ShortDescription
	campaign.Perks = campaignData.Perks
	campaign.GoalAmount = campaignData.GoalAmount

	updateCampaign, err := c.CampaignRepo.Update(campaign)
	if err != nil {
		return updateCampaign, errors.New("failed to update campaign")
	}
	return updateCampaign, nil
}
