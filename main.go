// Serve files identified only by a string
package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

var Conf map[string]string

func init() {
	// TODO: Read config file
	Conf = map[string]string{
		"addr":      "127.0.0.1:8080",
		"secretKey": "secretsecret",
	}

	ArchiveStore = make(map[string]Archive)
}

func main() {
	readConf()

	r := mux.NewRouter()
	// TODO: Check Auth & filter by content-type

	r.HandleFunc("/login", LoginHandler).Methods("POST")
	r.HandleFunc("/logout", LogoutHandler).Methods("GET")

	r.HandleFunc("/archives/{archiveKey}", DownloadArchiveHandler).Methods("GET")
	r.HandleFunc("/archives/{archiveKey}", UpdateArchiveHandler).Methods("PUT")
	r.HandleFunc("/archives/{archiveKey}", DeleteArchiveHandler).Methods("DELETE")

	r.HandleFunc("/archives", ListArchiveHandler).Methods("GET")
	r.HandleFunc("/archives", AddArchiveHandler).Methods("POST")

	// Shows a simple prompt for the user/password and file.
	// Serve static files
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./public/")))

	log.Fatal(http.ListenAndServe(Conf["addr"], r))
}

func readConf() {
	var fileConf map[string]string

	fh, err := os.Open("/etc/fileshare.json")
	if err != nil {
		log.Println("Couldn't open the config file")
		return
	}

	defer fh.Close()

	dec := json.NewDecoder(fh)
	err = dec.Decode(&fileConf)
	if err != nil {
		log.Println("Error reading the config file")
		return
	}

	for k := range Conf {
		if _, ok := fileConf[k]; ok {
			Conf[k] = fileConf[k]
		}
	}
}
