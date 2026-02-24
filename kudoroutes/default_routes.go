package kudoroutes

import (
	"encoding/json"
	"errors"
	"main/kudomodels"
	"main/kudoservice"
	"main/kudostore"
	"main/kudotypes"
	"net/http"
)

func UserGET(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "appplication/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(kudostore.Store)
}

func UserPOST(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Protect against large payloads (1MB limit)
	r.Body = http.MaxBytesReader(w, r.Body, 1<<20)

	var req kudotypes.UserRequest

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(
			errors.New("invalid request payload"),
		)

		return
	}

	// Validate input
	errs := kudoservice.ValidateUserRequest(&req)
	if len(errs) > 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errs)
		return
	}

	// Create user
	user := kudomodels.CreateUser(req.Name, req.Email)

	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(user); err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
		return
	}
}
