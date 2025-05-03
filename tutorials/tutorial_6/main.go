package main

import (
	"fmt"
	"math/rand"
	"time"
)

//Chanels

/*
func main() {
	var c = make(chan int, 5)
	go process(c)
	for i:= range c{
		time.Sleep(time.Second*1)
		fmt.Println(i)
	}
}

func process(c chan int) {
	defer close(c) //defer means this will be executed right before the function exits
	for i:=0; i<5; i++{
		c <- i
	}
	fmt.Println("Exiting process")
}
*/

var MAX_CHICKEN_PRICE float32 = 5
var MAX_TOFU_PRICE float32 = 3
func main() {
	var chickenChannel = make(chan string)
	var tofuChannel = make(chan string)
	var websites = []string{"walmart.com", "cotsco.com", "wholefoods.com"}
	for i := range websites {
		go checkChickenPrices(websites[i], chickenChannel)
		go checkTofuPrices(websites[i], tofuChannel)
	}
	sendMessage(chickenChannel, tofuChannel)
}

func checkTofuPrices(website string, c chan string){
	for {
		time.Sleep(time.Second*1)
		var tofu_price = rand.Float32()*20
		if (tofu_price < MAX_TOFU_PRICE) {
			c <- website
			break
		}
	}
}
func checkChickenPrices(website string, chickenChannel chan string) {
	for {
		time.Sleep(time.Second*1)
		var chickenPrice = rand.Float32()*20
		if chickenPrice <= MAX_CHICKEN_PRICE {
			chickenChannel <- website
			break
		}
	}
}

func sendMessage(chickenChannel chan string, tofuChannel chan string) {
	select {
		case website := <-chickenChannel:
			fmt.Printf("\nText Sent: Found a deal on chicken at %v.", website)
		case website := <-tofuChannel:
			fmt.Printf("\nEmail Sent: Found a deal on tofu at %v.", website)
	}
}