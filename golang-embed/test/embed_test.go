package golang_embed

import (
	"embed"
	_ "embed"
	"fmt"
	"io/fs"
	"os"
	"testing"
)

//go:embed version.txt
var version string
func TestString(t *testing.T) {
	fmt.Println(version)
}

//go:embed image.png
var logo []byte
func TestByte(t *testing.T) {
	err := os.WriteFile("new-image.png", logo, fs.ModePerm)
	if err != nil {
		panic(err)
	}
}

//go:embed a.txt
//go:embed b.txt
//go:embed c.txt
var files embed.FS

func TestMultiplefile(t *testing.T) {
	a, _ := files.ReadFile("files/a.txt")
	b, _ := files.ReadFile("files/b.txt")
	c, _ := files.ReadFile("files/c.txt")

	fmt.Println(string(a))
	fmt.Println(string(b))
	fmt.Println(string(c))
}

//go:embed *.txt
var path embed.FS

func TestPathMatcher(t *testing.T) {
	dirEntries, _ := path.ReadDir("files")
	
	for _, entry := range dirEntries {
		if !entry.IsDir() {
			fmt.Println(entry.Name())
			file, _ := path.ReadFile("files/" + entry.Name())
			fmt.Println(string(file))
		}
	}
}