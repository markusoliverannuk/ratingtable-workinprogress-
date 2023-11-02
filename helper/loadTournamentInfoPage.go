package helper

import (
	"fmt"
	"net/http"
)

func LoadTournamentAddPage(w http.ResponseWriter, r *http.Request, tournaments []Tournament) {
	data := struct {
		Tournaments []Tournament
	}{
		Tournaments: tournaments,
	}

	e := tournamentAddPageTMPLT.Execute(w, data)
	if e != nil {
		fmt.Println(e)
		http.Error(w, "Error while rendering tournament adder page!", http.StatusInternalServerError)
	}
}
