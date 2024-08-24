package server

import (
	"fmt"
	"net/http"
	"tucows-challenge/model"
)

func GetMenu(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Menu: %+v", model.Menu)
}
