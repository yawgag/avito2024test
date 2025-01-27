package user

import (
	"fmt"
	"net/http"
	"time"
)

func AuthChecker(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := r.Cookie("userRole")
		if err != nil {
			http.Error(w, "something went wrong. ", http.StatusForbidden)
		} else {
			next.ServeHTTP(w, r)
		}
	})
}

/*
add tokens for check
*/
func ModeratorChecker(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		roleCookie, err := r.Cookie("userRole")
		if err != nil {
			http.Error(w, "something went wrong", http.StatusForbidden)
		} else if roleCookie.Value == "user" {
			http.Error(w, "wrong user role", http.StatusForbidden)
		} else {
			next.ServeHTTP(w, r)
		}
	})
}

/*
complete this 2 functions below
add tolens for differents roles
*/
func Login(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "../../templates/login.html")
}

func Register(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "../../templates/register.html")
}

func DummyLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		if err := r.ParseForm(); err != nil {
			http.Error(w, "something went wrong", http.StatusBadRequest)
		}
		role := r.FormValue("role")
		userRoleCookie := http.Cookie{
			Name:    "role",
			Value:   role,
			Expires: time.Now().Add(time.Hour * 10),
		}
		fmt.Println(role)
		http.SetCookie(w, &userRoleCookie)

		http.Redirect(w, r, "/", http.StatusFound)
	} else {
		http.ServeFile(w, r, "../../templates/dummyLogin.html")
	}

}
