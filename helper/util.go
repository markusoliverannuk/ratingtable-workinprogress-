package helper

import "database/sql"

func getRatingPLPoints(db *sql.DB, id string) (int, error) {
	var ratingPoints int
	err := db.QueryRow("SELECT PP FROM leaderboard WHERE ID = ?", id).Scan(&ratingPoints)
	if err != nil {
		return 0, err
	}
	return ratingPoints, nil
}

func getWeightPoints(db *sql.DB, id string) (int, error) {
	var weightPoints int
	err := db.QueryRow("SELECT KL FROM leaderboard WHERE ID = ?", id).Scan(&weightPoints)
	if err != nil {
		return 0, err
	}
	return weightPoints, nil
}
