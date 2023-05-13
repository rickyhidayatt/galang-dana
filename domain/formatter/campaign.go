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

func FormatCampaign(campaign model.Campaign) CampaignFormater {
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
		camp := FormatCampaign(c)
		campaign = append(campaign, camp)
	}

	return campaign
}

type CampaignDetailFormater struct {
	Id               string                  `json:"id"`
	UserId           string                  `json:"user_id"`
	Name             string                  `json:"name"`
	ShortDescription string                  `json:"short_description"`
	Description      string                  `json:"description"`
	ImageURL         string                  `json:"image_url"`
	GoalAmount       int                     `json:"goal_amount"`
	BackerCount      int                     `json:"backer_count"`
	CurrentAmount    int                     `json:"current_amount"`
	Slug             string                  `json:"slug"`
	Perks            []string                `json:"perks"`
	User             campaignUserFormater    `json:"user"`
	Images           []campaignImageFormater `json:"images"`
}

type campaignUserFormater struct {
	Name     string `json:"name"`
	ImageURL string `json:"image_url"`
}

type campaignImageFormater struct {
	ImageUrl  string `json:"image_url"`
	IsPrimary bool   `json:"is_primary"`
}

func FormatCampaignDetail(campaign model.Campaign) CampaignDetailFormater {
	var perks []string

	user := campaignUserFormater{
		Name:     campaign.Name,
		ImageURL: campaign.User.AvatarFileName,
	}

	image := []campaignImageFormater{}

	for _, img := range campaign.Images {
		images := campaignImageFormater{}
		images.ImageUrl = img.FileName
		images.IsPrimary = img.IsPrimary

		image = append(image, images)
	}

	campaignFormat := CampaignDetailFormater{
		Id:               campaign.Id,
		UserId:           campaign.UserId,
		Name:             campaign.Name,
		ShortDescription: campaign.ShortDescription,
		Description:      campaign.Description,
		BackerCount:      campaign.BackerCount,
		CurrentAmount:    campaign.CurrentAmount,
		Slug:             campaign.Slug,
		User:             user,
		Images:           image,
	}

	if len(campaign.Images) > 0 {
		campaignFormat.ImageURL = campaign.Images[0].FileName
	}

	for _, perk := range strings.Split(campaign.Perks, ",") {
		perks = append(perks, strings.TrimSpace(perk))
	}

	campaignFormat.Perks = perks

	return campaignFormat
}
