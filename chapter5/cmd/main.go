package main

import (
	"context"
	"log"

	"coffeeco/internal/payment"
	"coffeeco/internal/purchase"
	"coffeeco/internal/store"
)

func main() {

	ctx := context.Background()

	someApiKey := "12345"
	mongoConString := "12345"
	csvc, err := payment.NewStripeService(someApiKey)
	if err != nil {
		log.Fatal(err)
	}

	prepo, err := purchase.NewMongoRepo(ctx, mongoConString)
	if err != nil {
		log.Fatal(err)
	}
	if err := prepo.Ping(ctx); err != nil {
		log.Fatal(err)
	}

	sRepo, err := store.NewMongoRepo(ctx, mongoConString)
	if err != nil {
		log.Fatal(err)
	}
	if err := sRepo.Ping(ctx); err != nil {
		log.Fatal(err)
	}

	sSvc := store.NewService(sRepo)

	_ = purchase.NewService(csvc, prepo, sSvc)

}
