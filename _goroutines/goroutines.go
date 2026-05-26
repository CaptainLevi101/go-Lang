package main

import (
	"fmt"
	"sync"
)

// here go routines deal with multi threading stuff, we gonna see how we gonna create this goroutine stuff
func task(id int,w* sync.WaitGroup){
	defer w.Done()
	// defer runs after the function completion 
	fmt.Println("Doing Task",id);
}
func main(){

	// now i want to run this task multiple times 
	var wg sync.WaitGroup
	for i:=0;i<10;i++{
		wg.Add(1)
		go task(i,&wg);
		// now this task is done in the specific order 
		// it means it is blocking 
		// we want that these task functions should run parallely 
		// to do this just write go in front of the function 
		// when we run this nothing is on hte screen 
		//why? golang put it in scheldular than schedular handles it 
		// when these threads were executing main finsihes executiong

	}
	// time.Sleep(time.Second*1);
	wg.Wait()  


	// so there are 3 things to notice in the case of wait groups 
	// add done and then wait 

	// now we gonna see the wait groups 
	// we shouldn't use the timeSleep here as wee don;t know how much time will it take
    // for that we gonna use wait groups 

	// to synchronise the goroutines 


}