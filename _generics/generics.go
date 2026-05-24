package main

import "fmt"

func main() {
    nums:=[]int{1,2,3,4}
	printSlice(nums)


	nums2:=[]string{"Ashish", "Parashar", "isBest"};
	printSlice(nums2)
}

// now let's see what is generic, we can print intgeer slics, what if we also have string slice 
// to solve the  repeteiveness of different functions we can juse generics 
//just like T in the typescript

// what if you opnly wanna allow integers and strign 
// T int | string    this thing is also fine and well 

func printSlice[T any](items []T) {
	for _,v:= range items {
		fmt.Println(v)
	}
}