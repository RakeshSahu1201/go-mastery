package kudoserver

import (
	"log"
	"main/kudoroutes"
	"net/http"
)

func SartServer() error {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", kudoroutes.UserGET)

	log.Println("Server starting on :8080...")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		return err
	}

	return nil
}
