package auth

import "fmt"

func LoginWithCredentials(username string, password string) {
	fmt.Println("Login user using username", username,password);
}

func CreateSession(username string){
	fmt.Println("Creating session for user", username);
}

// to use this package in our main file we need to import it and then we can use the function in it
// we also need to use the term scope 

//we use capital letter to make the function public and small letter to make it private
//dang this is really a cool information