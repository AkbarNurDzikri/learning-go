package main

import (
	"embed"
	"fmt"
	"io/fs"
	"os"
)

//go:embed test/version.txt
var version string

//go:embed test/image.png
var logo []byte

//go:embed test/*.txt
var path embed.FS

func main() {
	fmt.Println(version)

	err := os.WriteFile("new-image.png", logo, fs.ModePerm)
	if err != nil {
		panic(err)
	}

	dirEntries, _ := path.ReadDir("test")
	
	for _, entry := range dirEntries {
		if !entry.IsDir() {
			fmt.Println(entry.Name())
			file, err := path.ReadFile("test/" + entry.Name())
			if err != nil {
				panic(err)
			}

			fmt.Println(string(file))
		}
	}
}