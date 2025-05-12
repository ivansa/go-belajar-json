package usecase

import (
	"encoding/json"
	"fmt"
)

func PrimitiveJsonDecode() {

	var name string
	json.Unmarshal([]byte(`"Ivans"`), &name)

	var married bool
	json.Unmarshal([]byte(`true`), &married)

	var age int8
	json.Unmarshal([]byte(`20`), &age)

	fmt.Printf("Decode Primitive String : %s \n", name)
	fmt.Printf("Decode Primitive Bool : %t \n", married)
	fmt.Printf("Decode Primitive Int : %d \n", age)
}

func MapJsonDecode() {
	jsonMap := map[string]interface{}{}
	json.Unmarshal([]byte(`{"name":"ivan","age":20,"married":true,"hobbies":["Game","Football","Coding"]}`), &jsonMap)

	fmt.Printf("Decode Map Json : %v \n", jsonMap)
	fmt.Printf("Decode Map Json hobbies : %v \n", jsonMap["hobbies"])
}

func StructJsonDecode() {
	type Person struct {
		Name    string   `json:"name"`
		Age     int      `json:"age"`
		Married bool     `json:"married"`
		Hobbies []string `json:"hobbies"`
	}

	var person Person
	json.Unmarshal([]byte(`{"name":"ivan","age":20,"married":true,"hobbies":["Game","Football","Coding"]}`), &person)

	fmt.Printf("Decode Struct Json : %v \n", person)
	fmt.Printf("Decode Struct Json Hobbies : %v \n", person.Hobbies)
}

func ArrayJsonDecode() {
	type Person struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	var persons []Person
	json.Unmarshal([]byte(`[{"name":"ivan","age":20},{"name":"ard","age":21}]`), &persons)

	fmt.Printf("Decode Array Json : %v \n", persons)
	fmt.Printf("Decode Array Json 0 : %v \n", persons[0])
	fmt.Printf("Decode Array Json 1 : %v \n", persons[1])
}

func StructArrayJsonDecode() {
	type Job struct {
		Name string `json:"name"`
		Desc string `json:"desc"`
	}
	type Person struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
		Jobs []Job  `json:"jobs"`
	}

	var person Person
	json.Unmarshal([]byte(`{"name":"ivan","age":20,"jobs":[{"name":"programmer","desc":"programming"},{"name":"designer","desc":"designing"}]}`), &person)

	fmt.Printf("Decode Struct Array Json : %v \n", person)
	fmt.Printf("Decode Struct Array Json Jobs : %v \n", person.Jobs)
}
