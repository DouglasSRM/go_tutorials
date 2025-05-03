package main

import (
	"fmt"
	"strings"
)

//strings, runes, concatenation

func main() {
	var myString = "résumé"
	var indexed = myString[1]
	fmt.Printf("%v, %T\n", indexed, indexed)
	for i, v := range myString {
		fmt.Println(i, v)
	}
	fmt.Printf("The length of 'myString' is %v\n\n", len(myString))

	var myString2 = []rune("résumé")
	var indexed2 = myString2[1]
	fmt.Printf("%v, %T\n", indexed2, indexed2) // runes are just an alias to int32 ([]int32("résumé") would work the same way)
	for i, v := range myString2 {
		fmt.Println(i, v)
	}
	fmt.Printf("The length of 'myString2' is %v\n\n", len(myString2))

	var myRune = 'a'
	fmt.Printf("myRune = %v\n", myRune)

	//var strSlice = []string{"s","t","r","i","n","g"}
	//var conStr = ""
	//for i := range strSlice {
	//	conStr += strSlice[i]
	//}
	//fmt.Printf("\n%v", conStr)

	//wont work
	//conStr[0] = 'a'

	var strSlice = []string{"s","t","r","i","n","g"}
	var strBuilder strings.Builder
	for i := range strSlice {
		strBuilder.WriteString(strSlice[i])
	}
	var conStr = strBuilder.String()
	fmt.Printf("\n%v", conStr)
}