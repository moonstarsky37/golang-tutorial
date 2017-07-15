package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"gopkg.in/yaml.v2"
)

// The config struct, having only a single field Name that will be printed
type config struct {
	Name string
}

// Context is a struct to wrap the config file and guard it with a mutex
type context struct {
	config
	sync.Mutex
}

// The global context variable
var ctx *context

func (c *context) Config() config {
	c.Lock()
	cnf := c.config
	c.Unlock()
	return cnf
}

// A function that reads the file "config.yml" and returns a pointer to an instance of the config struct.
func readConfig() config {
	content, err := ioutil.ReadFile("config.yml")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	var tmpCnf config
	err = yaml.Unmarshal(content, &tmpCnf)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	return tmpCnf
}

// That's the interesting part
func swapConfig(c chan os.Signal) {
	for {
		// This line will block until the signal is received
		<-c

		// Whenever the signal is received. Read the configuration file and swap the old config out.
		ctx.Lock()
		ctx.config = readConfig()
		ctx.Unlock()
	}
}

func main() {
	fmt.Printf("Process PID : %v\n", os.Getpid())

	ctx = &context{config: readConfig()}

	c := make(chan os.Signal, 1)
	// Send to channel c whenver you receive a SIGUSR1 signal.
	signal.Notify(c, syscall.SIGUSR1)
	go swapConfig(c)

	for {
		fmt.Println(ctx.Config().Name)
		time.Sleep(time.Second)
	}
}
