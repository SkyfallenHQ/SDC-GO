package core

import (
	"SkyfallenDeveloperCenter/config_parser"
	"SkyfallenDeveloperCenter/webHandler"
	"fmt"
	"github.com/gorilla/securecookie"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"SkyfallenDeveloperCenter/database_connector"
	"os"
)

var config config_parser.ConfigStructure
var dbConnnection *mongo.Client

func Handle(){

	log.Printf("Welcome to Skyfallen App Engine, running Skyfallen Developer Center")
	fmt.Printf(" ____  _           __       _ _            \n/ ___|| | ___   _ / _| __ _| | | ___ _ __  \n\\___ \\| |/ / | | | |_ / _` | | |/ _ \\ '_ \\ \n ___) |   <| |_| |  _| (_| | | |  __/ | | |\n|____/|_|\\_\\\\__, |_|  \\__,_|_|_|\\___|_| |_|\n            |___/                          \n ____                 _                          ____           _            \n|  _ \\  _____   _____| | ___  _ __   ___ _ __   / ___|___ _ __ | |_ ___ _ __ \n| | | |/ _ \\ \\ / / _ \\ |/ _ \\| '_ \\ / _ \\ '__| | |   / _ \\ '_ \\| __/ _ \\ '__|\n| |_| |  __/\\ V /  __/ | (_) | |_) |  __/ |    | |__|  __/ | | | ||  __/ |   \n|____/ \\___| \\_/ \\___|_|\\___/| .__/ \\___|_|     \\____\\___|_| |_|\\__\\___|_|   \n                             |_|                                            \n")

	log.Printf("Generating a secure cookie key, please wait")
	os.Setenv("SDC_SESSION_KEY",string(securecookie.GenerateRandomKey(32)))

	config = config_parser.ParseConfig()
	log.Printf("Successfully read config, proceeding with database connection")
	dbConnnection = database_connector.Connect(config.Contents.Database)
	webHandler.HandleHTTP(config,dbConnnection)

}