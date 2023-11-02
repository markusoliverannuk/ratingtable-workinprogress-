package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"ratingtable/helper"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func main() {

	helper.Db = helper.DatabaseInitialization("leaderboard.db")
	helper.CreateDbTablePlayers(helper.Db)

	helper.CreateDbTableTournaments(helper.Db)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	if err := helper.ProcessXMLFile("players.xml"); err != nil {
		fmt.Println("Error:", err)
		return
	}

	http.HandleFunc("/", helper.RatingLoadHandler)

	http.HandleFunc("/reiting", helper.RatingLoadHandler)                    //displaying/loading player leaderboard
	http.HandleFunc("/filter", helper.RatingLoadHandlerFilter)               //filtering player leaderboard
	http.HandleFunc("/addtournament", helper.AddTournamentPageLoadHandler)   //loading tournamnet adder page
	http.HandleFunc("/addtournamentaction", helper.AddTournamentInfoHandler) // adding new tournament info
	http.HandleFunc("/insertgameinfo", helper.GameInsertHandler)             // inserting info about games (score, players)
	log.Fatal(http.ListenAndServe(":3089", nil))
	fmt.Println("hey")

}
