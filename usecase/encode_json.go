package usecase

import (
	"encoding/json"
	"fmt"
)

func PrimitiveJsonEncode() {
	jsonStr, _ := json.Marshal("Ivans")
	jsonBool, _ := json.Marshal(true)
	jsonNumber, _ := json.Marshal(2.0)
	jsonArr, _ := json.Marshal([]string{"ivan", "ard", "syah"})

	fmt.Printf("Primitive String : %s \n", jsonStr)
	fmt.Printf("Primitive Bool : %s \n", jsonBool)
	fmt.Printf("Primitive Number : %s \n", jsonNumber)
	fmt.Printf("Primitive Arr : %s \n", jsonArr)
}

func MapJsonEncode() {
	jsonMap := map[string]interface{}{
		"name":    "ivan",
		"age":     20,
		"married": true,
		"hobbies": []string{"Game", "Football", "Coding"},
	}

	jsonStr, _ := json.Marshal(jsonMap)
	fmt.Printf("Map Json : %s \n", jsonStr)
}

func StructJsonEncode() {
	type Person struct {
		Name    string            `json:"name"`
		Age     int               `json:"age"`
		Married bool              `json:"married"`
		Address map[string]string `json:"address"`
		Hobbies []string          `json:"hobbies"`
	}

	person := Person{
		Name:    "ivan",
		Age:     20,
		Married: true,
		Address: map[string]string{
			"street": "Jl. Merdeka",
			"city":   "Jakarta",
		},
		Hobbies: []string{"Game", "Football", "Coding"},
	}

	jsonStr, _ := json.Marshal(person)
	fmt.Printf("Struct Json : %s \n", jsonStr)
}
