package authentification

import (
	"encoding/json"
	"net/http"

	users "forum/server/api/user"
)

var Sessions = map[string]users.User{}

func GetSession(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session_token")
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Session not found"))
		return
	}

	// Vérifie si la session existe dans la map des sessions
	sessionData, exists := Sessions[cookie.Value]
	if !exists {
		cookie.MaxAge = -1
		http.SetCookie(w, cookie)
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Invalid session"))
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(sessionData)
}
