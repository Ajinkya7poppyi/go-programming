package dishes

import "errors"

// Dish ...
type Dish interface {
	LookIn()
	GetContents() string
}

const (
	BURGER = iota
	NOODLES
	PIZZA
)

// CreateDish ...
func CreateDish(myType int) (Dish, error) {
	switch myType {
	case BURGER:
		return new(Burger), nil
	case NOODLES:
		return new(Noodles), nil
	case PIZZA:
		return new(Pizza), nil
	default:
		return nil, errors.New("Invalid Dish Type")
	}
}
