package main

import (
	"fmt"
	"sync"
)

type post struct {
	views int 
	mu sync.Mutex
}

func (p *post) inc(wg *sync.WaitGroup) {
	defer func(){
		p.mu.Unlock()
		wg.Done()
	}()
	p.mu.Lock()
	p.views += 1
}

func main() {

	// jb bhi multiple process same resource ko modify kren
	// then there won't be futhure atomicity

	myPost := post{views: 0}

	// myPost.inc()

	// myPost.inc()
	// myPost.inc()
	// myPost.inc()
	// fmt.Println("MyPost Views",myPost.views);


	// but in normal scene, the reest may or may not be in order 
	// what if the reuest was concurrent in that case we gonna see what can be done in here 
	var wg sync.WaitGroup 

    for i:=0;i<100;i++{
		wg.Add(1)
		go myPost.inc(&wg)
	}

	wg.Wait()

	fmt.Println(myPost.views)  

	// mutext can b econsidered as a lock, as if it is being used by one of the resource it is lockd by it 
    // one problem in mutex is it makes the other process weight 
	// lock only the line where modification is done
	// don't lock the entire code 
	// so it is a good process to lock only at the modification 


}