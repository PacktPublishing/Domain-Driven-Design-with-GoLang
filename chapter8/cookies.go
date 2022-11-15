package chapter8

import (
	"context"
	"errors"
)

type (
	EmailSender interface {
		SendEmailReceipt(ctx context.Context, emailAddress string) error
	}
	CardCharger interface {
		ChargeCard(ctx context.Context, cardToken string, amountInCents int) error
	}
	CookieStockChecker interface {
		AmountInStock(ctx context.Context) int
	}

	CookieService struct {
		emailSender  EmailSender
		cardCharger  CardCharger
		stockChecker CookieStockChecker
	}
)

func NewCookieService(e EmailSender, c CardCharger, a CookieStockChecker) (*CookieService, error) {
	return &CookieService{
		emailSender:  e,
		cardCharger:  c,
		stockChecker: a,
	}, nil
}

func (c *CookieService) PurchaseCookies(
	ctx context.Context,
	amountOfCookiesToPurchase int,
	cardToken string,
	email string,
) error {
	priceOfCookie := 50

	cookiesInStock := c.stockChecker.AmountInStock(ctx)
	if cookiesInStock == 0 {
		return errors.New("no cookies in stock sorry :(")
	}
	if amountOfCookiesToPurchase > cookiesInStock {
		amountOfCookiesToPurchase = cookiesInStock
	}
	cost := priceOfCookie * amountOfCookiesToPurchase

	if err := c.cardCharger.ChargeCard(ctx, cardToken, cost); err != nil {
		return errors.New("your card was declined, you are banned!")
	}

	if err := c.emailSender.SendEmailReceipt(ctx, email); err != nil {
		return errors.New("we are sorry but the email receipt did not send")
	}
	return nil
}
