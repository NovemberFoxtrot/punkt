package server

import (
	"net/http"
)

func Init() {
	http.ListenAndServe(":8080", nil)
}
