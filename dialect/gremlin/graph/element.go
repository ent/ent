package graph

// Element defines a base struct for graph elements.
type Element struct {
	ID    interface{} `json:"id"`
	Label string      `json:"label"`
}

// NewElement create a new graph element.
func NewElement(id interface{}, label string) Element {
	return Element{id, label}
}
