package model

import "time"

type Campaign struct {
	Id               string
	UserId           string
	Name             string
	ShortDescription string
	Description      string
	Pearks           string
	BackerCount      int
	GoalAmount       int
	CurrentAmount    int
	Slug             string
	CreatedAt        time.Time
	UpdatedAt        time.Time
	Images           []Image
	User             User
}

type Image struct {
	Id         string
	CampaignId string
	FileName   string
	IsPrimary  bool
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
