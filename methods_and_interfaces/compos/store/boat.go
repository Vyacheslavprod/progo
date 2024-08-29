package store

type Boat struct {
	*Product
	Capasity  int
	Motorized bool
}

func NewBoat(name string, price float64, capasity int, motorized bool) *Boat {
	return &Boat{
		NewProduct(name, "Watersports", price),
		capasity,
		motorized,
	}
}
