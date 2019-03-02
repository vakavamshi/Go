package main

import (
	"fmt"
	"time"
)

func main() {

	go Consumer1()	
           Consumer2()  


}
func Consumer1() {
	for i := 0; i < 3; i++ {
		fmt.Println("Consumer1,  ->", i)
		time.Sleep(time.Duration(5 * time.Millisecond))
	}
	}
	
func Consumer2() {
	for i := 0; i < 3; i++ {
		fmt.Println("Consumer2,  ->", i)
		time.Sleep(time.Duration(5 * time.Millisecond))
	}
	
	}

