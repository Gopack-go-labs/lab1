package main

import (
	"encoding/json"
	"net/http"
	"time"
)

const PORT = ":8795"

func main() {
	http.HandleFunc("/time", func(w http.ResponseWriter, r *http.Request) {
		currentTime := time.Now().Format(time.RFC3339)
		response := map[string]string{"time": currentTime}
		jsonResponse, err := json.Marshal(response)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonResponse)
	})

  http.ListenAndServe(PORT, nil)
}
