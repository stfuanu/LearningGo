package main

import (
	"fmt"
	"time"
)

func main() {

	p1 := make(chan string)
	p2 := make(chan string)

	go func() {
		for {
			p1 <- "Process took 0.5 seconds"
			time.Sleep(time.Millisecond * 500)
		}
	}()

	go func() {
		for {
			p2 <- "Process took 2 seconds"
			time.Sleep(time.Second * 2)
		}
	}()

	// Example :1 : here p1 result is ready and sends it in 0.5 seconds but after than we have to wait for print p2 results to come in 2 seconds
	// So p1 on another cycle (while/for loop) takes 2 seconds delay .
	// imagine this in context for api calling , A slow endpoint will delay the fast one too which if called after it on loop.
	// To overcome this we use SELECT statement to print p1 and p2 which ever's result is received instantly .

	for i := 0; i < 5; i++ {
		mess := <-p1
		fmt.Println(mess)
		// Either save channel output to a var and then print it
		// OR receive the output directly into prinln function like below.
		// Later is better , cuz maybe it reduces time and resources , idk ??

		fmt.Println(<-p2)
	}

	// Example 1 : Output :

	//Process took 0.5 seconds
	//Process took 2 seconds
	//Process took 0.5 seconds
	//Process took 2 seconds
	//Process took 0.5 seconds
	//Process took 2 seconds
	//Process took 0.5 seconds
	//Process took 2 seconds
	//Process took 0.5 seconds
	//Process took 2 seconds
	//Process took 0.5 seconds
	//Process took 2 seconds

	// Example : 2
	// looping over a select statement
	for i := 0; i < 5; i++ {

		select {

		case res1 := <-p1:
			fmt.Println(res1)
		case res2 := <-p2:
			fmt.Println(res2)
		}
	}

	// Example 2 : Output :

	//Process took 2 seconds
	//Process took 0.5 seconds
	//Process took 0.5 seconds
	//Process took 0.5 seconds
	//Process took 0.5 seconds
	//Process took 2 seconds
	//Process took 0.5 seconds
	//Process took 0.5 seconds
	//Process took 0.5 seconds
	//Process took 0.5 seconds
	//Process took 2 seconds
	//Process took 0.5 seconds
	//Process took 0.5 seconds
	//Process took 0.5 seconds
	//Process took 0.5 seconds
	//Process took 2 seconds

}
