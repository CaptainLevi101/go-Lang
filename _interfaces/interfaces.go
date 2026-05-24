// package main

// import "fmt"

// type payment struct {
// }

// func (p payment) makePayment(amount float32) {
// 	// here i need payment gateway
// 	razorPayPaymentGw:=razorpay{};
// 	// what if we wan tmore payment gateways
// 	// so there are multiple reirements wanna use stripr
// 	// instead of razor pay
// 	// we can use another struct of typre stripr andd do the same shitt again

// 	razorPayPaymentGw.pay(amount);
// }

// type razorpay struct {
// }

// func (r razorpay) pay(amount float32) {
// 	// logic to make payment
// 	fmt.Println("Making payment using razorPay ",amount);

// }

// func main() {

// 	// firrst of all make instance of payment struct
// 	newPayment:=payment{};
// 	newPayment.makePayment(100);

// }

package main

import "fmt"



type paymentGateWay interface{
  makePayment(amount float32)
};

type payment struct{
	// now the payment will be if different types 
	gateway paymentGateWay
}

func (p payment) pay(amount float32){
	fmt.Println("paid amount ",amount);
	// now i wanna procees the payment through gateways 
}


type razorPay struct{

}
func (r razorPay) makePayment(amount float32){
	fmt.Println("Amount paid using razorPay",amount);
}


type stripe struct{

}
func (s stripe) makePayment(amount float32){
    fmt.Println("Amount paid using stripe ",amount);
}

func main(){
	//paymentUsingStripe:=stripe.makePayment(34);
	stripePayment:=stripe{};
	razorPayment:=razorPay{};
	// r:=payment{};
	r:=payment{
		gateway: stripePayment,
	};
	p:=payment{
		gateway: razorPayment,
	}
	p.gateway.makePayment(38);
	r.gateway.makePayment(64);
}

// now see what can be done in her e