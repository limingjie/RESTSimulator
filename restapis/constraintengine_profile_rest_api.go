package restapis

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"strconv"
	"time"

	"../models"
	"github.com/julienschmidt/httprouter"
)

// ConstraintEngineProfiles map
var ConstraintEngineProfiles = make(map[string]models.ConstraintEngineProfile)

func saveConstraintEngineProfiles() {
	WriteFile(filepath.Clean("GatewayData/ConstraintEngineProfiles.json"), ConstraintEngineProfiles)
}

// PostConstraintEngineProfile - POST /profiles/constraintengine
func PostConstraintEngineProfile(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	profile := models.ConstraintEngineProfile{}
	json.NewDecoder(r.Body).Decode(&profile)

	msg, _ := json.Marshal(profile)
	log.Println(r.RemoteAddr, "PostConstraintEngineProfile", string(msg))

	profileName := profile.Profile.ProfileName
	if len(profileName) > 0 {
		_, ok := ConstraintEngineProfiles[profileName]
		if ok {
			w.WriteHeader(409)
			fmt.Fprintf(w, "Error: Constraint Engine profile with same name already exists.")
		} else {
			profile.Profile.LastUpdated = time.Now().Format("2006/01/02 15:04:05")
			if len(profile.Profile.AccessPermission) == 0 {
				profile.Profile.AccessPermission = "ReadWrite"
			}
			ConstraintEngineProfiles[profileName] = profile

			w.WriteHeader(201)
			fmt.Fprintf(w, "Succeed.")

			saveConstraintEngineProfiles()
		}
	} else {
		w.WriteHeader(400)
		fmt.Fprintf(w, "Error: Expecting Profile Name.")
	}
}

// GetConstraintEngineProfiles - GET /profiles/constraintengine
func GetConstraintEngineProfiles(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var profilesJSON bytes.Buffer
	profilesJSON.WriteString("{\"ConstraintEngineProfile\":[")

	log.Println(r.RemoteAddr, "GetConstraintEngineProfiles", strconv.Itoa(len(ConstraintEngineProfiles)))

	ok := false
	for _, profile := range ConstraintEngineProfiles {
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

// GetConstraintEngineProfile - GET /profiles/constraintengine/:name
func GetConstraintEngineProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	profile, ok := ConstraintEngineProfiles[ps.ByName("name")]

	log.Println(r.RemoteAddr, "GetConstraintEngineProfile", ps.ByName("name"))

	if ok {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		profileJSON, _ := json.Marshal(profile)
		fmt.Fprintf(w, "%s", profileJSON)
	} else {
		w.WriteHeader(404)
		fmt.Fprintf(w, "Error: Constraint Engine profile does not exist.")
	}
}

// PutConstraintEngineProfile - PUT /profiles/constraintengine/:name
func PutConstraintEngineProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	profile := models.ConstraintEngineProfile{}
	json.NewDecoder(r.Body).Decode(&profile)

	msg, _ := json.Marshal(profile)
	log.Println(r.RemoteAddr, "PutConstraintEngineProfile", string(msg))

	profileName := ps.ByName("name")
	_, ok := ConstraintEngineProfiles[profileName]
	if ok {
		if profile.Profile.ProfileName == profileName {
			profile.Profile.LastUpdated = time.Now().Format("2006/01/02 15:04:05")
			if len(profile.Profile.AccessPermission) == 0 {
				profile.Profile.AccessPermission = "ReadWrite"
			}
			ConstraintEngineProfiles[profileName] = profile
			w.WriteHeader(200)
			fmt.Fprintf(w, "Succeed.")

			saveConstraintEngineProfiles()
		} else {
			w.WriteHeader(409)
			fmt.Fprintf(w, "Error: Constraint Engine profile name does not match.")
		}
	} else {
		w.WriteHeader(404)
		fmt.Fprintf(w, "Error: Constraint Engine profile does not exist.")
	}
}

// DeleteConstraintEngineProfile - DELETE /profiles/constraintengine/:name
func DeleteConstraintEngineProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	profileName := ps.ByName("name")

	log.Println(r.RemoteAddr, "DeleteConstraintEngineProfile", profileName)

	_, ok := ConstraintEngineProfiles[profileName]
	if ok {
		delete(ConstraintEngineProfiles, profileName)
		w.WriteHeader(200)
		fmt.Fprintf(w, "Succeed.")

		saveConstraintEngineProfiles()
	} else {
		w.WriteHeader(404)
		fmt.Fprintf(w, "Error: Constraint Engine profile does not exist.")
	}
}
