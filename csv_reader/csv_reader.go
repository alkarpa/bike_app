package csvreader

import (
	"os"
	"strings"
)

func ReadFromCSV(path string) ([]string, [][]string, error) {
	dat, err := os.ReadFile(path)
	if err != nil {
		return nil, nil, err
	}

	//fmt.Printf("Reading file %s\n", path)

	lines := strings.Split(string(dat), "\n")
	keys := strings.Split(lines[0], ",")

	//fmt.Printf("There are %d lines of data\n", len(lines)-1)

	ret := make([][]string, 0)

	const title_line_index = 0
	for i, line := range lines {
		if i == title_line_index {
			continue
		}

		// skip lines with missing data
		if strings.Contains(line, ",,") || strings.Contains(line, ",\n") {
			continue
		}
		values := strings.Split(line, ",")

		// skip lines if the lengths don't match
		if len(values) > len(keys) {
			continue
		}

		ret = append(ret, values)
	}
	//fmt.Printf("ret map %s", ret)

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
