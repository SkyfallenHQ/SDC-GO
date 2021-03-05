package webHandler

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"github.com/gabriel-vasile/mimetype"
	"time"
	"strings"
)

func handleStatic (w http.ResponseWriter, r *http.Request){

	startTime := time.Now()
	log.Printf("Handling a static request, reading request path")
	log.Printf("Request made to "+r.URL.Path)

	info, statErr := os.Stat(string(r.URL.Path[1:]))
	if statErr == nil {
		if info.IsDir() {
			log.Printf("Came across a directory in the request")
			log.Printf("403 Forbidden - "+ string(r.URL.Path[1:]))
			w.WriteHeader(http.StatusForbidden)
			fmt.Fprint(w, "403, Forbidden -- Skyfallen Static Server \n")
			elapsed := time.Since(startTime).Seconds()*1000
			fmt.Fprint(w,fmt.Sprintf("Request processed in %v ms", elapsed))

			return
		}
	}


	staticFile, err := os.Open(string(r.URL.Path[1:]))

	if err != nil {

		log.Printf("Error occured while serving static file")
		log.Printf("Err on fileopen: "+ string(r.URL.Path[1:]))
		log.Print(err)
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "404, File not found -- Skyfallen Static Server \n")
		elapsed := time.Since(startTime).Seconds()*1000
		fmt.Fprint(w,fmt.Sprintf("Request processed in %v ms", elapsed))


	} else {

		log.Printf("File found, getting mime type")

		mime, err := mimetype.DetectFile(string(r.URL.Path[1:]))

		if err != nil{

			log.Printf("Error: Cant detect file Mime type, aborting request.")
			w.WriteHeader(http.StatusBadRequest)
			return

		}

		var mimeType string

		mimeType = mime.String()

		b := string(r.URL.Path[len(r.URL.Path)-4:]) == ".css"
		c := string(r.URL.Path[len(r.URL.Path)-3:]) == ".js"
		d := string(r.URL.Path[len(r.URL.Path)-5:]) == ".html"
		if b {

			mimeType = "text/css"
			log.Printf("File extension .css detected, overriding mime type as text/css")

		}
		if c {

			mimeType = "text/javascript"
			log.Printf("File extension .js detected, overriding mime type as text/javascript")

		}
		if d {

			mimeType = "text/html"
			log.Printf("File extension .html detected, overriding mime type as text/html")

		}

		log.Printf("Detected file type, setting header: "+mimeType)

		w.Header().Set("Content-Type",mimeType)

		log.Printf("Sucessful Request, 200 OK - " + r.URL.Path)

		contentsNE, _ := ioutil.ReadAll(staticFile)

		contents := strings.ReplaceAll(string(contentsNE), "%", "%%")

		fmt.Fprintf(w, string(contents))

		elapsed := time.Since(startTime).Seconds()*1000
		log.Println(fmt.Sprintf("Served static file in %v ms", elapsed))

	}

}
