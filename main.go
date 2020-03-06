package main

import (
	"fmt"
	"reflect"
)

type Gender int

const (
	Male Gender = iota
	Female
)

type PhoneNumber struct {
	Mobile string
	Home   string
}

type Addresses struct {
	Work string
	Home string
}

type User struct {
	Name      string
	Sexuality Gender
	Phone     PhoneNumber
	Addr      Addresses
}

func main() {
	u := User{
		Name:      "me",
		Sexuality: Male,
		Phone: PhoneNumber{
			Mobile: "12345",
			Home:   "54321",
		},
		Addr: Addresses{
			Work: "this is work address",
			Home: "this is home address",
		},
	}

	fmt.Println("type of u: ", reflect.TypeOf(u))
	fmt.Println("type of u.Name: ", reflect.TypeOf(u.Name))
	fmt.Println("type of u.Phone: ", reflect.TypeOf(u.Phone))
}
