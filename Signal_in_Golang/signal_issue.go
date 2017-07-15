package main

import (
    "fmt"
    "os"
    "os/signal"
    "syscall"
)

func main() {
    c := make(chan os.Signal, 1)
    signal.Notify(c, syscall.SIGUSR1)
    go func() {
         syscall.Kill(syscall.Getpid(), syscall.SIGUSR1)
    }()
    s := <-c
    fmt.Println("Got signal:", s)
}
