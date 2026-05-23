package main

import "fmt"

func main(){
	fmt.Println("hello world")  // used for fomatting 


// what si 
// so we can first build and then run 
whatsmytype:=func(i interface{}){
	switch t:=i.(type){
	case int:
		fmt.Println("integer baby")
	case string:
		fmt.Println("string baby")
	case bool:
		fmt.Println("Boolean baby")
	default:
        fmt.Println("other baby",t)
	}
}
whatsmytype(44)
}
