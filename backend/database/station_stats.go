package database

import (
	"fmt"

	"alkarpa.fi/bike_app_be"
)

// An interface to ease the dynamic calling of similar or conditional queries
// and to keep the repeat code to a minimum for maintainability.
type stat_query interface {
	total_query() string
	monthly_query() string
	make_query(*StationService, bool) error
}

// Station statistics need very similar queries where only the
// table names get swapped. For this reason the swappable parts
// are written as %[1]s and %[2]s in the templates for this function.
// The id should be written as %[3]d in the template.
func build_query(s string, a string, b string, id int) string {
	return fmt.Sprintf(s, a, b, id)
}

// Returns the monthly or total stat_query query based on the monthly bool.
func stat_monthly_or_total(q stat_query, monthly bool) string {
	if monthly {
		return q.monthly_query()
	} else {
		return q.total_query()
	}
}

// Helps reduce duplicate code with stat_query implementations.
type stat_query_helper struct {
	target *map[string]*bike_app_be.Stat_count_avg
	a      string
	b      string
	id     int
}

// Queries that return a ride count and an average of distances
type query_count_average struct {
	*stat_query_helper
}

func (q *query_count_average) total_query() string {
	query := "SELECT 'all', COUNT(*), AVG(distance) FROM ride WHERE %[1]s_station=%[3]d"
	return build_query(query, q.a, q.b, q.id)
}
func (q *query_count_average) monthly_query() string {
	query := "SELECT CONCAT(YEAR(`%[1]s`),'-',MONTH(`%[1]s`)) yearmonth, " +
		"COUNT(*), AVG(distance) FROM ride WHERE %[1]s_station=%[3]d " +
		"GROUP BY yearmonth"
	return build_query(query, q.a, q.b, q.id)
}
func (q *query_count_average) make_query(s *StationService, monthly bool) error {
	var query = stat_monthly_or_total(q, monthly)
	return s.getStatCountAVG(q.target, query)
}

// Queries that return a slice of station connections with ride counts
type query_top_connections struct {
	*stat_query_helper
}

func (q *query_top_connections) total_query() string {
	query := `
	SELECT 'all', %[2]s_station, COUNT(*) c
	FROM ride WHERE %[1]s_station = %[3]d
	GROUP BY %[2]s_station
	ORDER BY c DESC
	LIMIT 5`
	return build_query(query, q.a, q.b, q.id)
}
func (q *query_top_connections) monthly_query() string {
	query :=
		"SELECT yearmonth, %[2]s_station, c " +
			"FROM ( SELECT CONCAT(YEAR(`%[1]s`),'-',MONTH(`%[1]s`)) yearmonth, " +
			" %[2]s_station, COUNT(*) c," +
			"ROW_NUMBER() OVER (PARTITION BY yearmonth order by c desc) rn " +
			"FROM ride WHERE %[1]s_station = %[3]d " +
			"GROUP BY yearmonth, %[2]s_station " +
			"ORDER BY yearmonth, c DESC " +
			") a " +
			"WHERE rn <= 5;"
	return build_query(query, q.a, q.b, q.id)
}

func (q *query_top_connections) make_query(s *StationService, monthly bool) error {
	var query = stat_monthly_or_total(q, monthly)
	return s.getTopConnections(q.target, query)
}
