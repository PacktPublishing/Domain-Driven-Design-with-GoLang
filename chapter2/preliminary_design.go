package chapter2

import (
	"context"
)

type UserType = int
type subscriptionType = int

const (
	unknownUserType UserType = iota
	lead
	customer
	churned
	lostLead
)

const (
	unknownSubscriptionType subscriptionType = iota
	basic
	premium
	exclusive
)

type UserAddRequest struct {
	userType       UserType
	email          string
	subType        subscriptionType
	paymentDetails PaymentDetails
}

type UserModifyRequest struct {
	id             string
	userType       UserType
	email          string
	subType        subscriptionType
	paymentDetails PaymentDetails
}

type User struct {
	id             string
	paymentDetails PaymentDetails
}

type PaymentDetails struct {
	stripeTokenID string
}

type UserManager interface {
	AddUser(ctx context.Context, request UserAddRequest) (User, error)
	ModifyUser(ctx context.Context, request UserModifyRequest) (User, error)
}
