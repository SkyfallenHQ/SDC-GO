package login

import (
	"SkyfallenDeveloperCenter/config_parser"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"net/http"
)

func AuthenticateUser(w http.ResponseWriter, r *http.Request, database *mongo.Client, cnf config_parser.ConfigStructure){

	log.Printf("Handling user authentication...")

	log.Printf("Extracting the session store out of the cookie store")

	SessionStore, _ := CookieStore.Get(r, "SDCSession")

	database.Database("SDC-DB").Collection("Users").InsertOne()

}
