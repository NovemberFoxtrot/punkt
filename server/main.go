package server

import (
	"net/http"
	"fmt"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "punkt")
}

func Init() {
	http.HandleFunc("/", Index)
	http.ListenAndServe(":8080", nil)
}
