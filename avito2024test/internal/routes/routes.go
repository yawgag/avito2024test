package routes

import (
	"avito2024test/internal/houses"
	"avito2024test/internal/mainPage"
	"avito2024test/internal/user"
	"net/http"

	"github.com/gorilla/mux"
)

func SetupRoutes() *mux.Router {
	houseMux := mux.NewRouter()

	houseMux.HandleFunc("/house/create", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "../../templates/createHouse.html")
	})
	houseMux.HandleFunc("/house/{id:[0-9]+}", houses.ShowHouseInfo)

	houseHandler := user.ModeratorChecker(houseMux)

	router := mux.NewRouter()

	router.HandleFunc("/login", user.Login)
	router.HandleFunc("/register", user.Register)
	router.HandleFunc("/dummyLogin", user.DummyLogin)
	router.Handle("/house/", houseHandler)

	router.HandleFunc("/", mainPage.MainPageHandler)

	return router
}
