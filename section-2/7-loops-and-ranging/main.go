package main

import "log"


func main() {
	for i := 0; i <= 5; i++ {
		log.Println(i)
	}


	animals := []string{"dog", "fish", "horse", "cow", "cat"}

	for i, animal := range animals {
		log.Println(i, animal)
	} 

	for _, animal := range animals {
		log.Println(animal)
	} 

	otherAnimals := make(map[string]string)

	otherAnimals["dog"] = "Odin"
	otherAnimals["cat"] = "Kira"

	for animalType, animal := range otherAnimals {
		log.Println(animalType, animal)
	}


	var firstLine string = "Once upn a midnight dreary"


	for i, l := range firstLine {
		log.Println(i, l)
	}

	type User struct {
		FirstName string
		LastName string
		Email string
		Age int
	}

	var users []User

	users = append(users, User{"John", "Smith", "johnsmith@email.com", 30})
	users = append(users, User{"Mary", "Henderson", "mary@email.com", 25})
	users = append(users, User{"Ally", "Johnson", "ally@email.com", 21})
	users = append(users, User{"Alex", "Tripoli", "alex@email.com", 33})


	for _, l := range users {
		log.Println(l.FirstName, l.LastName, l.Email, l.Age)
	}
}