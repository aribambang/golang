package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	// ari := person{"Ari", "Bambang"}

	// var ari person

	// ari.firstName = "ari"
	// ari.lastName = "bambang"

	ari := person{
		firstName: "Ari",
		lastName:  "Bambang",
		contact: contactInfo{
			email:   "aribk210@gmail.com",
			zipCode: 17411,
		},
	}

	ariPointer := &ari
	ariPointer.updateFirstName("Bambang")
	fmt.Println(&ari)

	// ari.print()
	// fmt.Println(ari)
}

// func (p person) updateFirstName(newFirstName string) {
// 	p.firstName = newFirstName
// }

func (pointerToPerson *person) updateFirstName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
