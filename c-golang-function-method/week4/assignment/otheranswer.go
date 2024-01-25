package main

import (
	"fmt"
)

// Animal interface defines methods for an animal
type Animal interface {
	Eat()
	Move()
	Speak()
}

// Cow type satisfies the Animal interface
type Cow struct{}

// Eat prints the food of a cow
func (c Cow) Eat() {
	fmt.Println("grass")
}

// Move prints the locomotion method of a cow
func (c Cow) Move() {
	fmt.Println("walk")
}

// Speak prints the spoken sound of a cow
func (c Cow) Speak() {
	fmt.Println("moo")
}

// Bird type satisfies the Animal interface
type Bird struct{}

// Eat prints the food of a bird
func (b Bird) Eat() {
	fmt.Println("worms")
}

// Move prints the locomotion method of a bird
func (b Bird) Move() {
	fmt.Println("fly")
}

// Speak prints the spoken sound of a bird
func (b Bird) Speak() {
	fmt.Println("peep")
}

// Snake type satisfies the Animal interface
type Snake struct{}

// Eat prints the food of a snake
func (s Snake) Eat() {
	fmt.Println("mice")
}

// Move prints the locomotion method of a snake
func (s Snake) Move() {
	fmt.Println("slither")
}

// Speak prints the spoken sound of a snake
func (s Snake) Speak() {
	fmt.Println("hsss")
}

func main() {
	animals := make(map[string]Animal)

	for {
		var command, name, action string
		fmt.Print("> ")
		fmt.Scan(&command, &name, &action)

		switch command {
		case "newanimal":
			createAnimal(name, action, animals)
		case "query":
			queryAnimal(name, action, animals)
		default:
			fmt.Println("Invalid command. Please use 'newanimal' or 'query'.")
		}
	}
}

func createAnimal(name, animalType string, animals map[string]Animal) {
	var newAnimal Animal

	switch animalType {
	case "cow":
		newAnimal = Cow{}
	case "bird":
		newAnimal = Bird{}
	case "snake":
		newAnimal = Snake{}
	default:
		fmt.Println("Invalid animal type. Please choose 'cow', 'bird', or 'snake'.")
		return
	}

	animals[name] = newAnimal
	fmt.Println("Created it!")
}

func queryAnimal(name, action string, animals map[string]Animal) {
	animal, exists := animals[name]

	if !exists {
		fmt.Println("Animal not found.")
		return
	}

	switch action {
	case "eat":
		animal.Eat()
	case "move":
		animal.Move()
	case "speak":
		animal.Speak()
	default:
		fmt.Println("Invalid action. Please use 'eat', 'move', or 'speak'.")
	}
}
