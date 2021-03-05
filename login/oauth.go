package login

import (
	"SkyfallenDeveloperCenter/config_parser"
	"encoding/json"
	"fmt"
	"github.com/gorilla/sessions"
	"go.mongodb.org/mongo-driver/mongo"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
)

var Oauth_bearer_resp oauth_bearer_request_response

var ReturnedIdentity oauth_identity_response

var CookieStore = sessions.NewCookieStore([]byte(os.Getenv("SDC_SESSION_KEY")))

func HandleOauthCallback(w http.ResponseWriter, r *http.Request, db *mongo.Client, cnf config_parser.ConfigStructure) {

	log.Printf("Extracting the session store out of the cookie store")
	SessionStore, _ := CookieStore.Get(r, "SDCSession")

	log.Printf("Received Oauth Callback from Skyfallen ID")

	code, ok := r.URL.Query()["code"]

	if !ok || len(code[0]) < 1 {
		log.Println("Oauth Code wasn't returned")
		fmt.Fprint(w,"NO CODE RETURNED")
		return
	}

	endpoint := cnf.Contents.IDP.AuthURL+"/oauth/token"
	data := url.Values{}
	data.Set("grant_type", "authorization_code")
	data.Set("code", code[0])
	data.Set("client_id", cnf.Contents.IDP.ClientID)
	data.Set("client_secret", cnf.Contents.IDP.ClientSecret)
	data.Set("redirect_uri", url.PathEscape(cnf.Contents.WebPath+"oauth/callback"))

	client := &http.Client{}
	r, err := http.NewRequest("POST", endpoint, strings.NewReader(data.Encode())) // URL-encoded payload
	if err != nil {
		log.Fatal(err)
	}
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	res, err := client.Do(r)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Skyfallen ID Auth API: "+res.Status)
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Skyfallen API Response (Request Bearer): "+string(body))

	json.Unmarshal(body,&Oauth_bearer_resp)

	/*
	if Oauth_bearer_resp.Error == "invalid_grant" {

		log.Printf("Skyfallen ID has returned an error: "+Oauth_bearer_resp.Error)
		fmt.Fprint(w, string(body))
		return

	} */

	respMe, errMe := http.Get(cnf.Contents.IDP.AuthURL+"/?oauth=me&access_token="+Oauth_bearer_resp.AccessToken)
	if errMe != nil {
		log.Fatal(errMe)
	}

	bodyMe, errP := ioutil.ReadAll(respMe.Body)
	if errP != nil {
		log.Fatal(errP)
	}

	stringMe := string(bodyMe)
	log.Printf("Oauth ME Returned from Skyfallen ID: "+stringMe)

	json.Unmarshal(bodyMe,&ReturnedIdentity)


	log.Printf("Saving Callback State to the session storage.")
	SessionStore.Values["LoginState"] = "LOGGEDIN"
	SessionStore.Values["Username"] = ReturnedIdentity.Username
	SessionStore.Values["Email"] = ReturnedIdentity.Email
	SessionStore.Values["ID"] = ReturnedIdentity.ID
	SessionStore.Values["DisplayName"] = ReturnedIdentity.DisplayName
	SessionStore.Values["AccountCreation"] = ReturnedIdentity.CreationDate

	SesSaveErr := SessionStore.Save(r, w)
	if SesSaveErr != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	} else {

		http.Redirect(w, r, cnf.Contents.WebPath+"developerid/auth", http.StatusSeeOther)

	}

}
