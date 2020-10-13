package api

import (
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	helpers "github.com/dwburke/go-tools/gorillamuxhelpers"
	"github.com/spf13/viper"
)

func Login(w http.ResponseWriter, r *http.Request) {
	username, ok := r.Context().Value(usernameKey("username")).(string)

	if !ok || username == "" {
		helpers.RespondWithJSON(w, 500, map[string]interface{}{
			"login": false,
			"error": "Could not determine username.",
		})

		return
	}

	expires := time.Now().Add(viper.GetDuration("auth.token.lifespan")).Unix()

	claims := &jwt.StandardClaims{
		ExpiresAt: expires,
		Issuer:    "prov-api",
		Subject:   username,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString([]byte(viper.GetString("auth.token.secret")))

	if err != nil {
		helpers.RespondWithError(w, 500, err.Error())
		return
	}

	helpers.RespondWithJSON(w, 200, map[string]interface{}{
		"login":    true,
		"token":    ss,
		"username": username,
		"expires":  expires,
	})
}
