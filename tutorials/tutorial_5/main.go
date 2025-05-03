package main

import (
	"fmt"
	"sync"
	"time"
)

//GoRoutines

var m = sync.RWMutex{} //short for (read/write) mutual exclusion. 2 main methods are Lock() and Unlock()
var wg = sync.WaitGroup{} //works like a counter
var dbData = []string{"id1", "id2", "id3", "id4", "id5"}
var results = []string{}

func main() {
	t0 := time.Now()
	for i:=0; i<len(dbData); i++{
		wg.Add(1) // increments the counter before starting the goroutine
		go dbCall(i)
	}
	wg.Wait() // waits for the counter to go back down to 0, meaning all the tasks have completed
	fmt.Printf("\nTotal execution time: %v", time.Since(t0))
	fmt.Printf("\nThe results are: %v", results)

	t0 = time.Now()
	for i:=0; i<10000; i++{
		wg.Add(1) 
		go count()
	}
	wg.Wait()
	fmt.Printf("\nTotal execution time: %v", time.Since(t0))
}

func dbCall(i int) {
	// Simulate DB call delay
	var delay float32 = 2000
	time.Sleep(time.Duration(delay)*time.Millisecond)	
	save(dbData[i])
	log()
	wg.Done() // decrements the counter at the end of the goroutine
}

func save(result string){
	m.Lock() //locks it so no other routine can append to the results at the same time. this can be full locks (Lock()) and read locks (Rlock()).
	results = append(results, result)
	m.Unlock() //unlocks so the next routine can go on
}

func log(){
	m.RLock() // Read locks dont wait for other read locks to go on, only for full locks!
	fmt.Println("The result from the database is:", results)
	m.RUnlock()
}

func count(){
	var res int
	for i:=0; i<10000000; i++{
		res+=1
	}
	wg.Done()
}