package bike_app_be

import (
	"strconv"
	"strings"
)

type Station struct {
	Id         int                          `db:"id" json:"id"`
	Operator   string                       `db:"operator" json:"operator"`
	Capacity   int                          `db:"capacity" json:"capacity"`
	X          float64                      `db:"x" json:"x"`
	Y          float64                      `db:"y" json:"y"`
	LangFields []StationLangField           `json:"-"`
	Text       map[string]map[string]string `json:"text"`
}

type StationLangField struct {
	Lang  string
	Key   string
	Value string
}

func getStationLangFields(row []string, index_map map[string]int) []StationLangField {
	fields := [][3]string{
		{"name", "fi", "Nimi"},
		{"name", "se", "Namn"},
		{"address", "fi", "Osoite"},
		{"address", "se", "Adress"},
		{"city", "fi", "Kaupunki"},
		{"city", "se", "Stad"},
	}
	langfields := make([]StationLangField, 0, len(fields))
	for _, field := range fields {
		slf := NewStationLangFieldFromRow(
			field[0], field[1], row, field[2], index_map,
		)
		if strings.TrimSpace(slf.Value) != "" {
			langfields = append(langfields, slf)
		}
	}
	return langfields
}

func BuildTextFromLangFields(lf []StationLangField) map[string]map[string]string {
	m := make(map[string]map[string]string)
	for _, field := range lf {
		if _, ok := m[field.Lang]; !ok {
			m[field.Lang] = make(map[string]string)
		}
		m[field.Lang][field.Key] = field.Value
	}
	return m
}

func NewStationsFromDataSlice(keys []string, data [][]string) [](*Station) {
	index_map := make(map[string]int)

	wanted_keys := map[string]struct{}{
		"ID":         {},
		"Operaattor": {},
		"Kapasiteet": {},
		"x":          {},
		"y":          {},
		"Nimi":       {}, "Namn": {},
		"Osoite": {}, "Adress": {},
		"Kaupunki": {}, "Stad": {},
	}

	for i, key := range keys {
		if _, ok := wanted_keys[key]; ok {
			index_map[key] = i
		}
	}
	//fmt.Println(index_map)

	stations := make([](*Station), 0, len(data))
	for _, row := range data {
		id, _ := strconv.Atoi(row[index_map["ID"]])
		cap, _ := strconv.Atoi(row[index_map["Kapasiteet"]])
		x, _ := strconv.ParseFloat(row[index_map["x"]], 64)
		y, _ := strconv.ParseFloat(row[index_map["y"]], 64)

		langfields := getStationLangFields(row, index_map)

		station := &Station{
			Id:         id,
			Operator:   row[index_map["Operaattor"]],
			Capacity:   cap,
			X:          x,
			Y:          y,
			LangFields: langfields,
		}
		//fmt.Println(station)
		stations = append(stations, station)
	}

	return stations
}

func NewStationLangFieldFromRow(
	key string, lang string,
	row []string,
	csv_key string, im map[string]int,
) StationLangField {
	return StationLangField{
		Key:   key,
		Lang:  lang,
		Value: row[im[csv_key]],
	}
}

type StationService interface {
	GetAll() ([]*Station, error)
	GetDetails(int) (*Stats, error)
}

type Stat_count_avg struct {
	Count            int               `json:"count"`
	Average_distance float64           `json:"average_distance"`
	Top_connections  []Stat_connection `json:"top_connections"`
}
type Stat_connection struct {
	Station_id int `json:"station_id"`
	Count      int `json:"count"`
}

type Stats struct {
	Departing map[string]*Stat_count_avg
	Returning map[string]*Stat_count_avg
}
