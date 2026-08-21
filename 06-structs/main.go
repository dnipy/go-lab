package main

import "fmt"

type User struct {
	id       int
	name     string
	age      int
	isActive bool
}

func (u *User) getName() string {
	return  u.name
}

func (u *User) getAge() int {
	return  u.age
}

func (u *User) activationStatus() bool {
	return  u.isActive
}


func (u *User) modifyActivate(active bool) bool {
	u.isActive = active
	return  active
}

func main() {
	zero_user := User{}
	user := User{
		id:       1,
		name:     "danial",
		age:      22,
		isActive: true,
	}
	user.getName()
	
	fmt.Println("zero user :",zero_user)
	fmt.Println("user :",user)
	fmt.Println(user.getName())
	fmt.Println(user.getAge())
	fmt.Println(user.activationStatus())
	user.modifyActivate(false)
	fmt.Println(user.activationStatus())
	user.modifyActivate(true)
	fmt.Println(user.activationStatus())

}