package main

import (
	_ "fmt"
	"io/ioutil"
	"strings"
)

func read_filelist(filename_list string, ch1 chan func(string), ch2 chan int) {
	dat, err := ioutil.ReadFile(filename_list)
	if err != nil {
		panic(err)
	}
	filenames := strings.Split(string(dat), "\n")
	for i := 0; i < len(filenames); i++ {
		ch1 <- func(input_text string) {
			ioutil.WriteFile(filenames[i], []byte(input_text), 0644)
		}
		_ = <- ch2
	}
	close(ch1)
}

func main() {
	ch1 := make(chan func(string))
	ch2 := make(chan int)
	go read_filelist("writing.list", ch1, ch2)
	for f := range ch1 { //loop until ch1 is closed
		f("Hello File!")
		ch2 <- 1
	}
}
