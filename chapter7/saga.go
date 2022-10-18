package chapter7

import "context"

type Saga interface {
	Execute(ctx context.Context) error
	Rollback(ctx context.Context) error
}

type OrderCreator struct{}

func (o OrderCreator) Execute(ctx context.Context) error {
	return o.createOrder(ctx)
}

func (o OrderCreator) Rollback(ctx context.Context) error {
	//Rollback Saga here
	return nil
}
func (o OrderCreator) createOrder(ctx context.Context) error {
	// Create Order here
	return nil
}

type PaymentCreator struct{}

func (p PaymentCreator) Execute(ctx context.Context) error {
	return p.createPayment(ctx)
}

func (p PaymentCreator) Rollback(ctx context.Context) error {
	//Rollback Saga here
	return nil
}

func (p PaymentCreator) createPayment(ctx context.Context) error {
	// Create payment here
	return nil
}

type SagaManager struct {
	actions []Saga
}

func (s SagaManager) Handle(ctx context.Context) {
	for i, action := range s.actions {
		if err := action.Execute(ctx); err != nil {
			for j := 0; j < i; j++ {
				if err := s.actions[j].Rollback(ctx); err != nil {
					// One of our compensation actions failed; we need to handle it (perhaps by emitting a message to a
					// a messagebus.
				}
			}
		}
	}
}
