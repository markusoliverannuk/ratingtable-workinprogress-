package helper

import "log"

func RetrievePlayerDataFromDatabase() ([]Person, error) {
	var players []Person

	rows, err := Db.Query("SELECT * FROM leaderboard")
	if err != nil {
		log.Fatal(err)
		return players, err
	}
	defer rows.Close()

	for rows.Next() {
		var player Person
		err := rows.Scan(
			&player.RateOrder,
			&player.RatePlpnts,
			&player.RatePoints,
			&player.RateWeight,
			&player.PersonID,
			&player.FirstName,
			&player.FamName,
			&player.BirthDate,
			&player.ClbName,
		)
		if err != nil {
			log.Fatal(err)
			return players, err
		}
		players = append(players, player)
	}

	return players, nil
}

func RetrieveTournamentsDataFromDatabase() ([]Tournament, error) {
	var tournaments []Tournament

	rows, err := Db.Query("SELECT * FROM tournaments")
	if err != nil {
		log.Fatal(err)
		return tournaments, err
	}
	defer rows.Close()

	for rows.Next() {
		var tournament Tournament
		err := rows.Scan(
			&tournament.AddedWhen,
			&tournament.Name,
			&tournament.HappenedWhen,
			&tournament.AmountOfPlayers,
			&tournament.MainJudge,
		)
		if err != nil {
			log.Fatal(err)
			return tournaments, err
		}
		tournaments = append(tournaments, tournament)
	}

	return tournaments, nil
}
