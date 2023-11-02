package helper

import "html/template"

var ratingPageTMPLT = template.Must(template.ParseFiles("templates/leaderboard.html"))
var tournamentAddPageTMPLT = template.Must(template.ParseFiles("templates/addtournamentinfo.html"))