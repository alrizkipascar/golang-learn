package main

import (
	"fmt"
	"os"
	"strings"
)

type Animals struct {
	food       string
	locomotion string
	speak      string
}

func (animals Animals) Eat() {
	fmt.Println(animals.food)
}

func (animals Animals) Move() {
	fmt.Println(animals.locomotion)
}

func (animals Animals) Speak() {
	fmt.Println(animals.speak)
}

func getRequest() (string, string) {
	var animalChoice, infoChoice string
	for {
		fmt.Print(">")
		if _, err := fmt.Scanln(&animalChoice, &infoChoice); err != nil {
			fmt.Println("Error:", err)
		} else {
			break
		}
		if strings.ToLower(animalChoice) == "exit" {
			fmt.Println("Good Bye")
			os.Exit(1)
		}
	}
	return animalChoice, infoChoice
}

func main() {
	fmt.Println("Please type '[cow|bird|snake] [eat|move|speak]' after the prompt \">\" to discover animals. Press CTRL+C or type exit to exit .")
	cow := Animals{"grass", "walk", "moo"}
	bird := Animals{"worms", "fly", "peep"}
	snake := Animals{"mice", "slither", "hsss"}

	for {
		animalChoice, infoChoice := getRequest()

		var animal Animals
		if animalChoice == "cow" {
			animal = cow
		} else if animalChoice == "bird" {
			animal = bird
		} else if animalChoice == "snake" {
			animal = snake
		}

		if infoChoice == "eat" {
			animal.Eat()
		} else if infoChoice == "move" {
			animal.Move()
		} else if infoChoice == "speak" {
			animal.Speak()
		}
	}
}
