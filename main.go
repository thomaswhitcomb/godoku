package main

import "fmt"
import "net/http"
import "strconv"

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path[1:])
	var grid = r.URL.Path[1:]
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
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
	fmt.Println("waiting to serve")
}
