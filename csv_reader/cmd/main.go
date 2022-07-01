package main

import (
	"fmt"
	"time"

	csvreader "alkarpa.fi/csv_reader"
)

func main() {

	fmt.Println(".csv reader manual benchmark")
	start := time.Now()

	csvreader.ReadFromCSV("../2021-05.csv")

	end := time.Now()
	diff := end.Sub(start)
	fmt.Printf("Took %s \n", diff)
}
