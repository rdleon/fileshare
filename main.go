// Serve files identified only by a string
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
)

type User struct {
	Name     string
	Password string
}

var (
	user User
	Conf map[string]string
)

func init() {
	user = User{Name: "rdleon", Password: "password"}
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

// Validate the webtoken and return the username
func ValidateJWT(tokenStr string) (string, bool) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method %v", token.Header["alg"])
		}

		return []byte(Conf["secretKey"]), nil
	})

	if err != nil {
		return "", false
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims["sub"].(string), true
	}

	return "", false
}

// Authenticates the user using name:password and generates a JWT
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var (
		credentials User
		response    map[string]interface{}
	)

	if content := r.Header.Get("Content-Type"); content == "application/json" {

		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&credentials)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "{\"error\": \"Bad Request\"}")
			return
		}

		response = make(map[string]interface{})

		if credentials == user {
			claims := jwt.StandardClaims{
				ExpiresAt: 15000,
				Issuer:    "fileshare",
				Subject:   user.Name,
			}
			token := jwt.NewWithClaims(jwt.SigningMethodHS256, &claims)

			response["token"], err = token.SignedString([]byte(Conf["secretKey"]))

			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				log.Println("Error: ", err)
				fmt.Fprintf(w, "{\"error\": \"Internal Server Error\"}")
				return
			}
		} else {
			response["error"] = "Wrong username or password"
		}

		json.NewEncoder(w).Encode(response)
	} else {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "{\"error\": \"Must send content as application/json\"}")
	}
}

// Logout the user by "destroying" the auth token before it expires
func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: destroy JWT token, expire session
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
