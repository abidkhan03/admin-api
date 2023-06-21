package auth

import (
	"net/http"

	"github.com/spongeling/admin-api/internal/dao"
	"github.com/spongeling/admin-api/internal/response"
)

var users []*dao.User

func UpdateUsers(newUsers []*dao.User) {
	users = newUsers
}

// Authenticator is a middleware for authentication using BasicAuth
func Authenticator(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		email, password, ok := r.BasicAuth()
		if !ok {
			response.RespondErrorMessage(w, http.StatusInternalServerError, "invalid credentials")
			// http.Error(w, "invalid credentials", http.StatusInternalServerError)
			return
		}

		for _, user := range users {
			if user.Email == email && user.Password == password {
				next.ServeHTTP(w, r)
				return
			}
		}
		response.RespondErrorMessage(w, http.StatusUnauthorized, "invalid username or password")
	}

}
