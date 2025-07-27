package responses

import(
	"fmt"
	"net/http"
	"encoding/json"
)

func Json(w http.ResponseWriter, statusCode int, data any) {
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(data)

	if err != nil {
		fmt.Fprintf(w, "%s", err.Error())
	}
}

func Error(w http.ResponseWriter, statusCode int, err error) {
	if err != nil {
		Json(w, statusCode, struct{
			Error string `json:"error"`
		}{
			Error: err.Error(),
		})

		return
	}

	Json(w, http.StatusBadRequest, nil)
}