package main

import "log"

var s string = "seven"

func main() {

	var s2 = "six"

	s := "eight"
	log.Println("s is", s)
	log.Println("s2 is", s2)

	saySomething("xxx")
}

func saySomething(s3 string) (string, string) {
	log.Println("s from the function", s)
	return s3, "World"
}


func Whatever() {
	log.Println("Public is with capital letter")
}

func something() {
	log.Println("This is private to main")
}