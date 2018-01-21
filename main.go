package main

import "fmt"
import "github.com/gorilla/mux"
import "net/http"
import "strconv"
import "log"

func handler(w http.ResponseWriter, r *http.Request) {
	var vars = mux.Vars(r)
	var grid = vars["grid"]
	if len(grid) != 81 {
		var s = strconv.Itoa(http.StatusNotImplemented) + ": Grid must be 81 characters (9x9)"
		fmt.Println(s)
		http.Error(w, s, http.StatusNotImplemented)
		return
	}
	for i, c := range grid {
		if c < '0' || c > '9' {
			var s = strconv.Itoa(http.StatusNotImplemented) + ": Bad character is at position: " + strconv.Itoa(i)
			fmt.Println(s)
			http.Error(w, s, http.StatusNotImplemented)
			return
		}
	}
	var sol = SolveGrid(&grid)
	fmt.Fprintf(w, *sol)
}

func main() {
	var router = mux.NewRouter()
	router.HandleFunc("/",
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "GODOKU - Sudoku written in Go.  Enter a string of 81 characters ranging 0 - 9 for a sudoku board.  A zero is an empty slot.  Usage: /game/{string of 81 characters of 0-9}. Example /game/123050000450789023709003006204005097360000214090010300001002970000900530900501040")
		})
	router.HandleFunc("/game/{grid}", handler)
	router.HandleFunc("/health",
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "OK")
		})
	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":5000", router))
}
