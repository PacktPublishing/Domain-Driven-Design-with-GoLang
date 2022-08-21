package chapter2_test

import (
	"testing"

	"ddd-golang/chapter2"
)

func Test_ToCampaign(t *testing.T) {
	// example from the book
	m := chapter2.MarketingCampaignModel{
		Id: "4cdd4ba9-7c04-4a3d-ac52-71f37ba75d7f",
		Metadata: struct {
			Name     string `json:"name"`
			Category string `json:"category"`
			EndDate  string `json:"endDate"`
		}{
			Name:     "some campaign",
			Category: "growth",
			EndDate:  "2023-04-12",
		},
	}

	_, err := m.ToCampaign()
	if err != nil {
		t.Fatalf("err was not nil: %v", err)
	}
}
