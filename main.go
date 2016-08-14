// Serve files identified only by a string
package main

import (
	"log"
	"net/http"

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

	r := mux.NewRouter()
	// TODO: Check Auth & filter by content-type

	r.HandleFunc("/login", LoginHandler).Methods("POST")
	r.HandleFunc("/logout", LogoutHandler).Methods("GET")

	r.HandleFunc("/archives/:id", DownloadArchiveHandler).Methods("GET")
	r.HandleFunc("/archives/:id", UpdateArchiveHandler).Methods("PUT")
	r.HandleFunc("/archives/:id", DeleteArchiveHandler).Methods("DELETE")

	r.HandleFunc("/archives", ListArchiveHandler).Methods("GET")
	r.HandleFunc("/archives", AddArchiveHandler).Methods("POST")

	// Shows a simple prompt for the user/password and file.
	// Serve static files
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./public/")))

	log.Fatal(http.ListenAndServe(Conf["addr"], r))
}
