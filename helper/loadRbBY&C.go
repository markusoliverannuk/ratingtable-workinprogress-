package helper

import (
	"log"
)

func RetrievePlayerDataFromDatabaseByBirthYearAndClub(birthyear string, clubname string) ([]Person, error) {
	var players []Person
	rows, err := Db.Query("SELECT * FROM leaderboard WHERE SA LIKE ? AND KLUBI LIKE ?", "%"+birthyear+"%", "%"+clubname+"%")
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
