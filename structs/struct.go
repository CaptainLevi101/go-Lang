package main

import (
	"fmt"
	"time"
)


type customer struct{
	fname string
	lname  string
	id     string
}

type order struct{
  orderDate time.Time  
  Address  string 
  payment  bool
  customer
}

func (o* order) orderFunc(){
	// her the order should be given by refernce if we wants to modify any value in it 
	// its goal is to change the payment status of the order struct to true
	o.payment=true;

}


// now see i have created this struct and i wanna 


func main() {
	myOrder:=order{
		orderDate: time.Now(),
		Address: "21434",
		payment: false,
	}
	// }{
	// 	// i can also give the intial values in here 
	// }
    fmt.Println(myOrder.orderDate.Day())
	// we can also fill it with less number of fields as reired 
    myOrder.orderFunc();
	fmt.Println(myOrder.payment);
	myOrder.customer.fname="Ashish";
	fmt.Println(myOrder.customer.fname);
    

	// we can also use the getters and setters functions in here 
	// depend upon the approach and the idea we gonna have 

	fmt.Println("Hii");
}