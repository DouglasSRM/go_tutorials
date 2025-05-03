package main

import "fmt"

// Structs and interfaces

/*
type gasEngine struct {
	mpg     uint8
	gallons uint8
	ownerInfo owner
	owner // adds the fields directly
}

type owner struct {
	name string
}

func main() {
	//var myEngine gasEngine = gasEngine{mpg:25, gallons:15}
	var myEngine gasEngine = gasEngine{25, 15, owner{"Douglas"}, owner{"test"}} //both work
	myEngine.mpg = 20
	fmt.Println(myEngine.mpg, myEngine.gallons, myEngine.ownerInfo.name, myEngine.name)

	// this works, but isnt reusable
	var myEngine2 = struct {
		mpg uint8
		gallons uint8
	}{20,15}
}
*/

type gasEngine struct {
	mpg     uint8
	gallons uint8
}

type electricEngine struct {
	mpkw uint8
	kwh uint8
}

//signing this function to the gasEngine type, becomes a method of the type.
func (e gasEngine) milesLeft() uint8 {
	return e.gallons*e.mpg
}

func (e electricEngine) milesLeft() uint8 {
	return e.kwh*e.mpkw
}

func canMakeIt(e engine, miles uint8) {
	if miles <= e.milesLeft() {
		fmt.Println("You can make it there!")
	} else {
		fmt.Println("Need to fuel up first!")
	}
}

//this interface makes it so the canMakeIt function can take anything as a parameter as long as
// the type passed has a "milesleft" method with the signature speficied in the interface :)
type engine interface{
	milesLeft() uint8 // the signature
}

func main() {
	var myGasEngine gasEngine = gasEngine{25, 15}
	canMakeIt(myGasEngine, 50)

	var myElectricEngine electricEngine = electricEngine{25,15}
	canMakeIt(myElectricEngine, 150)
}