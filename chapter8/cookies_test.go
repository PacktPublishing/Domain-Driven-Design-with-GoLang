package chapter8_test

import "testing"

func Test_CookiePurchases(t *testing.T) {
	t.Run(`Given a user tries to purchase a cookie and we have them in stock, 
		"when they tap their card, they get charged and then receive an email receipt a few moments later.`,
		func(t *testing.T) {
			t.FailNow()
		})
}
