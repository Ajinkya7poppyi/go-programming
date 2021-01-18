package dishes

// Pizza ...
type Pizza struct {
	dishType string
}

// LookIn ...
func (pz *Pizza) LookIn() {
	pz.dishType = "Pizza"
}

// GetContents ...
func (pz *Pizza) GetContents() string {
	return "This is description about " + pz.dishType + " Its dough base topped with vegies grilled in oven"
}
