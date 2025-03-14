package main

import (
	"fmt"
	"hangedman/api"
	"hangedman/model"
	"os"
	"slices"
)

func getUserInput() string {
	var userInput string
	fmt.Print("Enter a letter: ")
	fmt.Scanln(&userInput)
	return userInput
}

func cliGame() {
	game := model.Game{}
	game.PickWord()

	cliArgs := os.Args[1:]
	if slices.Contains(cliArgs, "DEBUG") {
		fmt.Println("[DEBUG] The word is:", game.Word)
	}

	fmt.Println("Welcome to Hanged Man!")
	fmt.Println("Good luck!")
	fmt.Println("You have 6 chances to guess the word.")

	for !game.IsGameOver() {
		game.ShowLivesLeft()
		game.ShowWord()
		game.ShowWrongGuesses()
		game.ShowHangedMan()

		userInput := getUserInput()
		if len(userInput) != 1 {
			fmt.Println("Please enter a single letter.")
			continue
		}

		if game.IsGuessed(userInput) {
			fmt.Println("You've already guessed that letter.")
			continue
		}

		game.AddGuess(userInput)
	}

	if game.IsGameWon() {
		fmt.Println("Congratulations! You won!")
		fmt.Println("The word is:", game.Word)
	}

	if game.IsGameLost() {
		game.ShowHangedMan()
		fmt.Println("You lost! The word was:", game.Word)
	}
}

func isApi() bool {
	cliArgs := os.Args[1:]
	return slices.Contains(cliArgs, "API")
}

func main() {
	game := model.Game{}
	game.PickWord()

	cliArgs := os.Args[1:]
	if slices.Contains(cliArgs, "DEBUG") {
		fmt.Println("[DEBUG] The word is:", game.Word)
	}

	if isApi() {
		api.StartWebServer()
	} else {
		cliGame()
	}
}
