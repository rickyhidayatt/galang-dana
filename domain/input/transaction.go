package input

import "github.com/galang-dana/domain/model"

type GetCampaignTransaction struct {
	ID   string `uri:"id" binding:"required"`
	User model.User
}
