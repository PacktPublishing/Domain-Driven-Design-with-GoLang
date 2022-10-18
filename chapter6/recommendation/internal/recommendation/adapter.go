package recommendation

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/Rhymond/go-money"
)

type partnerShipsResponse struct {
	AvailableHotels []struct {
		Name               string `json:"name"`
		PriceInUSDPerNight int    `json:"priceInUSDPerNight"`
	} `json:"availableHotels"`
}

type PartnershipAdaptor struct {
	client *http.Client
	url    string
}

func NewPartnerShipAdaptor(client *http.Client, url string) (*PartnershipAdaptor, error) {
	if client == nil {
		return nil, errors.New("client cannot be nil")
	}
	if url == "" {
		return nil, errors.New("url cannot be empty")
	}

	return &PartnershipAdaptor{client: client, url: url}, nil
}

func (p PartnershipAdaptor) GetAvailability(ctx context.Context, tripStart time.Time, tripEnd time.Time, location string) ([]Option, error) {
	from := fmt.Sprintf("%d-%d-%d", tripStart.Year(), tripStart.Month(), tripStart.Day())
	to := fmt.Sprintf("%d-%d-%d", tripEnd.Year(), tripEnd.Month(), tripEnd.Day())

	url := fmt.Sprintf("%s/partnerships?location=%s&from=%s&to=%s", p.url, location, from, to)
	res, err := p.client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to call partnerships: %w", err)
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("bad request to partnerships: %d", res.StatusCode)
	}

	var pr partnerShipsResponse
	if err := json.NewDecoder(res.Body).Decode(&pr); err != nil {
		return nil, fmt.Errorf("could not decode the response body of partnerships: %w", err)
	}

	opts := make([]Option, len(pr.AvailableHotels))
	for i, p := range pr.AvailableHotels {
		opts[i] = Option{
			HotelName:     p.Name,
			Location:      location,
			PricePerNight: *money.New(int64(p.PriceInUSDPerNight), "USD"),
		}
	}
	return opts, nil
}
