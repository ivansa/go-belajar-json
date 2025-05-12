package usecase

import (
	"encoding/json"
	"fmt"
	"os"
)

func StreamDecoder() {
	reader, _ := os.Open("./json/sample.json")
	decoder := json.NewDecoder(reader)

	profile := map[string]interface{}{}
	decoder.Decode(&profile)

	fmt.Printf("Stream from fileJson and decode : %v \n", profile)
}

func StreamEncoder() {
	writer, _ := os.Create("./json/output.json")
	encoder := json.NewEncoder(writer)

	profile := map[string]interface{}{
		"name":   "ivan",
		"status": "active",
		"age":    20,
	}

	encoder.Encode(profile)
}
