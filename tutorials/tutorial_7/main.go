package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	var intSlice = []int{1, 2, 3}
	fmt.Println(sumSlice(intSlice))

	fmt.Println(isEmpty(intSlice))

	var float32Slice = []float32{1,2,3}
	fmt.Println(sumSlice(float32Slice))

	run()
}

func sumSlice[T int | float32 | float64](slice []T) T {
	var sum T
	for _, v := range slice {
		sum += v
	}
	return sum
}

func isEmpty[T any](slice []T) bool {
	return len(slice)==0
}

type contactInfo struct {
	Name string
	Email string
}

type purchaseInfo struct {
	Name string
	Price float32
	Amount int
}

func run() {
	var contacts []contactInfo = loadJSON[contactInfo]("./contactInfo.json")
	fmt.Printf("\n%+v", contacts)

	var purchases []purchaseInfo = loadJSON[purchaseInfo]("./purchaseInfo.json")
	fmt.Printf("\n%+v", purchases)
}

func loadJSON[T contactInfo | purchaseInfo](filePath string) []T{
	data, _ := os.ReadFile(filePath)

	var loaded = []T{}
	json.Unmarshal(data, &loaded)

	return loaded
}