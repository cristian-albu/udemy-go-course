package main

import "log"

func main() {

	var isTrue bool
	isTrue = true

	if isTrue {
		log.Println("isTrue is", isTrue)
	} else {
		log.Println("isTrue is", isTrue)
	}

	cat := "cat"

	if cat == "cat" {
		log.Println("Cat is cat")
	} else {
		log.Println("Cat is not cat")
	}

	myNum := 100

	if myNum > 99 && !isTrue {
		log.Println("Yeah")
	} else if myNum < 100 || isTrue {
		log.Println("Kinda")
	} else {
		log.Println("nah")
	}

	myVar := "dog"

	switch myVar {
	case "cat": 
		log.Println("myVar is set to cat")
	case "dog": 
		log.Println("myVar is set to dog")
	case "bird": 
		log.Println("myVar is set to bird")
	case "horse": 
		log.Println("myVar is set to horse")
	default:
		log.Println("Sorry brah, probably chicken")
	}
}