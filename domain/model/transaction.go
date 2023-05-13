package model

import (
	"time"
)

type Transaction struct {
	ID         string
	CampaignID string
	UserID     string
	Amount     int
	Status     string
	Code       string
	PaymentURL string
	User       User
	Campaign   Campaign
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
