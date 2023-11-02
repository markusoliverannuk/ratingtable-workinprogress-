package helper

import (
	"database/sql"
	"fmt"
	"net/http"
)

func AddTournamentInfoWhenSubmitted(w http.ResponseWriter, r *http.Request, db *sql.DB) error {
	TournamentName := r.FormValue("tournamentname")
	MainJudge := r.FormValue("mainjudgename")
	DateWhenHappened := r.FormValue("datewhenhappened")
	PlayerAmount := r.FormValue("playeramount")

	fmt.Println(TournamentName)
	fmt.Println(MainJudge)
	fmt.Println(DateWhenHappened)
	fmt.Println(PlayerAmount)

	insertSQL := `
	INSERT INTO tournaments (LISATUD, NIMI, TOIMUMISKUUP2EV, M2NGIJATEARV, PEAKOHTUNIK)
	VALUES (?, ?, ?, ?, ?)
`
	_, err3 := db.Exec(
		insertSQL,
		0,
		TournamentName,
		DateWhenHappened,
		PlayerAmount,
		MainJudge,
	)
	http.Redirect(w, r, "/addtournament", http.StatusSeeOther)
	//LoadTournamentAddPage(w, r)
	return err3

}
