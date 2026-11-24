package golangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

func logJson(data interface{}) {
	byte, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(byte))
}

func TestEncodeJson(t *testing.T) {
	logJson("Udin")
	logJson(1)
	logJson(true)
	logJson([]string{"Udin", "Dindun", "Simalakadun"})
}