package input

type GetCampaignDetailInput struct {
	ID string `uri:"id" binding:"required"`
}
