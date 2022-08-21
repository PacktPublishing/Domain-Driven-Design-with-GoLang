package chapter3

import (
	"time"

	"github.com/Rhymond/go-money"
)

// Auction is an entity to represent our auction construct.
type Auction struct {
	ID int
	// We use a specific money library as floats are not good ways to represent money.
	startingPrice money.Money
	sellerID      int
	createdAt     time.Time
	auctionStart  time.Time
	auctionEnd    time.Time
}
