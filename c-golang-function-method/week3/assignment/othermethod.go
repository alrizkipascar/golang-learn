package main

import (
	"fmt"
	"strings"
)

// Animal is a struct to hold information about each animal.
type Animal struct {
	food       string
	locomotion string
	speak      string
}

// Eat prints the animal's food.
func (a Animal) Eat() {
	fmt.Println(a.food)
}

// Move prints the animal's locomotion.
func (a Animal) Move() {
	fmt.Println(a.locomotion)
}

// Speak prints the animal's spoken sound.
func (a Animal) Speak() {
	fmt.Println(a.speak)
}

func main() {
	// Define predefined animals
	animals := map[string]Animal{
		"cow":   {food: "grass", locomotion: "walk", speak: "moo"},
		"bird":  {food: "worms", locomotion: "fly", speak: "peep"},
		"snake": {food: "mice", locomotion: "slither", speak: "hsss"},
	}

	// Main loop for processing user requests
	for {
		// Prompt for user input
		fmt.Print(">")
		var animalName, action string
		fmt.Scan(&animalName, &action)

		// Convert user input to lowercase for case-insensitivity
		animalName = strings.ToLower(animalName)
		action = strings.ToLower(action)

		// Check if the requested animal exists
		if animal, ok := animals[animalName]; ok {
			// Process the user request based on the action
			switch action {
			case "eat":
				animal.Eat()
			case "move":
				animal.Move()
			case "speak":
				animal.Speak()
			default:
				fmt.Println("Invalid action. Please enter 'eat', 'move', or 'speak'.")
			}
		} else {
			fmt.Println("Invalid animal. Please enter 'cow', 'bird', or 'snake'.")
		}
	}
}
