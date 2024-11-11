package LineCounter

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// CountHandler handles the /count endpoint
func CountHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodGet {
		http.Error(w, InvalidMethodCall, http.StatusMethodNotAllowed)
		log.Println("Invalid method call")
		return
	}

	responseBody := LineCounterResponse{}

	folderPath := r.URL.Query().Get("folderPath")
	if folderPath == "" {
		responseBody.Message = InvalidParam
		responseJSON, err := json.Marshal(responseBody)
		if err != nil {
			http.Error(w, "Failed to encode response JSON", http.StatusInternalServerError)
			log.Println("Error marshalling response JSON")
			return
		}

		w.WriteHeader(http.StatusBadRequest)
		_, err = w.Write(responseJSON)
		if err != nil {
			return
		}
		return
	}

	result, err := analyzeFolder(folderPath)
	if err != nil {
		responseBody.Message = fmt.Sprintf(ErrLineCounter, err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Error while analysing folder", err.Error())
	} else {
		responseBody.Message = LineCounted
		responseBody.Body = result
		w.WriteHeader(http.StatusOK)
	}

	responseJSON, err := json.Marshal(responseBody)
	if err != nil {
		http.Error(w, "Failed to encode response JSON", http.StatusInternalServerError)
		log.Println("Error marshalling response JSON")
		return
	}

	w.WriteHeader(http.StatusBadRequest)
	_, err = w.Write(responseJSON)
	if err != nil {
		return
	}
}
