/*
package config_parser

import (
	"io/ioutil"
	"fmt"
	"log"
	"encoding/json"
	"os"
)

func ParseConfig(){

	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Getting Current Working Directory")
	log.Printf("Acquired path: "+path)

	log.Printf("Parsing Config")
	configText, error := ioutil.ReadFile("config.json")
	if error != nil {

		log.Fatal("Cannot open config file")

	} else {

		log.Printf("Acquired config")

	}

	var configStore config_file

	log.Printf("Reading config")
	json.Unmarshal(configText, &configStore)

	fmt.Printf(string(configText))

	log.Printf("Initiating DB Connection")
	fmt.Printf("%+v\n",configStore)

}
*/
package config_parser

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

// Users struct which contains
// an array of users
type Users struct {
	Users User `json:"configuration"`
}

// User struct which contains a name
// a type and a list of social links
type User struct {
	Name string `json:"database"`
	Type   string `json:"type"`
	Age    int    `json:"Age"`
	Social Social `json:"social"`
}

// Social struct which contains a
// list of links
type Social struct {
	Facebook string `json:"facebook"`
	Twitter  string `json:"twitter"`
}

func ParseConfig() {
	// Open our jsonFile
	jsonFile, err := os.Open("config.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened users.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// read our opened xmlFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// we initialize our Users array
	var users Users

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	json.Unmarshal(byteValue, &users)

	// we iterate through every user within our users array and
	// print out the user Type, their name, and their facebook url
	// as just an example
		fmt.Println("User Type: " + users.Users.Type)
		fmt.Println("User Age: " + strconv.Itoa(users.Users.Age))
		fmt.Println("User Name: " + users.Users.Name)
		fmt.Println("Facebook Url: " + users.Users.Social.Facebook)

}