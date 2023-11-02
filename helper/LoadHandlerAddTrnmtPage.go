package helper

import "net/http"

func AddTournamentPageLoadHandler(w http.ResponseWriter, r *http.Request) {

	tournaments, err := RetrieveTournamentsDataFromDatabase()
	if err != nil {
		http.Error(w, "Error retrieving player data", http.StatusInternalServerError)
		return
	}
	LoadTournamentAddPage(w, r, tournaments)
}
