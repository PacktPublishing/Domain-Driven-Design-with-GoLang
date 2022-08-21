package chapter3

import (
	"errors"
	"time"

	"github.com/Rhymond/go-money"
)

type AuctionRefactored struct {
	id            int
	startingPrice money.Money
	sellerID      int
	createdAt     time.Time
	auctionStart  time.Time
	auctionEnd    time.Time
}

func (a *AuctionRefactored) GetAuctionElapsedDuration() time.Duration {
	return a.auctionStart.Sub(a.auctionEnd)
}

func (a *AuctionRefactored) GetAuctionEndTimeInUTC() time.Time {
	return a.auctionEnd
}

func (a *AuctionRefactored) SetAuctionEnd(auctionEnd time.Time) error {
	if err := a.validateTimeZone(auctionEnd); err != nil {
		return err
	}
	a.auctionEnd = auctionEnd
	return nil
}

func (a *AuctionRefactored) GetAuctionStartTimeInUTC() time.Time {
	return a.auctionStart
}

func (a *AuctionRefactored) SetAuctionStartTimeInUTC(auctionStart time.Time) error {
	if err := a.validateTimeZone(auctionStart); err != nil {
		return err
	}

	// in reality, we would likely persist this to a database
	a.auctionStart = auctionStart
	return nil
}

func (a *AuctionRefactored) GetId() int {
	return a.id
}

func (a *AuctionRefactored) validateTimeZone(t time.Time) error {
	tz, _ := t.Zone()
	if tz != time.UTC.String() {
		return errors.New("time zone must be UTC")
	}
	return nil
}
