package main

//Making custom types

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

func newBill(name string) bill {
	b := bill{
		name:  name,
		items: map[string]float64{"pie": 34.55, "sandwich": 5.99},
		tip:   0,
	}
	return b
}

// bill type is recieved by the function and accessed using b
func (b bill) format() string {
	return b.name
}
