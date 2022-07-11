package csvreader

import (
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

	keys := strings.Split(line, ",")
	for i, key := range keys {
		keys[i] = strings.TrimSpace(key)
	}
	return keys
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
		if len(line) == 0 || strings.Contains(line, ",,") || strings.Contains(line, ",\n") {
			continue
		}
		values := strings.Split(strings.Replace(line, "\r", "", -1), ",") // TODO: test \r

		// skip lines if the lengths don't match
		if len(values) > len(keys) {
			continue
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
