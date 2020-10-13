package api

import (
	"context"
	"errors"
	"net/http"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"github.com/dwburke/raid-champ-api/db"
	"github.com/dwburke/raid-champ-api/types"
)

type usernameKey string

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		username, ok := checkAuth(r)

		if ok {
			ctx := context.WithValue(r.Context(), usernameKey("username"), username)
			next.ServeHTTP(w, r.WithContext(ctx))
		} else {
			w.Header().Set("WWW-Authenticate", "Basic realm=Prov")
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
		}
	})
}

func checkAuth(r *http.Request) (string, bool) {
	path := "UNKNOWN"

	if route := mux.CurrentRoute(r); route != nil {
		t, err := route.GetPathTemplate()
		if err != nil {
			log.Error(err)
			return "", false
		}

		path = t
	}

	if path == "/ping" {
		return "", true
	}

	var user *types.ApiUser
	var err error

	authHeader := r.Header.Get("Authorization")

	if strings.HasPrefix(authHeader, "Bearer ") {
		token := strings.TrimPrefix(authHeader, "Bearer ")
		user, err = authenticateToken(token)
	} else {
		username, password, ok := r.BasicAuth()
		if !ok {
			log.Debug("Couldn't parse basic auth header")
			return "", false
		}
		user, err = authenticateUserPass(username, password)
	}

	if err != nil {
		log.Errorf("Error processing credentials: %v", err)
		return "", false
	}

	if isUserAuthorized(user, path, r.Method) {
		return user.Username, true
	}

	return "", false
}

func authenticateToken(rawToken string) (*types.ApiUser, error) {
	token, err := jwt.ParseWithClaims(rawToken, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(viper.GetString("auth.token.secret")), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("Invalid token")
	}

	claims, ok := token.Claims.(*jwt.StandardClaims)
	if !ok {
		return nil, errors.New("Invalid token")
	}

	provdb := db.Open()

	user := &types.ApiUser{}
	if err := provdb.Where("username = ?", claims.Subject).First(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func authenticateUserPass(username, password string) (*types.ApiUser, error) {
	user := &types.ApiUser{}

	provdb := db.Open()

	if err := provdb.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}

	if user.CheckPassword(password) {
		return user, nil
	}

	return nil, errors.New("not authorized")
}

func isUserAuthorized(user *types.ApiUser, route, method string) bool {
	authorized, err := user.IsAuthorized(route, method)
	if err != nil {
		log.Error(err)
		return false
	}

	return authorized
}
