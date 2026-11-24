package golangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Product struct {
	Id string `json:"id"`
	Name string `json:"name"`
	ImageURL string `json:"image_url"`
}

func TestJsonTagEncode(t *testing.T) {
	product := Product{
		Id: "P001",
		Name: "Apple Mac Book Pro",
		ImageURL: "http://example.com/image.png",
	}

	bytes, _ := json.Marshal(product)
	fmt.Println(string(bytes))
}

func TestJsonTagDecode(t *testing.T) {
	// saat decode (baca) golang mengabaikan huruf besar kecil
	// tapi untuk encode harus persis
	// namun untuk key ber-underscore harus sama ada underscore juga
	jsonString := `{
		"ID": "P001",
		"NAME": "Apple Mac Book Pro",
		"IMAGE_URL": "http://example.com/image.png"
	}`

	jsonBytes := []byte(jsonString)

	product := &Product{}

	json.Unmarshal(jsonBytes,product)
	fmt.Println(product)
}