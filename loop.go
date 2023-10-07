package main

import "log"

func main() {
	for i := 0; i <= 3; i++ {
		log.Println(i)
	}

	animals := []string{"dog", "fish", "horse", "cow", "cat"}
	for _, animal := range animals {
		log.Println(animal)
	}

	newAnimals := make(map[string]string)
	newAnimals["dogs"] = "a"
	newAnimals["cat"] = "b"
	for animalType, animal := range newAnimals {
		log.Println(animalType, animal)
	}

	var firstLine = "Once upon a Time"
	for i, letter := range firstLine {
		log.Println(i, ":", letter)
	}

	type User struct {
		FirstName string
		LastName  string
		Email     string
		Age       int
	}
	var users []User
	users = append(users, User{"John", "Smith", "js@gmail.com", 30})
	users = append(users, User{"Lebron", "James", "lb@gmail.com", 38})
	users = append(users, User{"James", "Harden", "jh@gmail.com", 45})
	for _, l := range users {
		log.Println(l.FirstName, l.LastName, l.Email, l.Age)
	}

}
