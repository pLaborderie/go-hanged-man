package api

import (
	"fmt"
	"hangedman/model"
	"log"
	"net/http"
)

var game = model.Game{}
var rooms []model.Room

func initGame() {
	log.Println("Initializing game")
	game.PickWord()
	game.GuessedLetters = []string{}
	log.Println("The word to guess is: ", game.Word)
}

func writeWinMessage(w http.ResponseWriter) {
	_, err := fmt.Fprint(w, "Game won! The word was: ", game.Word)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func writeLossMessage(w http.ResponseWriter) {
	_, err := fmt.Fprint(w, "Game lost! The word was: ", game.Word)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func startGameHandler(w http.ResponseWriter, r *http.Request) {
	initGame()
	w.WriteHeader(http.StatusOK)
}

func hangedManHandler(w http.ResponseWriter, r *http.Request) {
	if game.IsGameWon() {
		writeWinMessage(w)
		return
	}
	if game.IsGameLost() {
		writeLossMessage(w)
		return
	}

	art := game.GetHangedMan()
	livesLeft := game.GetLivesLeft()
	word := game.GetWord()
	wrongGuesses := game.GetWrongGuesses()
	displayedText := fmt.Sprintf("%s\n%d\n%s\n%s", word, livesLeft, wrongGuesses, art)
	_, err := fmt.Fprint(w, displayedText)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func guessLetterHandler(w http.ResponseWriter, r *http.Request) {
	guess := r.PathValue("guess")
	if guess == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if len(guess) > 1 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if game.IsGuessed(guess) {
		w.WriteHeader(http.StatusAlreadyReported)
		return
	}
	if game.IsGameOver() {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	game.AddGuess(guess)
}

func registerRoutes() {
	http.HandleFunc("GET /", hangedManHandler)
	http.HandleFunc("POST /start", startGameHandler)
	http.HandleFunc("POST /{guess}", guessLetterHandler)
}

func StartWebServer() {
	initGame()
	registerRoutes()
	log.Fatal(http.ListenAndServe(":8080", nil))
}
