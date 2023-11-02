package helper

import (
	"net/http"
)

func GameInsertHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
			GameInsertExecution(w, r, Db)

	}
}
