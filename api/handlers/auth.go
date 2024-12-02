package handlers

import (
	"chatserver/models"
	"chatserver/services"
	"chatserver/utils"
	"encoding/json"
	"net/http"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
     
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
    if err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }
	
    // Insert the user into MongoDB
	response, err := services.LoginUser(r.Context(), user)
	if err != nil {
		http.Error(w, "Failed to register user", http.StatusInternalServerError)
		return
	}

	// Construct the response


	utils.SendResponse(w, http.StatusOK, response)
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	
	
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
    if err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }
	
    // Insert the user into MongoDB
	response, err := services.RegisterUser(r.Context(), user)
	if err != nil {
		http.Error(w, "Failed to register user", http.StatusInternalServerError)
		return
	}

	// Construct the response


	utils.SendResponse(w, http.StatusOK, response)
}
