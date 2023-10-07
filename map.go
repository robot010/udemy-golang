package main

import (
	"log"
	"sort"
)

type User struct {
	FirstName string
	LastName  string
}

func main() {
	myMap := make(map[string]string)

	myMap["dog"] = "Samson"
	myMap["other-dog"] = "Cassie"
	myMap["dog"] = "fido"

	log.Println(myMap["dog"])
	log.Println(myMap["other-dog"])

	mySecondMap := make(map[string]int)
	mySecondMap["first"] = 1
	mySecondMap["second"] = 2
	log.Println(mySecondMap["first"])

	myThirdMap := make(map[string]User)
	me := User{
		FirstName: "Trevor",
		LastName:  "Sawler",
	}
	myThirdMap["me"] = me
	log.Println(myThirdMap["me"].FirstName)

	var mySlice []string
	mySlice = append(mySlice, "abc")
	mySlice = append(mySlice, "ade")
	mySlice = append(mySlice, "bcd")

	sort.Strings(mySlice)

	log.Println(mySlice)

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	log.Println(numbers)
	log.Println(numbers[0:3])
}
