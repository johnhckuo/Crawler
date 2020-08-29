package warframe

type Riven struct {
	Positives map[string]string
	Negative  map[string]string
	Seller    *string
	Price     *string
	Name      *string
}

type Crawler interface {
	GetRivenByWeapon(*string) (*string, error)
	GetRivenByStats([]string) (*string, error)
}
