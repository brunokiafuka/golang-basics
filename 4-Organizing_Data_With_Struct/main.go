package main

import "fmt"

type contactInfo struct {
	email string
	zip   string
}

type person struct {
	firstName string
	lastName  string
	age       int
	contact   contactInfo
}

func main() {
	/*var bruno person

	bruno.firstName = "Bruno"
	bruno.lastName = "Kiafuka"
	bruno.age = 23
	bruno.contact.email = "bruno@email.com"
	bruno.contact.zip = "1021" */

	bruno := person{
		firstName: "Bruno",
		lastName:  "Kiafuka",
		age:       23,
		contact: contactInfo{
			email: "bruno@email.com",
			zip:   "1021",
		},
	}
	// bruno.updateName("Sebas")

	//Pointer
	//brunoPointer := &bruno
	bruno.updateName("Sebastian")
	bruno.print()
}

func (p *person) updateName(newFirstName string) {
	p.firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
