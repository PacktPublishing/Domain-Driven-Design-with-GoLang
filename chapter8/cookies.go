package chapter8

import "context"

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

func (c *CookieService) PurchaseCookies(ctx context.Context, amountOfCookiesToPurchase int) error {
	//TODO: ask how much cookies cost. This is a placeholder.
	priceOfCookie := 5

	cookiesInStock := c.stockChecker.AmountInStock(ctx)
	if amountOfCookiesToPurchase > cookiesInStock {
		//TODO: what do I do in this situation?
	}
	cost := priceOfCookie * amountOfCookiesToPurchase

	//TODO: where do I get cardtoken from?
	if err := c.cardCharger.ChargeCard(ctx, "some-token", cost); err != nil {
		//TODO: handle this later.
	}

	if err := c.emailSender.SendEmailReceipt(ctx, "some-email"); err != nil {
		//TODO: handle error later
	}
	return nil
}
