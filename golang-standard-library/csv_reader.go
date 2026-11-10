package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"strings"
)

func main() {
	csvString := "dzikri,nur,akbar\n" +
	"khalid,shalahuddin,akbar\n" +
	"bilal,shalahuddin,akbar"

	reader := csv.NewReader(strings.NewReader(csvString))

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}

		fmt.Println(record)
	}
}