package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"sync"
	tictactoe "tictactoe/pkg"
)

var CurrentGame tictactoe.Game
var InitGame tictactoe.Game
var Mutex *sync.Mutex

func main() {
	Mutex = &sync.Mutex{}
	InitGame = tictactoe.Game{
		CurrentPlayer: tictactoe.PlayerX,
		GameState:     tictactoe.InitGame,
	}
	CurrentGame = InitGame
	handleRequests()
}

func check(w http.ResponseWriter, err error) {
	if err != nil {
		_, errPrint := fmt.Fprintf(w, "error: %v", err)
		if errPrint != nil { panic("Could not print") }
	}
}

func homepage(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		postGame(r)
	}
	tmpl, err := template.ParseFiles("./templates/index.html")
	check(w, err)
	check(w, tmpl.Execute(w, CurrentGame))
}

func postGame(r *http.Request) {
	Mutex.Lock()

	val, err := strconv.Atoi(r.FormValue("index"))
	if err != nil { return }
	if val >= 0 && val < 9 {
		newState, err := CurrentGame.GameState.Play(val)
		if err != nil {
			Mutex.Unlock()
			return
		}
		CurrentGame = tictactoe.Game{
			CurrentPlayer: newState.CurrentPlayer(),
			GameState:     newState,
		}
		if winner, over := newState.IsOver(); over {
			CurrentGame.Winner = winner
		}
	} else if val == -1 {
		CurrentGame = InitGame
	} else if val == -2 {
		newState, err := CurrentGame.GameState.NextBestMove()
		if err != nil {
			Mutex.Unlock()
			return
		}
		CurrentGame = tictactoe.Game{
			CurrentPlayer: newState.CurrentPlayer(),
			GameState:     newState,
		}
		if winner, over := newState.IsOver(); over {
			CurrentGame.Winner = winner
		}
	}

	fmt.Println(CurrentGame.GameState)

	Mutex.Unlock()
}

func handleRequests() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", homepage)
	log.Fatal(http.ListenAndServe(":8000", nil))
}