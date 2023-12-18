package main

import (
	"fmt"
	"time"
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

	// ch <- "hello" // panic: send to closed channel

	values := []int{15, 8, 42, 16, 4, 23, 1}
	fmt.Println(sleepSort(values))
}

/*
For every value "n" in values, spin a goroutine that will
- slee "n" milliseconds
- send "n" to the channel

In the function bory, collect the values from the channel to a slice and return it
* this is not a really sort algorithm, but it's a good example for channels
*/
func sleepSort(values []int) []int {
	ch := make(chan int)
	for _, n := range values {
		n := n
		go func() {
			time.Sleep(time.Duration(n) * time.Millisecond)
			ch <- n
		}()
	}

	var out []int
	// for i := 0; i < len(values); i++ {
	for range values {
		n := <-ch
		out = append(out, n)
	}
	return out
}

/*
	Channel semantids:

- send & receive are blocking operations
- send & receive are atomic operations
- buffered channels are non-blocking  send & receive
- receive from closed channel is zero value without blocking
- send to closed channel panics
- closing a closed channel will panic
- send/receive to a nil channel blocks forever

See also https://www.353solutions.com/channel-semantics

Amdahl's law: https://en.wikipedia.org/wiki/Amdahl%27s_law
*/
func shadowExample() {
	n := 7
	{
		n := 2 // from here to } this is "n"
		// n = 2 // outer n
		fmt.Println("inner", n)
	}
	fmt.Println("outer", n)
}
