package champ

import (
	"encoding/json"
	"net/http"

	helpers "github.com/dwburke/go-tools/gorillamuxhelpers"
	//"github.com/gorilla/mux"
	"github.com/spf13/cast"

	"github.com/dwburke/raid-champ-api/db"
	"github.com/dwburke/raid-champ-api/types"
)

func CreateChamp(w http.ResponseWriter, r *http.Request) {
	conn := db.Open()

	champ := &types.Champ{}

	var obj map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&obj); err != nil {
		helpers.RespondWithError(w, 500, err.Error())
		return
	}

	if !helpers.CheckRequiredVar(w, obj, "name") {
		return
	}

	champ.Name = cast.ToString(obj["name"])

	if err := conn.Create(&champ).Error; err != nil {
		helpers.RespondWithError(w, 500, err.Error())
		return
	}

	helpers.RespondWithJSON(w, 200, &champ)
}
