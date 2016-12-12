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

var dataEnterprises = make(map[string]models.Enterprise)

// Index - For /
func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}

// GetEnterprise - GET /profiles/enterprises/:name
func GetEnterprise(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ent, ok := dataEnterprises[ps.ByName("name")]
	if ok {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		entj, _ := json.Marshal(ent)
		fmt.Fprintf(w, "%s", entj)
	} else {
		w.WriteHeader(404)
		fmt.Fprintf(w, "Error: Not Found.")
	}
}

// PostEnterprise - POST /profiles/enterprises
func PostEnterprise(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	ent := models.Enterprise{}

	json.NewDecoder(r.Body).Decode(&ent)

	if len(ent.EnterpriseProfile.ProfileName) > 0 {
		_, ok := dataEnterprises[ent.EnterpriseProfile.ProfileName]
		if ok {
			w.WriteHeader(409)
			fmt.Fprintf(w, "Error: Duplicate Profile.")
		} else {
			ent.EnterpriseProfile.LastUpdated = time.Now().Format("2006/01/02 15:04:05")
			dataEnterprises[ent.EnterpriseProfile.ProfileName] = ent

			w.WriteHeader(201)
			fmt.Fprintf(w, "Succeed.")
		}
	} else {
		w.WriteHeader(400)
		fmt.Fprintf(w, "Error: Expecting Profile Name.")
	}
}

func main() {
	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/profiles/enterprises/:name", GetEnterprise)
	router.POST("/profiles/enterprises", PostEnterprise)

	log.Fatal(http.ListenAndServe(":8888", router))
}
