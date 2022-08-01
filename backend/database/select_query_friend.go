package database

import (
	"fmt"
	"strconv"
	"strings"
)

// A struct to help organize query building for modifiable SELECTs.
type select_query_friend struct {
	language      string
	select_values []string
	from_table    string
	joins         []string
	order_by      string
	pagination    string
	values        []interface{}
}

// Sets up the pagination part of the query. Uses LIMIT and OFFSET
func (sqf *select_query_friend) setupParamsPage(parameters map[string][]string) {
	const page_size = 100
	default_limit := fmt.Sprintf("LIMIT %d", page_size)
	if page_params, found := parameters["page"]; found {
		page, err := strconv.Atoi(page_params[0])
		if err != nil || page < 1 {
			sqf.pagination = default_limit
			return
		}
		const page_limit_offset = "LIMIT %d OFFSET %d"
		sqf.pagination = fmt.Sprintf(page_limit_offset, page_size, page_size*page)
		return
	}
	sqf.pagination = default_limit
}

// Sets up the ORDER BY part of the query. Uses a whitelist for sanification.
func (sqf *select_query_friend) setupParamsOrdering(parameters map[string][]string) {
	if order_params, found := parameters["order"]; found {
		accepted_values := []string{
			"departure",
			"return",
			"departure_station",
			"return_station",
			"distance",
			"duration",
		}
		special_values := map[string]string{
			"departure_station": "departure_name",
			"return_station":    "return_name",
		}
		order_string := "ORDER BY %s"
		desc := "desc"
		orderings := make([]string, 0, len(order_params))
		for _, ordering := range order_params {
			parts := strings.Split(ordering, "_")
			for _, accepted_value := range accepted_values {
				q_value := parts[0]
				desc_pos := 1
				if len(parts) > 1 && parts[1] != desc {
					q_value += "_" + parts[1]
					desc_pos++
				}
				if q_value == accepted_value {

					if val, ok := special_values[q_value]; ok {
						q_value = val
					}
					query_ord := "`" + q_value + "`"
					if len(parts) == desc_pos+1 && parts[desc_pos] == desc {
						query_ord += " " + desc
					}
					orderings = append(orderings, query_ord)
					break
				}
			}
		}
		if len(orderings) > 0 {
			sqf.order_by = fmt.Sprintf(order_string, strings.Join(orderings, ","))
		}
	}
}

// For text joins and such. Defaults to Finnish. Uses a whitelist for sanification.
func (sqf *select_query_friend) setupLang(parameters map[string][]string) {
	accepted_languages := []string{"fi", "se"}
	if lang, found := parameters["lang"]; found {
		for _, al := range accepted_languages {
			if al == lang[0] {
				sqf.language = lang[0]
				return
			}
		}
	}
	sqf.language = "fi"
}

// Builds the completed query
func (sqf *select_query_friend) buildQuery() string {
	select_values := strings.Join(sqf.select_values, ",")
	inner_joins := strings.Join(sqf.joins, " ")

	query := fmt.Sprintf("SELECT %[1]s FROM %[2]s %[3]s %[4]s %[5]s",
		select_values,
		sqf.from_table,
		inner_joins,
		sqf.order_by,
		sqf.pagination,
	)
	return query
}

// Adds a join to the table and any related SELECT values
func (sqf *select_query_friend) addJoin(values []string, join string) {
	sqf.select_values = append(sqf.select_values, values...)
	sqf.joins = append(sqf.joins, join)
}

// Sets up the text search of the query. It's a front and back wild card text match and as such
// is quite slow. Only added to the query if the search parameter is found.
func (sqf *select_query_friend) setupTextSearch(parameters map[string][]string) {
	if search, search_found := parameters["search"]; search_found {
		sqf.addJoin(nil, " INNER JOIN "+
			"(SELECT DISTINCT id FROM station_lang_field WHERE station_lang_field.lang = '"+sqf.language+"' "+
			"AND station_lang_field.value LIKE CONCAT('%',?,'%') ) a "+
			"ON a.id IN (departure_station, return_station) ")
		sqf.values = append(sqf.values, search[0])
	}
}

// Builds a new select query friend from the parameters
func newRideSelectQueryFriend(parameters map[string][]string) *select_query_friend {

	sqf := &select_query_friend{
		select_values: []string{"departure", "`return`", "departure_station", "return_station", "distance", "duration"},
		from_table:    "ride",
	}
	sqf.setupLang(parameters)

	sqf.addJoin([]string{"slfds.value AS departure_name"},
		"INNER JOIN station_lang_field AS slfds ON slfds.id = departure_station AND slfds.lang='"+sqf.language+"' AND slfds.key='name'")
	sqf.addJoin([]string{"slfrs.value AS return_name"},
		"INNER JOIN station_lang_field AS slfrs ON slfrs.id = return_station AND slfrs.lang='"+sqf.language+"' AND slfrs.key='name'")

	sqf.setupTextSearch(parameters)
	sqf.setupParamsOrdering(parameters)
	sqf.setupParamsPage(parameters)
	return sqf
}
