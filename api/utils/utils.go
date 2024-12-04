package utils 

import (
    "encoding/json"
    "log"
    "net/http"
)
func SendResponse(w http.ResponseWriter, statusCode int, data interface{}) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(statusCode)

    if err := json.NewEncoder(w).Encode(data); err != nil {
        log.Println("Error encoding success response:", err)
        http.Error(w, "Failed to send response", http.StatusInternalServerError)
    }
}

func SendErrorResponse(w http.ResponseWriter, statusCode int, errorMessage string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	response := map[string]interface{}{
		"success": false,
		"error":   errorMessage,
	}
	json.NewEncoder(w).Encode(response)
}


