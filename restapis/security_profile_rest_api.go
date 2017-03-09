package restapis

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"../models"
	"github.com/julienschmidt/httprouter"
)

// SecurityProfiles map
var SecurityProfiles = make(map[string]models.SecurityProfile)

// PostSecurityProfile - POST /profiles/security
func PostSecurityProfile(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	profile := models.SecurityProfile{}
	json.NewDecoder(r.Body).Decode(&profile)

	msg, _ := json.Marshal(profile)
	log.Println("PostSecurityProfile", string(msg))

	profileName := profile.Profile.ProfileName
	if len(profileName) > 0 {
		_, ok := SecurityProfiles[profileName]
		if ok {
			w.WriteHeader(409)
			fmt.Fprintf(w, "Error: Security profile with same name already exists.")
		} else {
			profile.Profile.LastUpdated = time.Now().Format("2006/01/02 15:04:05")
			if len(profile.Profile.AccessPermission) == 0 {
				profile.Profile.AccessPermission = "ReadWrite"
			}
			SecurityProfiles[profileName] = profile

			w.WriteHeader(201)
			fmt.Fprintf(w, "Succeed.")
		}
	} else {
		w.WriteHeader(400)
		fmt.Fprintf(w, "Error: Expecting Profile Name.")
	}
}

// GetSecurityProfiles - GET /profiles/security
func GetSecurityProfiles(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var profilesJSON bytes.Buffer
	profilesJSON.WriteString("{\"SecurityProfiles\":[")

	log.Println("GetSecurityProfiles", strconv.Itoa(len(SecurityProfiles)))

	ok := false
	for _, profile := range SecurityProfiles {
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

// GetSecurityProfile - GET /profiles/security/:name
func GetSecurityProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	profile, ok := SecurityProfiles[ps.ByName("name")]

	log.Println("GetSecurityProfile", ps.ByName("name"))

	if ok {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		profileJSON, _ := json.Marshal(profile)
		fmt.Fprintf(w, "%s", profileJSON)
	} else {
		w.WriteHeader(404)
		fmt.Fprintf(w, "Error: Security profile does not exist.")
	}
}

// PutSecurityProfile - PUT /profiles/security/:name
func PutSecurityProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	profile := models.SecurityProfile{}
	json.NewDecoder(r.Body).Decode(&profile)

	msg, _ := json.Marshal(profile)
	log.Println("PutSecurityProfile", string(msg))

	profileName := ps.ByName("name")
	_, ok := SecurityProfiles[profileName]
	if ok {
		if profile.Profile.ProfileName == profileName {
			profile.Profile.LastUpdated = time.Now().Format("2006/01/02 15:04:05")
			if len(profile.Profile.AccessPermission) == 0 {
				profile.Profile.AccessPermission = "ReadWrite"
			}
			SecurityProfiles[profileName] = profile
			w.WriteHeader(200)
			fmt.Fprintf(w, "Succeed.")
		} else {
			w.WriteHeader(409)
			fmt.Fprintf(w, "Error: Security profile name does not match.")
		}
	} else {
		w.WriteHeader(404)
		fmt.Fprintf(w, "Error: Security profile does not exist.")
	}
}

// DeleteSecurityProfile - DELETE /profiles/security/:name
func DeleteSecurityProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	profileName := ps.ByName("name")

	log.Println("DeleteSecurityProfile", profileName)

	_, ok := SecurityProfiles[profileName]
	if ok {
		delete(SecurityProfiles, profileName)
		w.WriteHeader(200)
		fmt.Fprintf(w, "Succeed.")
	} else {
		w.WriteHeader(404)
		fmt.Fprintf(w, "Error: Security profile does not exist.")
	}
}
