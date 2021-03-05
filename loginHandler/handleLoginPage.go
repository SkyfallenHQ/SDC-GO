package loginHandler

import (
	"SkyfallenDeveloperCenter/config_parser"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path"
	"strings"
	"time"
)

func LoginHandle(w http.ResponseWriter, r *http.Request, config config_parser.ConfigStructure){

	startTime := time.Now()
	log.Printf("Received GET on login page, progressing with the page load")
	r.Header.Set("X-Powered-By","SkyfallenAppEngine")

	log.Printf("Reading template")
	fp := path.Join("web_templates", "developerIDLogin.html")
	tmpl, err := template.ParseFiles(fp)
	if err != nil {
		log.Printf("Unhandled error, can not parse the template file ")
		log.Printf(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		log.Printf("/developerid/login - 500 Internal Error")
		return
	}

	oauthCallbackNE :=  config.Contents.WebPath+"oauth/callback"
	oauthCallback := strings.ReplaceAll(string(oauthCallbackNE), "%", "%%")

	tmpltData := &LoginPagePassThrough{WebPath: config.Contents.WebPath, AuthEndpoint: config.Contents.IDP.AuthURL, AuthClientID: config.Contents.IDP.ClientID, EncodedRedirectURI: oauthCallback}

	log.Printf("Escaping Oauth Callback for "+oauthCallback)

	r.Header.Set("Content-Type","text/html")
	execErr := tmpl.Execute(w, tmpltData);

	if execErr != nil {
		log.Printf("Unhandled error, can not execute the template file")
		log.Printf(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		log.Printf("/developerid/login - 500 Internal Error")
		return
	} else {

		log.Printf("/developerid/login - 200 OK")
		elapsed := time.Since(startTime).Seconds()*1000
		log.Println(fmt.Sprintf("Served Login HTML file in %v ms", elapsed))

	}

}