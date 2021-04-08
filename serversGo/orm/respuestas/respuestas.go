package respuestas

import (
	"encoding/json"
	"net/http"
)

func ResponderJSON(obj interface{}, w http.ResponseWriter) {
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(obj)
}
