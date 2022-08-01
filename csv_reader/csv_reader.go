package csvreader

import (
	"fmt"
	"os"
	"strings"
)

func getKeys(first_line string) []string {
	// remove BOM if present
	bytes := []byte(first_line)
	var start_index int

	if bytes[0] == '\xEF' && bytes[1] == '\xBB' && bytes[2] == '\xBF' {
		start_index = 3
	} else {
		start_index = 0
	}
	line := string(bytes[start_index:])

	keys := parseLineToValues(line) // strings.Split(line, ",")
	for i, key := range keys {
		keys[i] = strings.TrimSpace(key)
	}
	return keys
}

func readQuotedValue(line string) (string, string) {
	return readValue(line[1:], "\",")
}
func readUnquotedValue(line string) (string, string) {
	return readValue(line, ",")
}
func readValue(line string, value_ender string) (string, string) {
	ender_index := strings.Index(line, value_ender)
	if ender_index != -1 {
		return line[:ender_index], strings.TrimSpace(line[ender_index+len(value_ender):])
	}
	return line, ""
}

func parseLineToValues(line string) []string {
	const q = "\""
	quote_index := strings.Index(line, q)
	if quote_index != -1 {
		//fmt.Printf("q: %s\n", line)
		values := make([]string, 0)
		for rest_of_line := line[:]; len(rest_of_line) > 0; {

			rqi := strings.Index(rest_of_line, q)
			if rqi == 0 {
				val, rest := readQuotedValue(rest_of_line)
				//fmt.Printf("%s\n'%s';'%s'\n", rest_of_line, val, rest)
				values = append(values, val)
				rest_of_line = rest
			} else {
				val, rest := readUnquotedValue(rest_of_line)
				values = append(values, val)
				rest_of_line = rest
			}
		}
		return values
	} else {
		return strings.Split(strings.Replace(line, "\r", "", -1), ",")
	}
}

func ReadFromCSV(path string) ([]string, [][]string, error) {
	dat, err := os.ReadFile(path)
	if err != nil {
		return nil, nil, err
	}

	//fmt.Printf("Reading file %s\n", path)

	lines := strings.Split(string(dat), "\n")
	keys := getKeys(lines[0]) // strings.Split(lines[0], ",")

	//fmt.Printf("There are %d lines of data\n", len(lines)-1)

	ret := make([][]string, 0)

	const title_line_index = 0
	for i, line := range lines {
		if i == title_line_index {
			continue
		}
		// skip lines with missing data
		if len(line) == 0 || strings.Contains(line, ",\n") {
			last_line_index := len(lines) - 1
			if i != last_line_index {
				fmt.Printf("Skipping line %d: '%s' \n", i, line)
			}
			continue
		}

		values := parseLineToValues(line) // strings.Split(strings.Replace(line, "\r", "", -1), ",")

		// skip lines if the lengths don't match
		if len(values) > len(keys) {
			fmt.Printf("Shortening line %d: '%s', too many values %d vs expected %d\n", i, line, len(values), len(keys))
			values = values[:len(keys)]
		}

		ret = append(ret, values)
	}

	return keys, ret, nil
}

func GetKeyIndex(key string, keys []string) int {
	for i, val := range keys {
		if val == key {
			return i
		}
	}
	return -1
}
