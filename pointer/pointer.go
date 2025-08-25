package main

import (
	"fmt"
)

func main() {
	userID := 42
	fmt.Println(&userID) //will print the address of userID
	anotherUserID := &userID
	fmt.Println(anotherUserID) //will print the address of userID
	//now what happens if we change the value of userID
	userID = 100
	fmt.Println(userID)         //the value of userID is changed as in overridden
	fmt.Println(&userID)        //the address of userID is still the same
	fmt.Println(anotherUserID)  //the address of anotherUserID is still the same
	fmt.Println(*anotherUserID) //the value of the box that anotherUserID is pointing to is changed
	// and this is how we also access the value of the box through pointer

	//fmt.Println(*userID)        //will give error as userID is not a pointer but an integer so we can't dereference it

	//now that we are talking about it, the * operator is also used to declare a type pointer
	var age *int      //age is a pointer to an integer and it is not initialized
	fmt.Println(age)  //will print nil as age is a pointer to an integer and it is not initialized
	age = &userID     //if we want to assign the value of userID to age, we need to use the & operator
	fmt.Println(*age) //will print the value of userID //dereferencing the pointer to get the value

	//you can also pass pointers to the functions
	update(age, 42)
	fmt.Println(*age) //will print the value of userID which is 42 as the value of userID is updated to 42
	//updateButNotPassingPointer(userID, 100)
	//fmt.Println(userID) //will print the value of userID which is 100 as the value of userID is updated to 100
	updateButNotPassingPointer(*age, 69)
	fmt.Println(*age) //will print the value of userID which is 42 as the value of userID is updated to 42
}
func update(val *int, to int) {
	*val = to
}
func updateButNotPassingPointer(val int, to int) { //this is a copy of the value of val
	val = to //this is a copy of the value of val
	fmt.Println(val) //will print the value of val which is 69 as the value of val is updated to 69
	//but the value of the original val is not updated as we are passing the value of val not the pointer to the function
}
