package chapter4

import "errors"

type CheckoutService struct {
	shoppingCart *ShoppingCart
}

func NewCheckoutService(shoppingCart *ShoppingCart) *CheckoutService {
	return &CheckoutService{shoppingCart: shoppingCart}
}

func (c CheckoutService) AddProductToBasket(p *Product) error {
	if c.shoppingCart.IsFull {
		return errors.New("cannot add to cart, its full")
	}
	if p.CanBeBought() {
		c.shoppingCart.Products = append(c.shoppingCart.Products, *p)
		return nil
	}
	if c.shoppingCart.maxCartSize == len(c.shoppingCart.Products) {
		c.shoppingCart.IsFull = true
	}
	return nil
}
