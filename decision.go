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
		log.Println("cat is cat")
	} else {
		log.Println("Cat is not cat")
	}

	myVar := "alien"

	switch myVar {
	case "cat":
		log.Println("myVar is set to cat")
	case "dog":
		log.Println("myVar is set to dog")
	case "bird":
		log.Println("myVar is set to bird")
	default:
		log.Println("myVar is set to", myVar)
	}
}
