package gofunctions

import (
	"encoding/csv"
	"fmt"
	"os"
)

func Createf(path string) [][]string {
	fil, err := os.Open(path)
	if err != nil {
		fmt.Println("error")
	}

	defer fil.Close()
	fileReader := csv.NewReader(fil)
	records, error := fileReader.ReadAll()

	if error != nil {
		fmt.Println(error)
	}
	return records
}
