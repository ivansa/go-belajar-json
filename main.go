package main

import (
	"fmt"
	"go-belajar-json/usecase"
)

func main() {
	usecase.PrimitiveJsonEncode()
	usecase.MapJsonEncode()
	usecase.StructJsonEncode()

	fmt.Println()
	fmt.Println("===========================")
	fmt.Println()

	usecase.PrimitiveJsonDecode()
	usecase.MapJsonDecode()
	usecase.StructJsonDecode()
	usecase.ArrayJsonDecode()
	usecase.StructArrayJsonDecode()
}
