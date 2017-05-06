package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type User struct {
	Name     string
	Password string
}

var MyUser = User{Name: "admin", Password: ""}
var loggedIn = make(map[string]bool)

// Validate the webtoken and return the username
func CheckAuth(r *http.Request) (string, bool) {
	tokenStr := r.Header.Get("Authorization")

	if tokenStr == "" {
		return "", false
	}

	i := strings.Index(tokenStr, "Bearer ")

	if i >= 0 {
		tokenStr = tokenStr[i+len("Bearer "):]
	}

	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method %v", token.Header["alg"])
		}

		return []byte(Conf["secretKey"]), nil
	})

	if err != nil {
		return "", false
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return "", false
	}

	name := claims["sub"].(string)
	if !loggedIn[name] {
		return "", false
	}

	return name, true
}

// Authenticates the user using name:password and generates a JWT
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var (
		credentials User
		response    map[string]interface{}
	)

	if _, ok := CheckAuth(r); ok {
		fmt.Fprintf(w, "{\"loggedin\": true}")
		return
	}

	if content := r.Header.Get("Content-Type"); content == "application/json" {

		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&credentials)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "{\"error\": \"Bad Request\"}")
			return
		}

		response = make(map[string]interface{})

		if credentials == MyUser {
			loggedIn[MyUser.Name] = true
			claims := jwt.StandardClaims{
				ExpiresAt: time.Now().Add(time.Hour * 2).Unix(),
				Issuer:    "fileshare",
				Subject:   MyUser.Name,
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
	if username, ok := CheckAuth(r); ok {
		delete(loggedIn, username)
	}

	fmt.Fprintf(w, "{\"loggedin\": false}")
}
