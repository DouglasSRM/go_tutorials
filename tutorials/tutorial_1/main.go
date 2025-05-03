package main

import (
	//"errors"
	"fmt"
	"time"
	//"unicode/utf8"
)

//basic types, functions, arrays, slices, maps

/*
func main1(){
	var intNum int = 32767
	fmt.Println(intNum)

	var floatNum float64 = 12345678.9
	fmt.Println(floatNum)


	var myString string = "Hello World"
	fmt.Println(utf8.RuneCountInString(myString))

	var myRune rune = 'a'
	fmt.Println(myRune)

	var myBoolean bool = false
	fmt.Println(myBoolean)

	//Default for all types of numbers + runes is 0
	// for strings its ''
	// And for booleans its false!

	var defaultInt uint32
	fmt.Println(defaultInt)
	var defaultRune rune
	fmt.Println(defaultRune)
	var defaultString string
	fmt.Println(defaultString)
	var defaultBoolean bool
	fmt.Println(defaultBoolean)

	myVar := "Text!" //:= my beloved can be used to replace var!
	fmt.Println(myVar)

	var1, var2 := 1,2 //initialize multiple variables in the same line!
	fmt.Println(var1, var2)

	const myConst = "constant" //constants must me initialized with a value, once they cant be reasigned
	fmt.Println(myConst)

	const pi float32 = 3.14159
	fmt.Println(pi)
}

func main2(){
	var printValue string = "hi!"
	printMe(printValue)

	var numerator int = 11
	var denominator int = 0
	var result, remainder, err = intDivision(numerator, denominator)

	switch{
		case (err != nil):
			fmt.Println(err.Error())
		case (remainder == 0):
			fmt.Printf("The result of the integer division is %v", result)
		default:
			fmt.Printf("The result of the integer division is %v with remainder %v",result,remainder)
	}

	switch remainder {
	case 0:
		fmt.Printf("The division was exact!")
	case 1,2:
		fmt.Printf("The division was close!")
	default:
		fmt.Printf("The division was not close!")
	}
}

func printMe(printValue string){
	fmt.Println(printValue)
}

func intDivision(numerator int, denominator int) (int, int, error) {
	var err error
	if (denominator == 0) {
		err = errors.New("cannot divide by zero")
		return 0, 0, err
	}
	var result int = numerator/denominator
	var remainder int = numerator%denominator
	return result, remainder, err
}

*/

func main(){
	/*
	var intArr [3]int32
	intArr[1] = 123
	fmt.Println(intArr[0])
	fmt.Println(intArr[1:3])

	//memory adresses in arrays are continuous, which makes it faster to compile :)
	fmt.Println(&intArr[0])
	fmt.Println(&intArr[1])
	fmt.Println(&intArr[2])
	*/

	//var intArr [3]int32 = [3]int32{1,2,3}
				//this is still an array of fixed size 3!
	intArr := [...]int32{1,2,3} //fancy
	fmt.Println(intArr)

	//by omitting the length value, we have a slice 
	var intSlice []int32 = []int32{4,5,6}
	fmt.Printf("The length is %v with capacity %v\n", len(intSlice), cap(intSlice))
	fmt.Println(intSlice)
	intSlice = append(intSlice, 7)
	fmt.Printf("The length is %v with capacity %v\n", len(intSlice), cap(intSlice))
	fmt.Println(intSlice)

	//var intSlice2 []int32 = []int32{8,9}
	intSlice = append(intSlice, []int32{8,9}...)
	fmt.Println(intSlice)

	//This makes it possible to declare both the length and the capicity of the slice, otherwise
	// the capacity would be equal to the length ( make([]int32,3) -> length 3, capacity 3)
	var intSlice3 []int32 = make([]int32, 3, 8)
	fmt.Printf("The length is %v with capacity %v\n", len(intSlice3), cap(intSlice3))

	var myMap map[string]uint8 = make(map[string]uint8)
	fmt.Println(myMap)

	var myMap2 = map[string]uint8{"Douglas":20, "Jorge":32}
	fmt.Println(myMap2) 
	fmt.Println(myMap2["Douglas"])

	//if you try to get a value from an unsigned key you will get the default value back
	fmt.Println(myMap2["Jubileu"]) //Returns 0 which is default uint8
	// this means that the map will always return something even if the key doesnt "exists"

	//BUT maps can return a second value indicating if the informed key was encountered in the map or not ;)
	var age, ok = myMap2["Jubileu"]
	fmt.Println(age, ok)

	if ok {
		fmt.Printf("The age is %v\n", age)
	} else {
		fmt.Printf("Invalid Name\n")
	}

	//so in this case its true
	age, ok = myMap2["Douglas"]
	fmt.Println(age, ok)

	// delete(myMap2,"Jorge") to delete

	//iterating through maps
	for name, age := range myMap2 {
		fmt.Printf("Name: %v, Age:%v \n", name, age)
	}

	//iterating through arrays/slices is similar
	for i, v := range intArr{
		fmt.Printf("Index: %v, Value:%v \n", i, v)
	}

	var i int = 1
	//for (i <= 10) {
	//	fmt.Println(i)
	//	i++
	//}
	// OR
	for {
		if (i >= 10) {
			break
		}
		fmt.Println(i)
		i++
	}
	// OR
	for i := 1; i<=10; i++ {
		fmt.Println(i)
	}

	testTime()
}

func testTime(){
	const n int = 1000000
	var testSlice = []int{}
	var testSlice2 = make([]int, 0, n)

	fmt.Printf("Total time without preallocation: %v\n", timeLoop(testSlice, n))
	fmt.Printf("Total time with preallocation: %v\n", timeLoop(testSlice2, n))
	//about 6 times faster here :o
}

func timeLoop(slice []int, n int) time.Duration {
	var t0 = time.Now()
	for (len(slice) < n) {
		slice = append(slice, 1)
	}
	return time.Since(t0)
}