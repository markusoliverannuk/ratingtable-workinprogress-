package helper

import (
	"net/http"
)

func RatingLoadHandlerFilter(w http.ResponseWriter, r *http.Request) {
	nameSearch := r.FormValue("name")
	clubSearch := r.FormValue("club")
	ratingRange := r.FormValue("ratingRange")
	birthYearSearch := r.FormValue("birthyear")

	players, err := RetrievePlayerDataFromDatabase()

	if r.Method == "GET" {
		if birthYearSearch != "" && clubSearch != "" && ratingRange != "all" {
			players, err = RetrievePlayerDataFromDatabaseByBirthYearAndClubAndRange(birthYearSearch, clubSearch, ratingRange)
		} else if nameSearch != "" && clubSearch != "" {
			players, err = RetrievePlayerDataFromDatabaseByNameAndClub(nameSearch, clubSearch)
		} else if birthYearSearch != "" && clubSearch != "" {
			players, err = RetrievePlayerDataFromDatabaseByBirthYearAndClub(birthYearSearch, clubSearch)
		} else if nameSearch != "" {
			players, err = RetrievePlayerDataFromDatabaseByName(nameSearch)
		} else if clubSearch != "" {
			players, err = RetrievePlayerDataFromDatabaseByClub(clubSearch)
			/*} else if ratingRange != "all" {
			players, err = RetrievePlayerDataFromDatabaseByRange(ratingRange)*/
		} else if birthYearSearch != "" {
			players, err = RetrievePlayerDataFromDatabaseByBirthYear(birthYearSearch)
		}
	}

	if err != nil {
		http.Error(w, "Error retrieving player data", http.StatusInternalServerError)
		return
	}

	LoadRatingPage(w, r, players)
}
func RatingLoadHandler(w http.ResponseWriter, r *http.Request) {

	players, err := RetrievePlayerDataFromDatabase()
	if err != nil {
		http.Error(w, "Error retrieving player data", http.StatusInternalServerError)
		return
	}
	LoadRatingPage(w, r, players)
}
