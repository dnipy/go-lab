package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func changeName(user User,name string) {
	fmt.Println("--------------------------")
	fmt.Println("[without pointer]")
	fmt.Println("before modify",user)
	user.Name = name
	fmt.Println("after modify",user)
	fmt.Println("--------------------------")
}


func changeNameWithPointer(user *User,name string) {
	fmt.Println("--------------------------")
	fmt.Println("[with pointer]")
	fmt.Println("before modify",user)
	user.Name = name
	fmt.Println("after modify",user)
	fmt.Println("--------------------------")
}

func main() {
	user := User{
		Name: "Danial",
		Age : 22,
	}

	fmt.Println("start app")
	fmt.Println("init user",user)

	changeName(user,"Alex")
	fmt.Println("changeName :",user)

	changeNameWithPointer(&user,"Alex")
	fmt.Println("changeNameWithPointer :",user)

	fmt.Println("--------------------------")


	pointer := &user

	fmt.Println("user:", user)
	fmt.Println("pointer:", pointer)
	fmt.Println("value behind pointer:", *pointer)
	fmt.Println("pointer name:", pointer.Name)
}