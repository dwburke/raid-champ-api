package api

import (
	"context"
	"fmt"
	"net"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	//endpoint_agents "github.com/dwburke/raid-champ-api/api/agents"
)

var ShutdownCh chan bool
var server *http.Server

func init() {
	viper.SetDefault("api.server.enabled", false)
	viper.SetDefault("api.server.address", "127.0.0.1")
	viper.SetDefault("api.server.port", 8443)
	viper.SetDefault("api.server.ssl-key", "key.pem")
	viper.SetDefault("api.server.ssl-cert", "cert.pem")
	viper.SetDefault("api.server.https", false)
}

func SetupRoutes(r *mux.Router) {
	r.HandleFunc("/ping", Ping).Methods("GET")
	r.HandleFunc("/login", Login).Methods("GET")

	//endpoint_agents.SetupRoutes(r)

	// display all endpoints that were created by called package(s)
	r.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		t, err := route.GetPathTemplate()
		if err != nil {
			log.Panic(err)
		}

		methods, err := route.GetMethods()
		if err != nil {
			log.Panic(err)
		}

		log.Printf("%s %#v\n", t, methods)
		return nil
	})
}

func Run() {
	if !viper.GetBool("api.server.enabled") {
		log.Println("api.server.enabled == false; not starting")
		return
	}

	var listen string = fmt.Sprintf("%s:%d", viper.GetString("api.server.address"), viper.GetInt("api.server.port"))

	log.WithFields(log.Fields{
		"api.server.https":    viper.GetBool("api.server.https"),
		"listen":              listen,
		"api.server.ssl-cert": viper.GetString("api.server.ssl-cert"),
		"api.server.ssl-key":  viper.GetString("api.server.ssl-key"),
	}).Info("api: starting")

	log.Info("api: not running on ", listen, "; starting")

	ShutdownCh = make(chan bool)

	r := mux.NewRouter()
	r.Use(loggingMiddleware)
	r.Use(AuthMiddleware)
	SetupRoutes(r)

	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Authorization", "Content-Type"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	//originsOk := handlers.AllowedOrigins([]string{os.Getenv("ORIGIN_ALLOWED")})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS", "DELETE"})

	server = &http.Server{
		Addr:    listen,
		Handler: handlers.CORS(originsOk, headersOk, methodsOk)(r),
	}

	server.SetKeepAlivesEnabled(false)

	go func() {
		if viper.GetBool("api.server.https") {
			if err := server.ListenAndServeTLS(
				viper.GetString("api.server.ssl-cert"),
				viper.GetString("api.server.ssl-key"),
			); err != http.ErrServerClosed {
				log.Fatalf("api: %s", err)
			}
		} else {
			if err := server.ListenAndServe(); err != http.ErrServerClosed {
				log.Fatalf("api: %s", err)
			}
		}
	}()

	log.Info("api: running")
}

func apiRunning() bool {
	var address string = fmt.Sprintf("%s:%d", viper.GetString("api.server.address"), viper.GetInt("api.server.port"))

	conn, err := net.Dial("tcp", address)
	if err != nil {
		return false
	}

	conn.Close()

	return true
}

func Shutdown() {
	if server != nil {
		log.Info("api: [shutdown] shutting down")
		if err := server.Shutdown(context.TODO()); err != nil {
			log.Panic("api: ", err)
		}

		server = nil
	} else {
		log.Info("api: [shutdown] server == nil in this process")
	}

	if ShutdownCh != nil {
		log.Info("api: [shutdown] signaling shutdown channel")
		ShutdownCh <- true
		close(ShutdownCh)
		ShutdownCh = nil
	} else {
		log.Info("api: [shutdown] ShutdownCh == nil in this process")
	}
}
