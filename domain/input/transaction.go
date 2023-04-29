package input

type GetCampaignTransaction struct {
	ID string `uri:"id" binding:"required"`
}
