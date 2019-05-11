package appliances

import "errors"

// Abstract type for factory to create
type Appliance interface {
	Start()
	GetPurpose() string
}

const (
	STOVE = iota
	FRIDGE
	MICROWAVE
)

// this is the factory
func CreateAppliance(myType int) (Appliance, error) {
	switch myType {
	case STOVE:
		return new(Stove), nil
	case FRIDGE:
		return new(Fridge), nil
	case MICROWAVE:
		return new(Microwave), nil
	default:
		return nil, errors.New("invalid appliance type")
	}
}