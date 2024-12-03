package handlers

import (
	"chatserver/models"
	"chatserver/services"
	"chatserver/utils"
	"encoding/json"
	"net/http"
	"time"
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
	result, err := services.LoginUser(r.Context(), user)
	if err != nil {
		http.Error(w, "Failed to register user", http.StatusInternalServerError)
		return
	}

	// Construct the response
	accessToken,accessMaxAge, err := utils.CreateToken(user.Email,time.Hour)
	if err != nil {
		http.Error(w, "Error creating token", http.StatusInternalServerError)
		return
	}
	refreshToken,refreshMaxAge, err := utils.CreateToken(user.Email, 7*24*time.Hour) // Refresh token valid for 7 days
	if err != nil {
		http.Error(w, "Error creating refresh token", http.StatusInternalServerError)
		return
	}

    // Set auth cookies
	utils.SetAuthCookies(w, accessToken, refreshToken, accessMaxAge, refreshMaxAge)
	response := map[string]interface{}{
		"message":      "User logged in successfully",
		"accessToken":  accessToken,
		"refreshToken": refreshToken,
		"data" : result,
	}
	
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
	result, err := services.RegisterUser(r.Context(), user)
	if err != nil {
		http.Error(w, "Failed to register user", http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"message": "User registered successfully",
		"data": result,
	}

	utils.SendResponse(w, http.StatusCreated, response)
}
