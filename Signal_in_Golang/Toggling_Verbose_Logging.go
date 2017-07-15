package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

// The io.Writer that the prints will write to
var writeTo io.Writer = ioutil.Discard
var mutex sync.Mutex

func toggleOutput(c chan os.Signal) {
	for {
		// This line will block until the signal is received
		<-c
		fmt.Println(c)

		// Whenver the signal is received toggle the target writer. If it's stdout swap it with ioutil.Discard and vice versa.
		mutex.Lock()
		if writeTo == os.Stdout {
			writeTo = ioutil.Discard
		} else {
			writeTo = os.Stdout
		}
		mutex.Unlock()
	}
}

func main() {

	fmt.Printf("Process PID : %v\n", os.Getpid())

	c := make(chan os.Signal, 1)
	// Send to channel c whenver you receive a SIGUSR1 signal.
	signal.Notify(c, syscall.SIGUSR1)
	go toggleOutput(c)

	// An infinte loops that prints and increments a counter
	counter := 1
	for {
		mutex.Lock()
		fmt.Fprintln(writeTo, counter)
		mutex.Unlock()
		counter++
		time.Sleep(time.Second)
	}
}
