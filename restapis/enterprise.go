package restapis

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"../models"
	"github.com/julienschmidt/httprouter"
)

// EnterpriseProfiles map
var EnterpriseProfiles = make(map[string]models.EnterpriseProfile)

// PostEnterprise - POST /profiles/enterprises
func PostEnterprise(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	profile := models.EnterpriseProfile{}
	json.NewDecoder(r.Body).Decode(&profile)

	profileName := profile.Profile.ProfileName
	if len(profileName) > 0 {
		_, ok := EnterpriseProfiles[profileName]
		if ok {
			w.WriteHeader(409)
			fmt.Fprintf(w, "Error: Enterprise profile with same name already exists.")
		} else {
			profile.Profile.LastUpdated = time.Now().Format("2006/01/02 15:04:05")
			EnterpriseProfiles[profileName] = profile

			w.WriteHeader(201)
			fmt.Fprintf(w, "Succeed.")
		}
	} else {
		w.WriteHeader(400)
		fmt.Fprintf(w, "Error: Expecting Profile Name.")
	}
}

// GetEnterprises - GET /profiles/enterprises
func GetEnterprises(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var profilesJSON bytes.Buffer
	profilesJSON.WriteString("{\"enterpriseProfile\":[")

	ok := false
	for _, profile := range EnterpriseProfiles {
		if ok {
			profilesJSON.WriteString(",")
		} else {
			ok = true
		}
		profileJSON, _ := json.Marshal(profile)
		profilesJSON.Write(profileJSON)
	}

	profilesJSON.WriteString("]}")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s", profilesJSON.String())
}

// GetEnterprise - GET /profiles/enterprises/:profilename
func GetEnterprise(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	profile, ok := EnterpriseProfiles[ps.ByName("profilename")]
	if ok {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		profileJSON, _ := json.Marshal(profile)
		fmt.Fprintf(w, "%s", profileJSON)
	} else {
		w.WriteHeader(404)
		fmt.Fprintf(w, "Error: Enterprise profile does not exist.")
	}
}

// PutEnterprise - PUT /profiles/enterprises/:profilename
func PutEnterprise(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	profile := models.EnterpriseProfile{}
	json.NewDecoder(r.Body).Decode(&profile)

	profileName := ps.ByName("profilename")
	_, ok := EnterpriseProfiles[profileName]
	if ok {
		if profile.Profile.ProfileName == profileName {
			profile.Profile.LastUpdated = time.Now().Format("2006/01/02 15:04:05")
			EnterpriseProfiles[profileName] = profile
			w.WriteHeader(200)
			fmt.Fprintf(w, "Succeed.")
		} else {
			w.WriteHeader(409)
			fmt.Fprintf(w, "Error: Enterprise profile name does not match.")
		}
	} else {
		w.WriteHeader(404)
		fmt.Fprintf(w, "Error: Enterprise profile does not exist.")
	}
}

// DeleteEnterprise - DELETE /profiles/enterprises/:profilename
func DeleteEnterprise(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	profileName := ps.ByName("profilename")
	_, ok := EnterpriseProfiles[profileName]
	if ok {
		delete(EnterpriseProfiles, profileName)
		w.WriteHeader(200)
		fmt.Fprintf(w, "Succeed.")
	} else {
		w.WriteHeader(404)
		fmt.Fprintf(w, "Error: Enterprise profile does not exist.")
	}
}
