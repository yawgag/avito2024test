package mainPage

import (
	"net/http"
)

func MainPageHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "../../templates/mainPage.html")
}
