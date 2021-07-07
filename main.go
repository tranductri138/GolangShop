package main

import (
	"fmt"
	"net/http"
)

func echo(){
	
}
func main() {
	fmt.Println("tris dep zai")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

	})
	http.ListenAndServe(":3000", nil)
}
