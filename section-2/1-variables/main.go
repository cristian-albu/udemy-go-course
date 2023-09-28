package main

import "fmt"


func main(){
	fmt.Println("Hello world")

	var whatToSay string
	var i int

	whatToSay = "Goodbye, cruel world"

	fmt.Println(whatToSay)

	i = 15

	fmt.Println("i is set to", i)

	whatWasSaid, whatWasSaid2 := saySomething()

	fmt.Println(whatWasSaid, whatWasSaid2)
}

func saySomething() (string, string) {
	return "something", "otherThings"
}