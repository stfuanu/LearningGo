// By default channels are _unbuffered_, meaning that they
// will only accept sends (`chan <-`) if there is a
// corresponding receive (`<- chan`) ready to receive the
// sent value. _Buffered channels_ accept a limited
// number of  values without a corresponding receiver for
// those values.

package main

import "fmt"

func main() {

	// Here we make a channel of strings buffering up to
	// 2 values.
	messages := make(chan string, 2)

	// if we don't use buffer here , then it will be stuck at sender and nothing can receive it beacause , only receriver in this code is in same main function which it can't jump .

	// Because this channel is buffered, we can send these
	// values into the channel without a corresponding
	// concurrent receive.
	messages <- "first"
	messages <- "second"
	//messages <- "third"

	// Later we can receive these two values as usual.
	fmt.Println(<-messages)
	fmt.Println(<-messages)
	//fmt.Println(<-messages)
}
