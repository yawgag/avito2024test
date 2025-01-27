package houses

import (
	"fmt"
	"net/http"
)

func CreateNewHouse(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "../../templates/createHouse.html")
}

func ShowHouseInfo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "house info")
}
