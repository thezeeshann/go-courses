package main

import (
	"fmt"
	"time"
)

func printMessage(text string, channel chan string) {
	for i := 0; i < 10; i++ {
		fmt.Println(text)
		time.Sleep(800 * time.Millisecond)
	}
	channel <- "Done" // sending data into channel
}

func main() {

	var channel chan string

	go printMessage("GO is great", channel)
	message := <-channel // we are waiting for data form channel
	fmt.Println(message)
	// go printMessage("frontend master is rock")
	// go printMessage("We miss classes")
	// time.Sleep(time.Minute)

}
