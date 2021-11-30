package main

import (
	"encoding/json"
	"github.com/esmaeilmirzaee/random-time-sleeper/pkg/handlers"
	"log"
	"net/http"

	"github.com/esmaeilmirzaee/random-time-sleeper/helpers"
)

const upperBoundLimit = 1000
const portNumber = ":8080"

type Person struct {
	Id        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	HairColor string `json:"hair_color"`
	HasDog    bool   `json:"has_dog"`
}

func calculateValue(intChan chan int) {
	intChan <- helpers.GenerateRandomNumber(upperBoundLimit)
}


func main() {
	myJson := `
[
	{"id": 1, "first_name": "John", "last_name": "Doe", "hair_color": "Black", "has_dog": true},
	{"id": 2, "first_name": "Jane", "last_name": "Doe", "hair_color": "Black", "has_dog": false},
	{"id": 3, "first_name": "James", "last_name": "Jason", "hair_color": "Black", "has_dog": false}
]
`
	intChan := make(chan int)
	defer close(intChan)

	go calculateValue(intChan)

	log.Println(<-intChan)

	var unmarshalledPerson []Person

	if err := json.Unmarshal([]byte(myJson), &unmarshalledPerson); err != nil {
		log.Fatalln(err)
	}

	for _, item := range unmarshalledPerson {
		item.printPerson()
	}
	
	http.HandleFunc("/", handlers.HomePageHandler)
	_ = http.ListenAndServe(portNumber, nil)
}

func (person Person) printPerson() {
	log.Println("Person: ", person.Id, person.FirstName, person.LastName, person.HairColor, person.HasDog)
}
