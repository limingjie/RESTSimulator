package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"./models"
	"github.com/julienschmidt/httprouter"
)

// Index - For /
func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}

// Users - For /users/:name
func Users(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	u := models.User{
		Id:     ps.ByName("name"),
		Name:   "foobar",
		Gender: "male",
		Age:    30,
	}

	uj, _ := json.Marshal(u)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s", uj)
}

// Enterprises - for /profiles/enterprises/:name
func Enterprises(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	u := models.Enterprises{
		EnterprisesProfile: models.EnterprisesProfile{
			ProfileName:      ps.ByName("name"),
			AccessPermission: "READWRITE",
			LastUpdated:      time.Now().Format("2006/01/02 03:04:05"),
		},
		EnterpriseConfigParams: models.EnterpriseConfigParams{
			DatabasePlatform: "MSSQL_SERVER",
			ConnectString:    "slc02jzc",
			DBUsername:       "SADMIN",
			Encrypt:          "RSA",
			TableOwner:       "db",
		},
	}

	uj, _ := json.Marshal(u)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s", uj)
}

func main() {
	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/users/:name", Users)
	router.GET("/profiles/enterprises/:name", Enterprises)

	log.Fatal(http.ListenAndServe(":8888", router))
}
