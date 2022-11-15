package chapter2

import (
	"context"
)

type UserType = int
type SubscriptionType = int

const (
	unknownUserType UserType = iota
	lead
	customer
	churned
	lostLead
)

const (
	unknownSubscriptionType SubscriptionType = iota
	basic
	premium
	exclusive
)

type UserAddRequest struct {
	UserType       UserType
	Email          string
	SubType        SubscriptionType
	PaymentDetails PaymentDetails
}

type UserModifyRequest struct {
	ID             string
	UserType       UserType
	Email          string
	SubType        SubscriptionType
	PaymentDetails PaymentDetails
}

type User struct {
	ID             string
	PaymentDetails PaymentDetails
}

type PaymentDetails struct {
	stripeTokenID string
}

type UserManager interface {
	AddUser(ctx context.Context, request UserAddRequest) (User, error)
	ModifyUser(ctx context.Context, request UserModifyRequest) (User, error)
}
