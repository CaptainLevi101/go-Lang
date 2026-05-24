package main

import (
	"fmt"
	"slices"
)

func add(a int,b int) int{
	return a+b
  }

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
// now let's study about slice 
//slice is like a vector in the case of go, dynamic arrays 
  // uniintialised slice is nil 
  var nums=make([]int,2);
  var nums2=make([]int,len(nums))
 
  nums=append(nums,2);
  nums2=append(nums2,2)
  fmt.Println(nums,nums2);

  // so for adding an element in the vector for go we use append function 

// slice operator 
// slice ke andr ke elements ko return krke dega 

fmt.Println(nums[0:2]);
// lets make a slice 
// var nums1=[4]int{1,2,3,4};
fmt.Println(slices.Contains(nums,2));

// we can also maek 2d slices in here 
var nums3=[][]int{{1,2},{3,4}}
fmt.Println(nums3[0]);


// now we will be using maps in go 
   m:=map[string]string{
	"Name":"Ashish",
	"Age":"24",
	"Work":"Currently focusing on goals",
   }

   fmt.Println(m)


   // so use ok, to know whether the key exist in the map or not 
   // if the key doesn't exist then we get something diff here 


   // so maps are reference types in here 
   // if i create a map 
   // and
  // do something like b=a then change b then the value of a will also be changed 


  // now what do we ahve is ran ge in the go 
  // range is for iterating over ds 
  // lets we hace a slice 
 
   // now let's do the sum over this and use range 
   a:=[]int{1,2,3,4};
   sum:=0;
   for _,ai:=range a{
	sum+=ai
	 fmt.Print(ai)
   }
   fmt.Print(sum);

   // this is how we iterate over the slice and get the sum 
   // so first value is more or less like the index we will be loooping obver 

   m1:=map[string]string{
	"fname":"Ashish",
	"lName":"Parashar",
   }

   for k,v:=range m1{
	// k is the key and v is the value 
	fmt.Print(k,v);
   }


   // now range can also be used in over stirng 
   var name string="Ashish";

   for _,c:=range name{
	 fmt.Print(string(c))
	 // here c is the unicode given so it is more or less like a unicode 
	 // first parameter is the starting byte of rune
   }

   // so here is our range we can use it for map, slice and string as well 

  // now we will be covering the functions in go 


  fmt.Print(add(5,4))
  println(sumo(3,4,5,6,7,8,9,9,0,546,32,2,34))
}


// now we gonna see is vaiadic functions 
// the functions that can take any number of parameters are known as variadic functinons
// if we want to recieve the type of any  then we suee interface 

func sumo(nums...int)int{
	total:=0;
	for _,num:=range nums{
		total+=num;
	}
	return total;
}

