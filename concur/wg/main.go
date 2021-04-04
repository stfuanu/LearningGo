package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup // counter name defined
	// increment the counter by 1 , to say that we have one go routine here .
	wg.Add(1)

	//count("samosa")

	// decrement the counter when goroutine finishes , means WAIT IS OVER .
	// put wg.Done() after function is called DONE , to tell that the funtion work is over

	// we can also create an anonymous funtion by[        go func(){CODE}()       ], this syntax creates a function and immediatly invokes it .
	// functions called from inside this will run as goroutine .
	go func() {
		count("inside anon ,SAMOSA and FIRST GOROUTINE")
		wg.Done() // this decrements the counter by 1
	}()

	go func() {
		count("another one  , runs concur[NOT PARALLELLY] with FIRST function (time b/w) SECOND GOROUTINE")
		wg.Done() // this decrements the counter by 1
	}()

	wg.Wait() // wait untill gorountine is over , OR till decrement is done , basically blocks func main() till wg.Done() is hit to decrement.
}

func count(x string) {
	for i := 0; i < 10; i++ {
		fmt.Println(i, "is", x)
		time.Sleep(time.Millisecond * 500)
	}

}
