package main

import (
	"fmt"
	"math/rand"
)

var BOARD = [100]uint8{
	0, 0, 0, 14, 0, 0, 0, 0, 31, 0,
	0, 0, 0, 0, 0, 0, 7, 0, 0, 0,
	42, 0, 0, 0, 0, 0, 0, 84, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	67, 0, 0, 34, 0, 0, 0, 0, 0, 0,
	0, 19, 0, 60, 0, 0, 0, 0, 0, 0,
	0, 91, 0, 0, 0, 0, 0, 0, 0, 99,
	0, 0, 0, 0, 0, 0, 36, 0, 0, 0,
	0, 0, 73, 0, 75, 0, 0, 79, 0, 0,
}

func roll(space uint8) uint8 {
	roll := uint8(rand.Intn(6) + 1)
	space += roll

	fmt.Printf("You rolled a %v.\n", roll)
	fmt.Printf("You are now on space %v.\n", space)

	return space
}

func check_over_goal(space uint8) uint8 {
	if space > 100 {
		spaces_over := space - 100
		space -= spaces_over * 2

		fmt.Printf("You have rolled %v spaces over 100.\n", spaces_over)
		fmt.Printf("You are now on space %v.\n", space)
	}

	return space
}

func check_movement(space uint8) uint8 {
	space_index := BOARD[space-1]

	var thing string

	if space_index != 0 {
		if space_index > space {
			thing = "ladder"
		} else {
			thing = "snake"
		}

		space = space_index

		fmt.Printf("You landed on a %v!\n", thing)
		fmt.Printf("You are now on space %v\n", space)
	}

	return space
}

func main() {
	// declares variables
	var space, turn uint8 = 1, 0
	play := true

	// prints name of the game to user
	fmt.Println("Snakes & Ladders")

	// code loop
	for play {
		// takes input for exit code
		var input string
		fmt.Scanln(&input)
		if len(input) == 1 && input == "q" {
			return
		}

		turn++

		space = roll(space)

		play = space != 100

		space = check_over_goal(space)
		space = check_movement(space)
	}

	fmt.Printf("You win!!!\nIt took %v turns.\n", turn)
}
