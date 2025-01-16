package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	

	file, err := os.Open("test.csv")
	if err != nil {
		panic(err)
	}

	csvReader := csv.NewReader(bufio.NewReader(file))
	records, err := csvReader.ReadAll()
	
	if err != nil {
		panic(err)
	}

	
	for record := range records {

		for _, field := range records[record] {
			fmt.Println(field)
		}
	}

}