package main

import "fmt"

// Pointers

func main() {
	var p *int32 = new(int32)
	var i int32 

	showValues(&i,p)

	p = &i
	*p = 1

	showValues(&i,p)

	fmt.Println("\n----------------")

	var thing1 = [5]float64{1,2,3,4,5}
	fmt.Printf("The memory location of the thing1 array is %p", &thing1)
	var result [5]float64 = square(&thing1)
	fmt.Printf("\nThe result is: %v", result)
	fmt.Printf("\nThe value of thing1 is: %v", thing1)
}

func showValues(ipointer *int32, p *int32) {
	fmt.Println("\n----------------")
	fmt.Printf("The value p points to is: %v", *p)
	fmt.Printf("\nThe memory location p points to is: %v", p)
	fmt.Printf("\nThe memory location p is ar is: %v", &p)
	fmt.Printf("\nThe value i is at is : %v", ipointer)
	fmt.Printf("\nThe value i has the value of : %v", *ipointer)
}

func square(thing2 *[5]float64) [5]float64{
	fmt.Printf("\nThe memory location of the thing2 array is %p", thing2)
	for i := range thing2{
		thing2[i] = thing2[i]+thing2[i]
	}
	return *thing2
}