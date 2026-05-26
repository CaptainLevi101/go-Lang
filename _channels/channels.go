package main

import (
	"fmt"
	"math/rand/v2"
	"time"
	//"time"
)

func processNum(numChan chan int){
	for num:=range numChan{
		fmt.Println("Processing number",num);
		time.Sleep(time.Second*2);
	}
}

func main() {
	numChan:=make(chan int)
	go processNum(numChan)
	for{
		numChan<-rand.IntN((100))
	}
//	numChan<-5

	//time.Sleep(time.Second*2);
	// channels are used for communication between different goroutines
	// messageChan := make(chan string)
	// // channels are bloking in nature 
	// messageChan <- "ping" // sending data to channel
	// //	<-messageChan    // recieving data from the channel
	// msg := <-messageChan

//	fmt.Println(msg);

}