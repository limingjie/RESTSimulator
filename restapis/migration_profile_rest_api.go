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

// MigrationProfiles map
var MigrationProfiles = make(map[string]models.MigrationProfile)

func saveMigrationProfiles() {
	WriteFile(filepath.Clean("GatewayData/MigrationProfiles.json"), MigrationProfiles)
}

// PostMigrationProfile - POST /profiles/migrations
func PostMigrationProfile(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	profile := models.MigrationProfile{}
	json.NewDecoder(r.Body).Decode(&profile)

	msg, _ := json.Marshal(profile)
	log.Println(r.RemoteAddr, "PostMigrationProfile", string(msg))

	profileName := profile.Profile.ProfileName
	if len(profileName) > 0 {
		_, ok := MigrationProfiles[profileName]
		if ok {
			w.WriteHeader(409)
			fmt.Fprintf(w, "Error: Migration profile with same name already exists.")
		} else {
			profile.Profile.LastUpdated = time.Now().Format("2006/01/02 15:04:05")
			if len(profile.Profile.AccessPermission) == 0 {
				profile.Profile.AccessPermission = "ReadWrite"
			}
			MigrationProfiles[profileName] = profile

			w.WriteHeader(201)
			fmt.Fprintf(w, "Succeed.")

			saveMigrationProfiles()
		}
	} else {
		w.WriteHeader(400)
		fmt.Fprintf(w, "Error: Expecting Profile Name.")
	}
}

// GetMigrationProfiles - GET /profiles/migrations
func GetMigrationProfiles(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var profilesJSON bytes.Buffer
	profilesJSON.WriteString("{\"MigrationProfile\":[")

	log.Println(r.RemoteAddr, "GetMigrationProfiles", strconv.Itoa(len(MigrationProfiles)))

	ok := false
	for _, profile := range MigrationProfiles {
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

// GetMigrationProfile - GET /profiles/migrations/:name
func GetMigrationProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	profile, ok := MigrationProfiles[ps.ByName("name")]

	log.Println(r.RemoteAddr, "GetMigrationProfiles", ps.ByName("name"))

	if ok {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		profileJSON, _ := json.Marshal(profile)
		fmt.Fprintf(w, "%s", profileJSON)
	} else {
		w.WriteHeader(404)
		fmt.Fprintf(w, "Error: Migration profile does not exist.")
	}
}

// PutMigrationProfile - PUT /profiles/migrations/:name
func PutMigrationProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	profile := models.MigrationProfile{}
	json.NewDecoder(r.Body).Decode(&profile)

	msg, _ := json.Marshal(profile)
	log.Println(r.RemoteAddr, "PutMigrationProfile", string(msg))

	profileName := ps.ByName("name")
	_, ok := MigrationProfiles[profileName]
	if ok {
		if profile.Profile.ProfileName == profileName {
			profile.Profile.LastUpdated = time.Now().Format("2006/01/02 15:04:05")
			if len(profile.Profile.AccessPermission) == 0 {
				profile.Profile.AccessPermission = "ReadWrite"
			}
			MigrationProfiles[profileName] = profile
			w.WriteHeader(200)
			fmt.Fprintf(w, "Succeed.")

			saveMigrationProfiles()
		} else {
			w.WriteHeader(409)
			fmt.Fprintf(w, "Error: Migration profile name does not match.")
		}
	} else {
		w.WriteHeader(404)
		fmt.Fprintf(w, "Error: Migration profile does not exist.")
	}
}

// DeleteMigrationProfile - DELETE /profiles/migrations/:name
func DeleteMigrationProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	profileName := ps.ByName("name")

	log.Println(r.RemoteAddr, "DeleteMigrationProfile", profileName)

	_, ok := MigrationProfiles[profileName]
	if ok {
		delete(MigrationProfiles, profileName)
		w.WriteHeader(200)
		fmt.Fprintf(w, "Succeed.")

		saveMigrationProfiles()
	} else {
		w.WriteHeader(404)
		fmt.Fprintf(w, "Error: Migration profile does not exist.")
	}
}
