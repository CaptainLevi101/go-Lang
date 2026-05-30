package main

import "github.com/CaptainLevi101/go-Lang/_packages/auth"

func main(){
	// here we gonna learn what are packages 
	// package is a collection of files in the same directory that are compiled together
	// we can import the package in our main file and use the functions in it 
	// we can also create our own package and use it in our main file 
	// we can also use the standard library packages in go 
	// we can also use third party packages in go 
	// we can also use the go get command to install the third party packages in go 
	// we can also use the go mod command to manage the dependencies in go 
	// we can also use the go build command to build the go files and create an executable file 
	// we can also use the go run command to run the go files without building them

	//go mod 
	//go get 
	//go build
	//go run 
	// go mod tidy to remove the unused dependencies in go and add the missing dependencies in go and also to update the go.mod file in go
	// let's see the authentication package in go
	// we can use the authentication package in go to authenticate the user 
	// we can also use the authentication package in go to generate the token for the user 
	// we can also use the authentication package in go to validate the token for the user 
	// we can also use the authentication package in go to refresh the token for the user 
	// we can also use the authentication package in go to revoke the token for the user


	auth.LoginWithCredentials("Ashish","password123");
	auth.CreateSession("Ashish");

}