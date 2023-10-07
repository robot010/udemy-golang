package main

import (
	"log"

	"helpers"
)

const numPool = 10

func CalculateValue(intChan chan int) {
	RandomNumber := helpers.RandomNumber(numPool)
	intChan <- RandomNumber
}

func main() {
	intChan := make(chan int)
	defer close(intChan)
	go CalculateValue(intChan)

	num := <-intChan
	log.Println(num)
}
