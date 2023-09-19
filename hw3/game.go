package homework

import ( 
	"fmt"
	"os"
	"math/rand"
	"bufio"
	"strconv"
	"strings"
)

type Character struct {
	Name string
	Sleeping bool
	Stuff []string 
}

// Check if main character is sleeping
func (character *Character) IsSleeping() bool {
	return character.Sleeping
}

// Lets wake up main character
func (character *Character) WakeUP() {
	character.Sleeping = false
}

func (character *Character) Sleep() {
	character.Sleeping = true
}

var deadBodyAnimalAnswer string
var caveEntryAnswer string
var tentAnswer string
var min, max, attempts int
var secretNumber int


func Game() {
	character := Character {"Steven", true, []string{"backpack", "knife", "flashlight", "match"}}
	if character.Sleeping == true {
		fmt.Printf("%s is sleeping right now, Lets wake up him\n", character.Name)
		goto WakeUp 
	} else {
		goto CaveEntry
	}

WakeUp:
	character.WakeUP()
	if character.Sleeping == false {
		fmt.Printf("%s is not sleeping\n", character.Name)
		goto CaveEntry 
	}

CaveEntry:
	fmt.Printf("\n")
	fmt.Println("hello, i am next to Cave. Its very dark in a cave but i have a lighter and i  can check what is inside \n")
	fmt.Print("Should i go inside to a cave or use path into a forest? yes -> go to a cave, no -> go to a forest\n")
	fmt.Scan(&caveEntryAnswer)
	switch caveEntryAnswer {
		case "yes":
			fmt.Print("You answered yes, lets go to a cave and check what is going on there\n")
			goto Cave
		case "no":
			fmt.Print("You answered no, lets use that path and go to the Forest\n")
			goto Forest
		default:
			fmt.Print("You didn't input right answer, you could use yes or no (lowercase)")
			goto CaveEntry
	}

Cave:
	fmt.Print("I am in a cave, lets turn on our lighter. There is a ladder inside from the cave\n")
	goto Exit 

Forest:
	fmt.Print("I am in a forest, I see dead body of the animal, What should i do? skip or check ?  \n")
	fmt.Print("possible answer : skip and check \n")
	fmt.Scan(&deadBodyAnimalAnswer)
	switch deadBodyAnimalAnswer {
		case "skip":
			fmt.Print("Ok, your answer is skip and you will go ahead\n ")
			goto Camp
		case "check":
			fmt.Print("Ok, your answer is check, animal was poisened and you are dead too\n ")
			goto LoseGame
		default:
			fmt.Print("Your answer is not correct, you can use skip and check (lowercase)\n ")
			goto Forest 
	}
Camp:
	fmt.Print("I am in a empty camp and i see a tent. I am tired and would like to have a rest \n ")
	fmt.Print("Should i use a tent or should i go. Possible answer: go or rest ?\n")
	fmt.Scan(&tentAnswer)
	switch tentAnswer {
		case "rest":
			fmt.Print("Lets use that tent and have a rest\n")
			goto Tent
		case "go":
			fmt.Print("It saved your life, there is a dangerous bug in a tent\n ")
			goto Exit
		default:
			fmt.Print("Your answer is not correct, you can use rest or go (all lowercase), lets try again\n")
			goto Camp
	}

Tent:
	fmt.Print("I am in a tent and see a safe, let me check what is inside")
	// i would like to make Steven life easier and correct nubmer between 10 and 20
	fmt.Println("Guess a number between 10 and 20")
	fmt.Println("Please input your guess")
	min = 10
	max = 20
	attempts = 0
	secretNumber = rand.Intn(max-min) + min
	
	for {
		attempts++
		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("An error reading input. Please try again\n", err)
			continue
		}

		input = strings.TrimSuffix(input, "\n")

		guess, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid input. Please enter an integer value\n")
			continue
		}

		fmt.Println("Your guess is", guess)

		if guess != secretNumber {
			fmt.Println("Your guess not right. Try again\n")
		} else {
			fmt.Println("Correct, safe is open\n")
			break
		}
	}
	fmt.Print("There is a poisened bug inside in a safe\n")
	goto LoseGame

LoseGame:
	fmt.Print("GAME OVER\n")
	os.Exit(0)
Exit: 
	fmt.Print("you are saved, you found how to espace\n")
}
