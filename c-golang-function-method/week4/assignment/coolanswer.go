package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct {
	food       string
	locomotion string
	noise      string
}

func (c Cow) Eat() {
	fmt.Println(c.food)
}

func (c Cow) Move() {
	fmt.Println(c.locomotion)
}

func (c Cow) Speak() {
	fmt.Println(c.noise)
}

type Bird struct {
	food       string
	locomotion string
	noise      string
}

func (b Bird) Eat() {
	fmt.Println(b.food)
}

func (b Bird) Move() {
	fmt.Println(b.locomotion)
}

func (b Bird) Speak() {
	fmt.Println(b.noise)
}

type Snake struct {
	food       string
	locomotion string
	noise      string
}

func (s Snake) Eat() {
	fmt.Println(s.food)
}

func (s Snake) Move() {
	fmt.Println(s.locomotion)
}

func (s Snake) Speak() {
	fmt.Println(s.noise)
}

func CreateNewAnimal(name string, atype string, store *map[string]Animal) (bool, error) {
	var err error
	switch atype {
	case "cow":
		(*store)[name] = Cow{
			food:       "grass",
			locomotion: "walk",
			noise:      "moo",
		}
	case "bird":
		(*store)[name] = Bird{
			food:       "worms",
			locomotion: "fly",
			noise:      "peep",
		}
	case "snake":
		(*store)[name] = Snake{
			food:       "mice",
			locomotion: "slither",
			noise:      "hsss",
		}
	default:
		err = fmt.Errorf("%s - Unsupported animal type: [cow|bird|snake] are only allowed", atype)
	}
	if err != nil {
		return false, err
	}
	return true, nil
}

func validate(input string) (bool, error) {
	words := strings.Split(input, " ")
	if len(words) < 3 {
		return false, fmt.Errorf("Command should be of 3 words separated by <SPACE> %s", input)
	}
	return true, nil
}

func main() {
	var input string
	store := map[string]Animal{}
	for {
		fmt.Print(">")
		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			input = scanner.Text()
		} else {
			panic("Unexpected error while capturing user input")
		}
		if status, err := validate(input); status {
			work := strings.Split(input, " ")
			command, name, ctype := work[0], work[1], work[2]
			switch command {
			case "newanimal":
				if status, err := CreateNewAnimal(name, ctype, &store); status {
					fmt.Println("Created it!")
				} else {
					fmt.Println(err.Error())
				}
			case "query":
				if animal, exists := store[name]; exists {
					switch ctype {
					case "eat":
						animal.Eat()
					case "move":
						animal.Move()
					case "speak":
						animal.Speak()
					default:
						fmt.Println("Invalid request type - ", ctype, " [eat|move|speak] are only allowed.")
					}
				}
			}

		} else {
			fmt.Println(err.Error())
		}
	}
}
