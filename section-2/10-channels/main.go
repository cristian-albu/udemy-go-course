package main

import (
	"log"

	"github.com/cristian-albu/channels-project/helpers"
)

const numpool = 1000

func CalculateValue(intChan chan int) {
	randomNumber := helpers.RandomNumber(numpool)
	intChan <- randomNumber
}

func main() {
	intChan := make(chan int)

	defer close(intChan)

	go CalculateValue(intChan)

	num := <-intChan

	log.Println(num)

}
