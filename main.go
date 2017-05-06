// Serve files identified only by a string
package main

import (
	"encoding/json"
	"errors"
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

var Conf map[string]string

func init() {
	Conf = map[string]string{
		"listen":    "*:8080",
		"user":      "admin",
		"password":  "",
		"saveDir":   "/tmp",
		"secretKey": "",
	}

	ArchiveStore = make(map[string]Archive)
}

func main() {
	err := readConf()
	if err != nil {
		log.Fatal(err)
	}

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

	log.Println("listening on", Conf["listen"])
	log.Fatal(http.ListenAndServe(Conf["listen"], r))
}

func readConf() error {
	var fileConf map[string]string

	var conf = flag.String("conf", "./config.json", "The configuration file")
	flag.Parse()

	fh, err := os.Open(*conf)
	if err != nil {
		return err
	}
	defer fh.Close()

	dec := json.NewDecoder(fh)
	err = dec.Decode(&fileConf)
	if err != nil {
		return err
	}

	for k := range Conf {
		if _, ok := fileConf[k]; ok {
			Conf[k] = fileConf[k]
		}
	}

	if user, ok := Conf["user"]; ok {
		MyUser.Name = user
	}

	pass, ok := Conf["password"]

	if !ok || len(pass) < 3 {
		return errors.New("Password missing or too short")
	}

	MyUser.Password = pass

	return nil
}
