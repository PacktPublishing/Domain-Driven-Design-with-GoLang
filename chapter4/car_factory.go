package chapter4

import "errors"

type Car interface {
	BeepBeep()
}

type BMW struct {
	heatedSeatSubscriptionEnabled bool
}

func (B BMW) BeepBeep() {
	//TODO implement me
	panic("implement me")
}

type Tesla struct {
	autoPilotEnabled bool
}

func (t Tesla) BeepBeep() {
	//TODO implement me
	panic("implement me")
}

func BuildCar(carType string) (Car, error) {
	switch carType {
	case "bmw":
		return BMW{heatedSeatSubscriptionEnabled: true}, nil

	case "tesla":
		return Tesla{autoPilotEnabled: true}, nil
	default:
		return nil, errors.New("unknown car type")
	}
}

func main() {
	myCar, _ := BuildCar("tesla")
}
