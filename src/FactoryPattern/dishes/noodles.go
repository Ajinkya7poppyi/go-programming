package dishes

// Noodles ...
type Noodles struct {
	dishType string
}

// LookIn ...
func (nl *Noodles) LookIn() {
	nl.dishType = "Noodles"
}

// GetContents ...
func (nl *Noodles) GetContents() string {
	return "This is description about " + nl.dishType + " Its rice noodles tossed in sour sauces with onion and vegetables"
}
