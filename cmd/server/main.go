package main

import (
	"fmt"
	"net/http"
)

/*
HANDLER FUNCTION TO THE MAIN PAGE
ResponseWriter allows sending a response back to the client.
Request provides information about the request.
*/
func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to my Book Store!\n")
}

func main() {
	http.HandleFunc("/", homeHandler)

	http.ListenAndServe(":8080", nil)
}
