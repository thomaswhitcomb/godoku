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
	router.HandleFunc("/run/{grid}", handler)
	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":8080", router))
}
