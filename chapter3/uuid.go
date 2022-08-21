package chapter3

import "github.com/google/uuid"

type SomeEntity struct {
	id uuid.UUID
}

func NewSomeEntity() *SomeEntity {
	id := uuid.New()
	return &SomeEntity{id: id}
}
