package input

import "github.com/galang-dana/domain/model"

type GetCampaignDetailInput struct {
	ID string `uri:"id" binding:"required"`
}

type CreateCampaign struct {
	Name             string `json:"name" binding:"required"`
	ShortDescription string `json:"short_description" binding:"required"`
	Description      string `json:"description" binding:"required"`
	GoalAmount       int    `json:"goal_amount" binding:"required"`
	Perks            string `json:"perks" binding:"required"`
	User             model.User
}
