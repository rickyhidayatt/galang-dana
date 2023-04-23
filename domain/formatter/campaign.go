package formatter

import (
	"strings"

	"github.com/galang-dana/domain/model"
)

type CampaignFormater struct {
	Id               string `json:"id"`
	UserId           string `json:"user_id"`
	Name             string `json:"name"`
	ShortDescription string `json:"short_description"`
	ImageURL         string `json:"image_url"`
	GoalAmount       int    `json:"goal_amount"`
	CurrentAmount    int    `json:"current_amount"`
	Slug             string `json:"slug"`
}

func formatCampaign(campaign model.Campaign) CampaignFormater {
	campaignFormatter := CampaignFormater{
		Id:               campaign.Id,
		UserId:           campaign.UserId,
		Name:             campaign.Name,
		ShortDescription: campaign.ShortDescription,
		GoalAmount:       campaign.GoalAmount,
		CurrentAmount:    campaign.CurrentAmount,
		Slug:             campaign.Slug,
		ImageURL:         "",
	}

	if len(campaign.Images) > 0 {
		campaignFormatter.ImageURL = campaign.Images[0].FileName
	}

	return campaignFormatter
}

func FormatCampaigns(campaigns []model.Campaign) []CampaignFormater {

	campaign := []CampaignFormater{}
	for _, c := range campaigns {
		camp := formatCampaign(c)
		campaign = append(campaign, camp)
	}

	return campaign
}

type CampaignDetailFormater struct {
	Id               string   `json:"id"`
	UserId           string   `json:"user_id"`
	Name             string   `json:"name"`
	ShortDescription string   `json:"short_description"`
	Description      string   `json:"description"`
	ImageURL         string   `json:"image_url"`
	GoalAmount       int      `json:"goal_amount"`
	BackerCount      int      `json:"backer_count"`
	CurrentAmount    int      `json:"current_amount"`
	Slug             string   `json:"slug"`
	Perks            []string `json:"perks"`
}

func FormatCampaignDetail(campaign model.Campaign) CampaignDetailFormater {

	campaignFormat := CampaignDetailFormater{
		Id:               campaign.Id,
		UserId:           campaign.UserId,
		Name:             campaign.Name,
		ShortDescription: campaign.ShortDescription,
		Description:      campaign.Description,
		// ImageURL:         "",
		BackerCount:   campaign.BackerCount,
		CurrentAmount: campaign.CurrentAmount,
		Slug:          campaign.Slug,
	}

	if len(campaign.Images) > 0 {
		campaignFormat.ImageURL = campaign.Images[0].FileName
	}

	var perks []string

	for _, perk := range strings.Split(campaign.Perks, ",") {
		perks = append(perks, strings.TrimSpace(perk))
	}

	campaignFormat.Perks = perks

	return campaignFormat
}
