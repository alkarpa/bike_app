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

	/*
		for i, key := range keys {
			fmt.Printf("key %d %s\n", i, key)
		}
	*/

	for i, line := range lines {
		if i == 0 {
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
