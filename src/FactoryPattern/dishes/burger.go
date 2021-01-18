package dishes

// Burger ...
type Burger struct {
	dishType string
}

// LookIn ...
func (br *Burger) LookIn() {
	br.dishType = "Burger"
}

// GetContents ...
func (br *Burger) GetContents() string {
	return "This is description about " + br.dishType + " Its patty sandwiched between bread buns"
}
