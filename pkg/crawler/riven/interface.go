package riven

type Crawler interface {
	GetRivenByWeapon(*string) (*string, error)
	GetRivenByStats([]string) (*string, error)
}
