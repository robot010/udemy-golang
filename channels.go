package main

import (
	"/Users/bdeng/Library/CloudStorage/OneDrive-Personal/coursera/golang-udemy/go-overview/helpers.go"
	"log"
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
