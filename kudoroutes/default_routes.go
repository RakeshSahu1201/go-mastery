package kudoroutes

import (
	"encoding/json"
	"main/kudostore"
	"net/http"
)

func UserGET(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "appplication/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(kudostore.Store)
}
