package main

import (
	"fmt"
)

func main() {
	go fmt.Println("goroutine")
	fmt.Println("main routine")

	for i := 0; i < 3; i++ {

		// FIX 2: use a loop body variable
		i := i // i shadowing the outer i
		go func() {
			fmt.Println(i)
		}()
		/* FIX 1: use a local variable
		go func(n int) {
			fmt.Println(n)
		}(i)
		*/
		/*	BUG: all goroutines use the "i" variable from the same scope
			go func() {
				fmt.Println(i)
			}()*/
	}

	// time.Sleep(10 * time.Millisecond)

	ch := make(chan string)
	go func() {
		ch <- "hello" // send to channel
	}() // anonymous function
	msg := <-ch // receive from channel
	fmt.Println(msg)

	go func() {
		for i := 0; i < 3; i++ {
			// fmt.Println(i)
			msg := fmt.Sprintf("message #%d", i+1)
			ch <- msg
		}
		close(ch)
	}()

	for msg := range ch {
		fmt.Println("got: ", msg)
	}

	msg, ok := <-ch // ch is closed, msg is zero value
	fmt.Printf("closed channel: %#v (ok=%v)\n", msg, ok)

	/* Channel semantids:
	- send & receive are blocking operations
	- send & receive are atomic operations
	- receive from closed channel is zero value without blocking
	- send to closed channel panics
	- closing a closed channel will panic
	- send/receive to a nil channel blocks forever
	*/

	// ch <- "hello" // panic: send to closed channel
}
