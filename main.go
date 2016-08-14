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
}

func main() {

	r := mux.NewRouter()
	// TODO: Check Auth & filter by content-type

	r.HandleFunc("/login", LoginHandler).Methods("POST")
	r.HandleFunc("/logout", LogoutHandler).Methods("GET")

	r.HandleFunc("/files/:id", DownloadFileHandler).Methods("GET")
	r.HandleFunc("/files/:id", UpdateFileHandler).Methods("PUT")
	r.HandleFunc("/files/:id", DeleteFileHandler).Methods("DELETE")

	r.HandleFunc("/files", ListFilesHandler).Methods("GET")
	r.HandleFunc("/files", AddFileHandler).Methods("POST")

	// Shows a simple prompt for the user/password and file.
	// Serve static files
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./public/")))

	log.Fatal(http.ListenAndServe(Conf["addr"], r))
}

// Given an string id searches for a file and delivers it to the client
func DownloadFileHandler(w http.ResponseWriter, r *http.Request) {
}

// Updates the file name and expire date
func UpdateFileHandler(w http.ResponseWriter, r *http.Request) {
}

// Deletes a file from the server
func DeleteFileHandler(w http.ResponseWriter, r *http.Request) {
}

// Lists all the files that the server has available to download
func ListFilesHandler(w http.ResponseWriter, r *http.Request) {
}

// Uploads a file to the server, returns the status and expire date
func AddFileHandler(w http.ResponseWriter, r *http.Request) {
}
