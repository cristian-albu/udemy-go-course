package main

import (
	"log"
	"sort"
)


type User struct {
	FirstName string
	LastName string
}

func main() {

	// MAPS

	// var myOtherMap map[string]string

	myMap := make(map[string]string)

	myMap["dog"] = "Odin"
	myMap["cat"] = "Kira"

	log.Println(myMap["dog"])
	log.Println(myMap["cat"])

	myMap["dog"] = "Odinut"

	log.Println(myMap["dog"])

	myMap2 := make(map[string]int)

	myMap2["First"] = 1

	myMap2["Second"] = 2

	log.Println(myMap2["First"], myMap2["Second"])

	myUserMap := make(map[string]User)

	me := User {
		FirstName: "Cristian",
		LastName: "Albu",
	}

	myUserMap["me"] = me

	log.Println(myUserMap["me"].FirstName)


	// SLICES


	var mySlice []string

	mySlice = append(mySlice, "Cristi")
	mySlice = append(mySlice, "Andreea")
	mySlice = append(mySlice, "Odin")
	mySlice = append(mySlice, "Kira")

	log.Println(mySlice)

	var myIntsSlice []int

	myIntsSlice = append(myIntsSlice, 2)
	myIntsSlice = append(myIntsSlice, 1)
	myIntsSlice = append(myIntsSlice, 3)

	sort.Ints(myIntsSlice)

	log.Println(myIntsSlice)

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	log.Println(numbers)

	log.Println(numbers[0:2])
	
}