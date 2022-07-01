package csvreader

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(".csv reader")
}

func ReadFromCSV(path string) (map[string][]string, error) {
	dat, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(dat), "\n")
	keys := strings.Split(lines[0], ",")

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
		for j, val := range values {
			ret[keys[j]] = append(ret[keys[j]], val)
		}
	}
	//fmt.Printf("ret map %s", ret)

	return ret, nil
}
