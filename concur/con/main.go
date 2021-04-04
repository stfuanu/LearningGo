package main

import (
	"fmt"
	"time"
)

func main() {

	// calling a function without go prefix will run usual , first it will complete then next thing .
	// but with go routine , the function call works in background , much like screen . but like when if calling api then the we can use
	// time taken to call api and call another in meantime , thus concurrency
	go count("sleep")
	time.Sleep(time.Millisecond * 1000)
}

func count(x string) {
	for i := 0; i < 10; i++ {
		fmt.Println(i, "is", x)
		time.Sleep(time.Millisecond * 500)
	}
}
