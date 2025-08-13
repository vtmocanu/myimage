package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// ANSI escape codes for red bold text
	redBold := "\033[1;31m"
	reset := "\033[0m"
	
	fmt.Printf("%sHello METAPRO%s\n", redBold, reset)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// For web response, use HTML with inline CSS for red bold text
		w.Header().Set("Content-Type", "text/html")
		fmt.Fprintf(w, "<h1 style='color: red; font-weight: bold;'>Hello METAPRO</h1>\n")
	})

	log.Println("Server starting on port 80...")
	if err := http.ListenAndServe(":80", nil); err != nil {
		log.Fatal(err)
	}
}