package kudoserver

import (
	"log"
	"main/kudoroutes"
	"net/http"
)

func StartDefaultServer() error {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /users", kudoroutes.DefaultUserGET)
	mux.HandleFunc("POST /users", kudoroutes.DefaultUserPOST)

	log.Println("Server starting on :8080...")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		return err
	}

	return nil
}
