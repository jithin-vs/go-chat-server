package utils

import (
	"net/http"
	"time"
)

func SetAuthCookies(w http.ResponseWriter, accessToken string, refreshToken string, accessExp int64, refreshExp int64) {
	// Calculate remaining time for MaxAge
	accessMaxAge := int(accessExp - time.Now().Unix())
	refreshMaxAge := int(refreshExp - time.Now().Unix())

	// Set Access Token Cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "accessToken",
		Value:    accessToken,
		Path:     "/",
		HttpOnly: true,
		Secure:   false, // Set to true in production
		MaxAge:   accessMaxAge,
	})

	// Set Refresh Token Cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "refreshToken",
		Value:    refreshToken,
		Path:     "/",
		HttpOnly: true,
		Secure:   false, // Set to true in production
		MaxAge:   refreshMaxAge,
	})
}