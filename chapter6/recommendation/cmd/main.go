package main

import (
	"log"
	"net/http"

	"github.com/hashicorp/go-retryablehttp"

	"github.com/PacktPublishing/Domain-Driven-Design-with-GoLang/chapter6/recommendation/internal/recommendation"
	"github.com/PacktPublishing/Domain-Driven-Design-with-GoLang/chapter6/recommendation/internal/transport"
)

func main() {
	c := retryablehttp.NewClient()
	c.RetryMax = 10

	partnerAdaptor, err := recommendation.NewPartnerShipAdaptor(
		c.StandardClient(),
		"http://localhost:3031",
	)
	if err != nil {
		log.Fatal("failed to create a partnerAdaptor: ", err)
	}

	svc, err := recommendation.NewService(partnerAdaptor)
	if err != nil {
		log.Fatal("failed to create a service: ", err)
	}

	handler, err := recommendation.NewHandler(*svc)
	if err != nil {
		log.Fatal("failed to create a handler: ", err)
	}

	m := transport.NewMux(*handler)

	if err := http.ListenAndServe(":4040", m); err != nil {
		log.Fatal("server errored: ", err)
	}
}
