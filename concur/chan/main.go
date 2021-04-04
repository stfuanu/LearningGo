package main

import (
	"fmt"
	"strconv"
	"time"
)

// channel is a way for GoRoutines to communicate with each other .
// basically a tunnel/pipe though which we can receive or send messages .
// chan have type of data string , int or even chan .

func main() {

	// make a channel using make function
	CHANNELVAR := make(chan string)

	// go rountine function
	go count("FIRST GOROUTINE", CHANNELVAR)

	// recieve message/data from channel here [ONE AT A TIME] from tunnel CHANNELVAR
	// use forloop to overcome this if not then will receive only one and shut .
	// send and receive work together , both should be ready while any one of them ask to send/receive , which blockes their go routine .

	//------------------------------------------------------------------------------------------------------------

	// 1). EXAMPLE 1 :
	//fmt.Print("\n\n THIS PRINTs ONE LINE cuz sender sends one and then another when it evaluates it [it forloop goddammit : ]\n\n")
	//messagevar := <-CHANNELVAR
	//fmt.Println(messagevar)

	//------------------------------------------------------------------------------------------------------------
	// 2). EXAMPLE 2 :

	//fmt.Print("\n\n Uses for loop to print all results one after the another , as sender sends it , it receives and prints it \n\n")
	// In Go , there's no while , modified for loop is used in it's place , like this one
	//for {
	//	messagevar, open := <-CHANNELVAR

	// this break statement is important because , when sender won't send anything beause it's closed , it would print empty stuff .
	// this open variable has a value then it means it's open if not then it's closed , so it will break then .
	//	if !open {
	//		break
	//	}
	//	fmt.Println(messagevar)

	//}
	//------------------------------------------------------------------------------------------------------------
	// 3). EXAMPLE 3 :

	fmt.Print("\n\n A nicer way to do EXAMPLE 2 , using range Iterating over range of a channel , now we won't need openvar or while/for loop  \n\n")
	// Iterating over range of a channel .
	for x := range CHANNELVAR {
		fmt.Println(x)
	}

	//------------------------------------------------------------------------------------------------------------

	// close the channel [THERE1] , to tell chan to stop sending .
	// or will cause fatal error because RECEIVER is waiting or data , SENDER gate is open but forloop (any action) is over ,
	// Go detects it at runtime(not compile) , and understands that Go routine aren't making any progress and NOTHING WILL be received becoz  nothing left to send hence error.
	// Since channel is string type , and if sends nothing becos action(forloop in this case) is over  , so empty!=string , hence error
	// don't close pre-maturely , causes panic .

	// COUNT FUNCTION IS FINISHED AND THERE IS NO-ONE TO TELL RECEIVER (main function) THAT COUNT FUNCTION IS FINSIHED ,
	// SO PROGRAM WILL NEVER TERMINATE , BECAUSE main function (receiver) WILL KEEP WAITING FOR MESSAGE
	// IF close() function doesn't close the sender channel , indicating that it's OVER/COMPLETE. , close(<CHANNELVAR>) DOES THAT WORK , at the end of function .
	// RECEIVER CAN'T CLOSE CHANNEL BECOZ IT DOESN'T KNOW WHEN COUNT FUNTION WILL COMPLETE/OVER .

}

func count(x string, CHANNELVAR chan string) {
	for a := 0; a < 15; a++ {

		result := strconv.Itoa(a) + " is with " + x
		// strconv package includes functions to do stuff [TO/FROM String]
		// Itoa function is [Integer to ASCII] (idiots can't get a simple name)
		//also can use fmt.SprintF() , but idk how , check params allowed

		// use channel instead to print on terminal
		// like this
		CHANNELVAR <- result
		time.Sleep(time.Millisecond * 50)
	}
	//[HERE1]
	close(CHANNELVAR)

}

// SENDING AND RECEIVING THOUGH CHANNEL IS BLOCKING Operations of go routine

// When you try to recieve something , you need to have a value/data at sender(c <- x)

// if you have two go routines and 2nd one is receiving addition result of x+y , and your compiler reach (result := CHANNELVAR) in 2nd channel (receiver) .
// then it will stop there and wait for result which will be sent by sender (1st go routine) , doesn't matter how much time it takes .
// Thus blocking the go routine because it can't bypass it  and wait for result from sender .

// Similarly if our compiler reaches sender (CHANNELVAR <- result ) , so it means it has result and waiting for
