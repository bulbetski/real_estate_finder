package types

type Location struct {
	Lat float64 `db:"latitude"`
	Lng float64 `db:"longitude"`
}
