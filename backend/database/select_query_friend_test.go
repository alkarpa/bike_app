package database

import (
	"strings"
	"testing"
)

func TestSelectQueryFriend_Pagination(t *testing.T) {
	t.Run("no page parameter makes LIMIT 100", func(t *testing.T) {
		parameters := map[string][]string{}
		sqf := &select_query_friend{}
		sqf.setupParamsPage(parameters)
		query_part := sqf.pagination
		expected := "LIMIT 100"
		if query_part != expected {
			t.Errorf("expected '%s', got '%s'", expected, query_part)
		}

	})
	t.Run("page=-1 makes LIMIT 100", func(t *testing.T) {
		parameters := map[string][]string{
			"page": {"-1"},
		}
		sqf := &select_query_friend{}
		sqf.setupParamsPage(parameters)
		query_part := sqf.pagination
		expected := "LIMIT 100"
		if query_part != expected {
			t.Errorf("expected '%s', got '%s'", expected, query_part)
		}

	})
	t.Run("page=2 makes LIMIT 100 OFFSET 200", func(t *testing.T) {
		parameters := map[string][]string{
			"page": {"2"},
		}
		sqf := &select_query_friend{}
		sqf.setupParamsPage(parameters)
		query_part := sqf.pagination
		expected := "LIMIT 100 OFFSET 200"
		if query_part != expected {
			t.Errorf("expected '%s', got '%s'", expected, query_part)
		}
	})
}
func TestSelectQueryFriend_Ordering(t *testing.T) {
	t.Run("no order makes empty string", func(t *testing.T) {
		parameters := map[string][]string{}
		sqf := &select_query_friend{}
		sqf.setupParamsOrdering(parameters)
		query_part := sqf.order_by
		expected := ""
		if query_part != expected {
			t.Errorf("expected '%s', got '%s'", expected, query_part)
		}
	})
	t.Run("order=return makes ORDER BY `return`", func(t *testing.T) {
		parameters := map[string][]string{
			"order": {"return"},
		}
		sqf := &select_query_friend{}
		sqf.setupParamsOrdering(parameters)
		query_part := sqf.order_by
		expected := "ORDER BY `return`"
		if query_part != expected {
			t.Errorf("expected '%s', got '%s'", expected, query_part)
		}
	})
	t.Run("order=return_station_desc makes ORDER BY `return_name` desc", func(t *testing.T) {
		parameters := map[string][]string{
			"order": {"return_station_desc"},
		}
		sqf := &select_query_friend{}
		sqf.setupParamsOrdering(parameters)
		query_part := sqf.order_by
		expected := "ORDER BY `return_name` desc"
		if query_part != expected {
			t.Errorf("expected '%s', got '%s'", expected, query_part)
		}
	})
	t.Run("order=departure_desc makes ORDER BY `departure` desc", func(t *testing.T) {
		parameters := map[string][]string{
			"order": {"departure_desc"},
		}
		sqf := &select_query_friend{}
		sqf.setupParamsOrdering(parameters)
		query_part := sqf.order_by
		expected := "ORDER BY `departure` desc"
		if query_part != expected {
			t.Errorf("expected '%s', got '%s'", expected, query_part)
		}
	})
}
func TestSelectQueryFriend_Language(t *testing.T) {
	t.Run("no language defaults to Finnish", func(t *testing.T) {
		parameters := map[string][]string{}
		sqf := &select_query_friend{}
		sqf.setupLang(parameters)
		query_part := sqf.language
		expected := "fi"
		if query_part != expected {
			t.Errorf("expected '%s', got '%s'", expected, query_part)
		}
	})
	t.Run("lang=se is Swedish", func(t *testing.T) {
		parameters := map[string][]string{
			"lang": {"se"},
		}
		sqf := &select_query_friend{}
		sqf.setupLang(parameters)
		query_part := sqf.language
		expected := "se"
		if query_part != expected {
			t.Errorf("expected '%s', got '%s'", expected, query_part)
		}
	})
	t.Run("lang=not_good defaults", func(t *testing.T) {
		parameters := map[string][]string{}
		sqf := &select_query_friend{}
		sqf.setupLang(parameters)
		query_part := sqf.language
		expected := "fi"
		if query_part != expected {
			t.Errorf("expected '%s', got '%s'", expected, query_part)
		}
	})
}

func TestSelectQueryFriend_addJoin(t *testing.T) {
	t.Run("Values to empty", func(t *testing.T) {
		sqf := &select_query_friend{}
		test_values := []string{"test_value", "test_second"}
		sqf.addJoin(test_values, "")
		for _, i := range []int{0, 1} {
			query_part := sqf.select_values[i]
			expected := test_values[i]
			if query_part != expected {
				t.Errorf("expected '%s', got '%s'", expected, query_part)
			}
		}
	})
	t.Run("Values to existing", func(t *testing.T) {
		sqf := &select_query_friend{
			select_values: []string{"existing"},
		}
		test_values := []string{"test_value", "test_second"}
		sqf.addJoin(test_values, "")
		for _, i := range []int{0, 1} {
			query_part := sqf.select_values[i+1]
			expected := test_values[i]
			if query_part != expected {
				t.Errorf("expected '%s', got '%s'", expected, query_part)
			}
		}
	})
	t.Run("Join station added to joins", func(t *testing.T) {
		sqf := &select_query_friend{}
		test_join := "JOIN station ON departure_station = station.id"
		sqf.addJoin(nil, test_join)
		query_part := sqf.joins[0]
		expected := test_join
		if query_part != expected {
			t.Errorf("expected '%s', got '%s'", expected, query_part)
		}

	})
}

func TestSelectQueryFriend_queries(t *testing.T) {
	t.Run("COUNT(*) from station, nothing more", func(t *testing.T) {
		sqf := &select_query_friend{
			select_values: []string{"COUNT(*)"},
			from_table:    "station",
		}
		query := strings.TrimSpace(sqf.buildQuery())
		expected := "SELECT COUNT(*) FROM station"
		if query != expected {
			t.Errorf("expected '%s', got '%s'", expected, query)
		}
	})
	t.Run("COUNT(*) from station, nothing more", func(t *testing.T) {
		sqf := &select_query_friend{
			select_values: []string{"COUNT(*)"},
			from_table:    "station",
		}
		query := strings.TrimSpace(sqf.buildQuery())
		expected := "SELECT COUNT(*) FROM station"
		if query != expected {
			t.Errorf("expected '%s', got '%s'", expected, query)
		}
	})
}
