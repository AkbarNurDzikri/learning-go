package main

import (
	"encoding/csv"
	"os"
)

func main() {
	writer := csv.NewWriter(os.Stdout)

	_ = writer.Write([]string{"dzikri","nur","akbar"})
	_ = writer.Write([]string{"bilal","shalahuddin","akbar"})
	_ = writer.Write([]string{"khalid","shalahuddin","akbar"})

	writer.Flush()
}