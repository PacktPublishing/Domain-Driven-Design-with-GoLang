package gen

import _ "github.com/golang/mock/mockgen/model"

//go:generate mockgen -package mocks -destination chapter8/mocks/cookies.go github.com/PacktPublishing/Domain-Driven-Design-with-GoLang/chapter8 CookieStockChecker,CardCharger,EmailSender
