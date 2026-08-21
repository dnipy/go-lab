package main

import "fmt"

var package_scope_string string;
// it use zero assignment based on type
// or we can define like 
// var package_scope_string = ""; 

var package_scope_string_with_value = "hi"
// but we cant define like 
// package_scope_string_with_value := "hi"
// bcz := stle is for functions not global scope

const APP_PORT = 3000
const DEBUG = false
// constant and imutate
// but if it was var we could modify and mutate it 

func main () {
	name := "Daniel"
	age := 23
	active := true

	fmt.Println(name)
	fmt.Println("chagne name to Alex")
	name = "Alex"
	// we cant use := here 
	// cuz its just redeclare the same variable inside same scope again

	
	var username string
	var count int
	var enabled bool


	fmt.Println("[vars] :")
	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(active)


	fmt.Println("[zero assignment] :")
	fmt.Printf("%q\n", username)
	fmt.Println(count)
	fmt.Println(enabled)


	fmt.Println("[outside main()] :")
	fmt.Printf("%q\n",package_scope_string)
	fmt.Println(package_scope_string_with_value)
	fmt.Println(APP_PORT)
	fmt.Println(DEBUG)
}	