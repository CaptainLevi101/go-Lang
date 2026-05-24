package main

import "fmt"

// now we gonna see what are enums
func changeOrderStatus(status OrderStatus){
	fmt.Println("changind ordere status", status);
}


// now there can be many status like
// confirmed, recieved etc 
// what if in futire we want to vchange thi sconfirmed status to something else 
// what if there are multilple status 

// for enum we use custom data type 
type OrderStatus int 

const(
	recieved OrderStatus=iota   // predeclared identifer that goes on increasing 
	confirmed
	prepared
	delivered
)
func main(){

	changeOrderStatus(confirmed);


}