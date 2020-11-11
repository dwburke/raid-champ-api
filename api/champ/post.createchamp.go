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

	if v, ok := obj["rarity"]; ok {
		champ.Rarity = cast.ToInt(v)
	}
	if v, ok := obj["affinity_id"]; ok {
		champ.AffinityId = cast.ToInt(v)
	}
	if v, ok := obj["faction_id"]; ok {
		champ.FactionId = cast.ToInt(v)
	}

	if dbc := db.Conn.Create(&champ); dbc.Error != nil {
		helpers.RespondWithError(w, 500, dbc.Error.Error())
		return
	}

	helpers.RespondWithJSON(w, 200, &champ)
}
