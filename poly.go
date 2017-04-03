package main

import (
	"fmt"
)

//polymorphism
type Address interface {
	GetAddress(string) string
}

type Company struct {
	Suite   int
	Number  int
	Street  string
	City    string
	State   string
	Zip     int
	address string //private
}

type Employee struct {
	Number  int
	Street  string
	City    string
	State   string
	Zip     int
	address string //private
}

func (c *Company) GetAddress(password string) (address string) {
	if password != "Lets Rock" {
		address = "Not Enough"
		return
	}

	if c.address != "" {
		address = c.address
	} else {
		address = fmt.Sprintf("%d %s, #%d\n%s, %s %d", c.Number, c.Street, c.Suite, c.City, c.State, c.Zip)
		c.address = address
	}
	return
}

func (e *Employee) GetAddress(password string) (address string) {
	if password != "Piece of Cake" {
		address = "Not Enough"
		return
	}

	if e.address != "" {
		address = e.address
	} else {
		address = fmt.Sprintf("%d %s\n%s, %s %d", e.Number, e.Street, e.City, e.State, e.Zip)
		e.address = address
	}
	return
}

func main() {
	things := make([]Address, 2)
	things[0] = &Employee{
		Number: 123,
		Street: "Foo Bazz Street",
		City:   "Malang",
		State:  "East Java",
		Zip:    666,
	}

	things[1] = &Company{
		Suite:  666,
		Number: 124,
		Street: "Boo Street",
		City:   "Malang",
		State:  "East Java",
		Zip:    666,
	}

	for _, item := range things {
		fmt.Println(item.GetAddress("Lets Rock"))
	}
}
