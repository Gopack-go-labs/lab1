package main

import (
	// "encoding/json"
	"net/http"
	// "time"
)

const PORT = ":8795"

func main() {
	http.HandleFunc("/time", getTime)

  http.ListenAndServe(PORT, nil)
}
