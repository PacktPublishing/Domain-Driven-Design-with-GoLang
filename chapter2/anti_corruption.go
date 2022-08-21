package chapter2

import (
	"errors"
	"time"
)

type Campaign struct {
	ID      string
	Title   string
	Goal    string
	EndDate time.Time
}

type MarketingCampaignModel struct {
	Id       string `json:"id"`
	Metadata struct {
		Name     string `json:"name"`
		Category string `json:"category"`
		EndDate  string `json:"endDate"`
	} `json:"metadata"`
}

func (m *MarketingCampaignModel) ToCampaign() (*Campaign, error) {
	if m.Id == "" {
		return nil, errors.New("campaign ID cannot be empty")
	}
	formattedDate, err := time.Parse("2006-01-02", m.Metadata.EndDate)
	if err != nil {
		return nil, errors.New("endDate was not in a parsable format")
	}

	return &Campaign{
		ID:      m.Id,
		Title:   m.Metadata.Name,
		Goal:    m.Metadata.Category,
		EndDate: formattedDate,
	}, nil
}
