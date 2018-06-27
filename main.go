package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	dat, err := ioutil.ReadFile("/etc/adgest.conf")

	var adgest, nom string

	if err == nil {
		adgest = string(dat)
	} else {
		adgest = err.Error()
	}

	name, err := os.Hostname()

	if err == nil {
		nom = name
	} else {
		nom = err.Error()
	}

	fmt.Fprintf(w, "ADPARTS TEST WEBSERVER - CURS DE DOCKER - JUNY 2018\n\n"+
		"URI: %s\n"+
		"HOSTNAME LOCAL: %s\n"+
		"VARIABLE ENTORN $MISSATGE: %s\n"+
		"CONTINGUT FITXER /etc/adgest.conf: \n" +
		"%s\n"+
		"", r.URL.Path[1:], nom, os.Getenv("MISSATGE"), adgest)
}

func main() {
	log.Print("Inicialitzant el webserver...")
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
