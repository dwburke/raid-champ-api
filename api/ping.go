package api

import (
	"net/http"

	helpers "github.com/dwburke/go-tools/gorillamuxhelpers"
)

func Ping(w http.ResponseWriter, r *http.Request) {
	helpers.RespondWithJSON(w, 200, map[string]interface{}{
		"ping": 1,
	})
}
