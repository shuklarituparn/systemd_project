package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		return
	}
	rootMux := http.NewServeMux()
	rootMux.Handle("/status", http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json")
		err := json.NewEncoder(writer).Encode(map[string]string{
			"status": "up",
		})
		if err != nil {
			return
		}
	}))
	fmt.Print(fmt.Sprintf("Listening on port %s", os.Getenv("PORT")))
	serverError := http.ListenAndServe(fmt.Sprintf("0.0.0.0:%s", os.Getenv("PORT")), rootMux)
	if serverError != nil {
		return
	}
}
