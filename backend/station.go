package bike_app_be

type Station struct {
	Id   int    `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
}

type StationService interface {
	GetAll() ([]*Station, error)
}
