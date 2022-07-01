package csvreader

import (
	"os"
	"strings"
)

func ReadFromCSV(path string) (map[string][]string, error) {
	dat, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	//fmt.Printf("Reading file %s\n", path)

	lines := strings.Split(string(dat), "\n")
	keys := strings.Split(lines[0], ",")

	//fmt.Printf("There are %d lines of data\n", len(lines)-1)

	ret := make(map[string][]string)

	const title_line_index = 0
	for i, line := range lines {
		if i == title_line_index {
			continue
		}
		if strings.Contains(line, ",,") || strings.Contains(line, ",\n") {
			//fmt.Printf("Line %d is missing data\n", i)
			continue
		}
		values := strings.Split(line, ",")

		if len(values) > len(keys) {
			continue
		}

		for j, val := range values {
			ret[keys[j]] = append(ret[keys[j]], val)
		}
	}
	//fmt.Printf("ret map %s", ret)

	return ret, nil
}
