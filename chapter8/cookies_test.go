package chapter8_test

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"

	"github.com/PacktPublishing/Domain-Driven-Design-with-GoLang/chapter8"
	"github.com/PacktPublishing/Domain-Driven-Design-with-GoLang/chapter8/mocks"
)

func Test_CookiePurchases(t *testing.T) {
	t.Run(`Given a user tries to purchase a cookie and we have them in stock, 
		"when they tap their card, they get charged and then receive an email receipt a few moments later.`,
		func(t *testing.T) {
			var (
				ctrl = gomock.NewController(t)
				e    = mocks.NewMockEmailSender(ctrl)
				c    = mocks.NewMockCardCharger(ctrl)
				s    = mocks.NewMockCookieStockChecker(ctrl)

				ctx       = context.Background()
				email     = "some@email.com"
				cardToken = "token"
			)
			cookiesToBuy := 5
			totalExpectedCost := 250

			cs, err := chapter8.NewCookieService(e, c, s)
			if err != nil {
				t.Fatalf("expected no error but got %v", err)
			}

			gomock.InOrder(
				s.EXPECT().AmountInStock(ctx).Times(1).Return(cookiesToBuy),
				c.EXPECT().ChargeCard(ctx, cardToken, totalExpectedCost).Times(1).Return(nil),
				e.EXPECT().SendEmailReceipt(ctx, email).Times(1).Return(nil),
			)

			err = cs.PurchaseCookies(ctx, cookiesToBuy, cardToken, email)
			if err != nil {
				t.Fatalf("expected no error but got %v", err)
			}
		})

	t.Run(`Given a user tries to purchase a cookie and we don’t have any in stock, we return an error to the cashier 
			so they can apologize to the customer.`, func(t *testing.T) {
		var (
			ctrl = gomock.NewController(t)
			e    = mocks.NewMockEmailSender(ctrl)
			c    = mocks.NewMockCardCharger(ctrl)
			s    = mocks.NewMockCookieStockChecker(ctrl)

			ctx       = context.Background()
			email     = "some@email.com"
			cardToken = "token"
		)
		cookiesToBuy := 5

		cs, err := chapter8.NewCookieService(e, c, s)
		if err != nil {
			t.Fatalf("expected no error but got %v", err)
		}

		gomock.InOrder(
			s.EXPECT().AmountInStock(ctx).Times(1).Return(0),
		)

		err = cs.PurchaseCookies(ctx, cookiesToBuy, cardToken, email)
		if err == nil {
			t.Fatal("expected no error but got none")
		}
	})

	t.Run(`Given a user tries to purchase a cookie, we have them in stock, but their card gets declined, we return 
		an error to the cashier so that we can ban the customer from the store.`, func(t *testing.T) {
		var (
			ctrl = gomock.NewController(t)
			e    = mocks.NewMockEmailSender(ctrl)
			c    = mocks.NewMockCardCharger(ctrl)
			s    = mocks.NewMockCookieStockChecker(ctrl)

			ctx       = context.Background()
			email     = "some@email.com"
			cardToken = "token"
		)
		cookiesToBuy := 5
		totalExpectedCost := 250

		cs, err := chapter8.NewCookieService(e, c, s)
		if err != nil {
			t.Fatalf("expected no error but got %v", err)
		}

		gomock.InOrder(
			s.EXPECT().AmountInStock(ctx).Times(1).Return(cookiesToBuy),
			c.EXPECT().ChargeCard(ctx, cardToken, totalExpectedCost).Times(1).Return(errors.New("some error")),
		)

		err = cs.PurchaseCookies(ctx, cookiesToBuy, cardToken, email)
		if err == nil {
			t.Fatal("expected an error but got none")
		}
		if err.Error() != "your card was declined, you are banned!" {
			t.Fatalf("error was unexpected, got %v", err.Error())
		}
	})
	t.Run(`Given a user purchases a cookie and we have them in stock, their card is charged successfully but we 
		fail to send an email, we return a message to the cashier so they know can notify the customer that they will not
		get an e-mail, but the transaction is still considered done.`, func(t *testing.T) {
		var (
			ctrl = gomock.NewController(t)
			e    = mocks.NewMockEmailSender(ctrl)
			c    = mocks.NewMockCardCharger(ctrl)
			s    = mocks.NewMockCookieStockChecker(ctrl)

			ctx       = context.Background()
			email     = "some@email.com"
			cardToken = "token"
		)
		cookiesToBuy := 5
		totalExpectedCost := 250

		cs, err := chapter8.NewCookieService(e, c, s)
		if err != nil {
			t.Fatalf("expected no error but got %v", err)
		}

		gomock.InOrder(
			s.EXPECT().AmountInStock(ctx).Times(1).Return(cookiesToBuy),
			c.EXPECT().ChargeCard(ctx, cardToken, totalExpectedCost).Times(1).Return(nil),
			e.EXPECT().SendEmailReceipt(ctx, email).Times(1).Return(errors.New("failed to send email")),
		)

		err = cs.PurchaseCookies(ctx, cookiesToBuy, cardToken, email)
		if err == nil {
			t.Fatal("expected an error but got none")
		}
		if err.Error() != "we are sorry but the email receipt did not send" {
			t.Fatalf("error was unexpected, got %v", err.Error())
		}
	})

	t.Run(`Given someone wants to purchase more cookies than we have in stock we only charge them for the ones we do have`,
		func(t *testing.T) {
			var (
				ctrl = gomock.NewController(t)
				e    = mocks.NewMockEmailSender(ctrl)
				c    = mocks.NewMockCardCharger(ctrl)
				s    = mocks.NewMockCookieStockChecker(ctrl)

				ctx       = context.Background()
				email     = "some@email.com"
				cardToken = "token"
			)
			requestedCookiesToBuy := 5
			inStock := 3
			totalExpectedCost := 150

			cs, err := chapter8.NewCookieService(e, c, s)
			if err != nil {
				t.Fatalf("expected no error but got %v", err)
			}

			gomock.InOrder(
				s.EXPECT().AmountInStock(ctx).Times(1).Return(inStock),
				c.EXPECT().ChargeCard(ctx, cardToken, totalExpectedCost).Times(1).Return(nil),
				e.EXPECT().SendEmailReceipt(ctx, email).Times(1).Return(nil),
			)

			err = cs.PurchaseCookies(ctx, requestedCookiesToBuy, cardToken, email)
			if err != nil {
				t.Fatalf("expected no error but got %v", err)
			}
		})
}
