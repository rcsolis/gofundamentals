package concurrency

import (
	"fmt"
	"runtime"
	"strconv"
	"sync"
	"time"
)

func SayHello(msg string) {
	fmt.Println("Named function:", msg)
}

func basis() {
	// Create and executes an anonymous function
	var msg string = "Hello Concurrency"
	// Invoque a function as a gorutine
	go SayHello(msg)
	// It has a dependecy with the external variable
	go func() {
		msg = "Change variable"
		fmt.Println("First Anonymous: ", msg)
	}()
	msg = "Message changed after gorutine"
	// Create an executes and anonymous function
	// without external variable dependency
	// Takes the value of the variable when its called
	// Passing by value
	go func(msg string) {
		fmt.Println("Second Anonymous: ", msg)
	}(msg)
	fmt.Println("Execution end, Final value: ", msg)
	// *** AVOID USE SLEEP
	time.Sleep(100 * time.Millisecond)
}

// Using wait groups
var wg = sync.WaitGroup{}

func waitGroupsSample() {
	var msg = "Using waitgroups!!"
	// Increments the counter of the goroutines that will be executed
	wg.Add(1)
	go func(msg string) {
		SayHello(msg)
		// Notify the ends of goroutine
		wg.Done()
	}(msg)
	// We could add more than one
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			SayHello("Executing the goroutine: " + strconv.Itoa(i))
			wg.Done()
		}(i)
	}
	// Wait until go routines ends
	wg.Wait()
	fmt.Println("Finishing wait groups execution.")
}

var wgMxt = sync.WaitGroup{}
var mtx = sync.RWMutex{}
var counter = 0

func incrementCounter() {
	counter++
	// Task done
	wgMxt.Done()
}

func printCounter() {
	fmt.Println("The counter is: ", counter)
	// Task done
	wgMxt.Done()
}
func incrementCounterMutex() {
	counter++
	// finish write
	mtx.Unlock()
	// Task done
	wgMxt.Done()
}

func printCounterMutex() {
	fmt.Println("The counter is: ", counter)
	// unlock mutex
	mtx.RUnlock()
	// Task done
	wgMxt.Done()
}

func mutexFunc() {
	// Use mutex to synchronize tasks and protect data for read/write access
	fmt.Println("Synchronization using Mutex!")
	fmt.Println("Without using mutex, gorutines try to executes as fast as posible")

	for i := 0; i < 10; i++ {
		wgMxt.Add(2)
		go printCounter()
		go incrementCounter()
	}
	wgMxt.Wait()
	fmt.Println("With mutex, we loose the paralelism or currency capabilities of gorrutines")
	fmt.Printf("Threads: %v\n", runtime.GOMAXPROCS(-1))
	runtime.GOMAXPROCS(100)
	fmt.Printf("New Threads: %v\n", runtime.GOMAXPROCS(-1))
	// RWMutex, has infinite number of readers, but only one writer.
	for i := 0; i < 10; i++ {
		wgMxt.Add(2)
		// Lock for read
		mtx.RLock()
		go printCounterMutex()
		// Lock to write
		mtx.Lock()
		go incrementCounterMutex()
	}
	wgMxt.Wait()

}

var wgChan = sync.WaitGroup{}

func channelsFunc() {

	fmt.Println("Channels for concurrency.")
	fmt.Println("Unbuffered channels (default)")
	// channels are strongly typed, we only can send or receive data of the declared type
	// By default, we work with unbuffered channels, only ONE message can be in the channel at a time
	// the routine/function execution are paused until the channel are free
	ch := make(chan int)
	wgChan.Add(2)
	// Receiving gorutine
	go func() {
		// Get data from channel
		i := <-ch
		fmt.Println("Data Received: ", i)
		wgChan.Done()
	}()
	// Sender go routine
	go func() {
		i := 80
		ch <- i
		i = 27
		fmt.Println("Data in Sender:", i)
		wgChan.Done()
	}()
	wgChan.Wait()
	// Work with multiple calls
	for j := 0; j < 6; j++ {
		wgChan.Add(2)
		go func() {
			rec := <-ch
			fmt.Println("Received: ", rec)
			wgChan.Done()
		}()
		go func(idx int) {
			ch <- idx * 2
			fmt.Println("Sending double of:", idx)
			wgChan.Done()
		}(j)
	}
	wgChan.Wait()
	// Declares to be only sender or receiver
	fmt.Println("Create exclusive senders or receivers")

	wgChan.Add(2)
	// Receiver (declares that has a channel parameter to only receive integers
	go func(ch <-chan int) {
		i := <-ch
		fmt.Println("From exclusive receiver: ", i)
		wgChan.Done()
	}(ch)
	// Sender (declares that has a channel parameter to only send integers
	go func(ch chan<- int) {
		fmt.Println("Sending 99")
		ch <- 99
		fmt.Println("Message sent!")
		wgChan.Done()
	}(ch)
	wgChan.Wait()

	// Buffered channels
	fmt.Println("Create Buffered Channels")
	// Create a channel with store capacity limited 10 messages
	bufferedChannel := make(chan int, 10)
	//declare one receiver
	wgChan.Add(2)
	go func(ch <-chan int) {
		// When we could receive multiple messages, we need to loop into the channel
		for msg := range ch {
			fmt.Println("Receive message from buffer:", msg)
		}
		wgChan.Done()
	}(bufferedChannel)

	go func(ch chan<- int) {
		for i := 10; i < 20; i++ {
			ch <- i
		}
		// We need to close the channel to send a signal that there is no more messages to be posted
		close(ch)
		wgChan.Done()
	}(bufferedChannel)
	wgChan.Wait()

	fmt.Println("Channels with select statement and signal channel")
	// Using select statement
	msgChannel := make(chan string, 5)
	// delcares a signal only channel
	// Using type of struct{} that is unique and requires CERO memory allocation but let us know
	// that a message was sent
	doneChannel := make(chan struct{})

	// receiver
	go func(msg <-chan string, done <-chan struct{}) {

		for {
			// Monitor channels
			// By default select its a blocking statement, if we want to exectute some code
			// and convert to a non blocking statement we need to add a default case
			select {
			// Receive from channel with data
			case data := <-msg:
				fmt.Println("-", data)
			// REceive from signal channel
			case <-done:
				break
			//default converts to nonblocking select and
			//executes when there is no messages available in the channels
			default:
				fmt.Println("Waiting for messages")
			}
		}
	}(msgChannel, doneChannel)

	// Sends multple messages
	go func(msg chan<- string, done chan<- struct{}) {
		// Sends messages
		msg <- "Top languages that I want to use"
		msg <- "Golang"
		msg <- "Rust"
		msg <- "C++"
		msg <- "Dart/Flutter"
		msg <- "Javascript/react"
		msg <- "NodeJs"
		msg <- "Python"
		// Send signal for break
		done <- struct{}{}
	}(msgChannel, doneChannel)

	fmt.Println("This is the last statement in channels function.")
}

func Init() {
	fmt.Println("*** Concurrency")
	basis()
	waitGroupsSample()
	mutexFunc()
	channelsFunc()
}
