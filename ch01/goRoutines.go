package main

import (
	"fmt"
	"time"
)

var counter int = 0

func myPrint(start, finish int) {
	for i := start; i <= finish; i++ {
		//fmt.Print(i, " ")
		counter++
		fmt.Println(counter)
	}
	fmt.Println()
	time.Sleep(100000000 * time.Second)
}

func main() {
	for i := 0; i < 400000000000; i++ {
		go myPrint(i, 54000000000000)
	}
	time.Sleep(time.Second)
}
