package model

import (
	"hangedman/data"
	"math/rand"
	"slices"
	"strings"
)

const maxTries = 6
const letterPlaceholder = "_"

type Game struct {
	// The current state of the game
	Word           string
	GuessedLetters []string
}

func (game *Game) PickWord() {
	// Randomly pick a word from the wordList slice
	randomIndex := rand.Intn(len(data.WordList))
	game.Word = data.WordList[randomIndex]
}

func (game *Game) GetWord() string {
	// Display the word with underscores for letters not guessed yet
	displayedWord := ""
	for _, letter := range game.Word {
		if slices.Contains(game.GuessedLetters, string(letter)) {
			displayedWord += string(letter) + " "
		} else {
			displayedWord += letterPlaceholder + " "
		}
	}
	return displayedWord
}

func (game *Game) IsGameWon() bool {
	for _, letter := range game.Word {
		if !slices.Contains(game.GuessedLetters, string(letter)) {
			return false
		}
	}
	return true
}

func (game *Game) IsGameLost() bool {
	return len(game.GetWrongGuesses()) >= maxTries && !game.IsGameWon()
}

func (game *Game) IsGameOver() bool {
	return game.IsGameWon() || game.IsGameLost()
}

func (game *Game) GetWrongGuesses() []string {
	wrongGuesses := []string{}
	for _, letter := range game.GuessedLetters {
		if !strings.Contains(game.Word, letter) {
			wrongGuesses = append(wrongGuesses, letter)
		}
	}
	return wrongGuesses
}

func (game *Game) GetLivesLeft() int {
	return maxTries - len(game.GetWrongGuesses())
}

func (game *Game) GetHangedMan() string {
	return data.HangedManArt[game.GetLivesLeft()]
}

func (game *Game) IsGuessed(letter string) bool {
	return slices.Contains(game.GuessedLetters, letter)
}

func (game *Game) AddGuess(guess string) {
	game.GuessedLetters = append(game.GuessedLetters, guess)
}
