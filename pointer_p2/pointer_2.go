package main

import (
	"fmt"
)

type User struct {
	ID   int
	Name string
}

func (u *User) updateName(name string, id int) *User {
	u.Name = name
	u.ID = id
	return u
}

func main() {
	user := User{ID: 1, Name: "John"}
	user.updateName("Jane", 2)
	//updateName(user, "Jane", 2)
	fmt.Println(user.Name, user.ID)
}
