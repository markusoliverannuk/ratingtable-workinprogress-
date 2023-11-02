package helper

import (
	"fmt"
	"net/http"
)

func LoadRatingPage(w http.ResponseWriter, r *http.Request, players []Person) {
	data := struct {
		Players []Person
	}{
		Players: players,
	}

	e := ratingPageTMPLT.Execute(w, data)
	if e != nil {
		fmt.Println(e)
		http.Error(w, "Error while rendering ELTL reiting!", http.StatusInternalServerError)
	}

}
