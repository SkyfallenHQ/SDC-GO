package webHandler

import (
	"SkyfallenDeveloperCenter/config_parser"
	"SkyfallenDeveloperCenter/login"
	"SkyfallenDeveloperCenter/loginHandler"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"net/http"
)


func HandleHTTP(config config_parser.ConfigStructure, dbConnection *mongo.Client) {
	log.Printf("Initialising Skyfallen App Engine")

	log.Printf("Processing web handlers...")
	/* Begin Handlers */
	//http.HandleFunc("/", handleHome)
	http.HandleFunc("/developerid/login", func(w http.ResponseWriter, r *http.Request) {
		loginHandler.LoginHandle(w, r, config)
	})
	http.HandleFunc("/oauth/callback", func(w http.ResponseWriter, r *http.Request) {
		login.HandleOauthCallback(w, r, dbConnection, config)
	})

	http.HandleFunc("/developerid/auth", func(w http.ResponseWriter, r *http.Request) {
		login.AuthenticateUser(w, r, dbConnection, config)
	})

	http.HandleFunc("/static/",handleStatic)
	/* End Handlers */

	log.Printf("Running App Engine")
	log.Printf("Listening for connections...")
	log.Fatal(http.ListenAndServe(":80", nil))
}