package main

import "net/http"

type Archive struct {
	SavePath string
	Name     string
	Key      string
	Expire   int
}

var ArchiveStore map[string]Archive

// Given an string id searches for a file and delivers it to the client
func DownloadArchiveHandler(w http.ResponseWriter, r *http.Request) {
}

// Updates the file name and expire date
func UpdateArchiveHandler(w http.ResponseWriter, r *http.Request) {
}

// Deletes a file from the server
func DeleteArchiveHandler(w http.ResponseWriter, r *http.Request) {
}

// Lists all the files that the server has available to download
func ListArchiveHandler(w http.ResponseWriter, r *http.Request) {
}

// Uploads a file to the server, returns the status and expire date
func AddArchiveHandler(w http.ResponseWriter, r *http.Request) {
}
