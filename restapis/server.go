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

// ServerProfile map
var ServerProfiles = make(map[string]models.ServerProfile)

// PostServer - POST /profiles/servers
func PostServer(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	profile := models.ServerProfile{}
	json.NewDecoder(r.Body).Decode(&profile)

	profileName := profile.Profile.ProfileName
	if len(profileName) > 0 {
		_, ok := ServerProfiles[profileName]
		if ok {
			w.WriteHeader(409)
			fmt.Fprintf(w, "Error: Server profile with same name already exists.")
		} else {
			profile.Profile.LastUpdated = time.Now().Format("2006/01/02 15:04:05")
			ServerProfiles[profileName] = profile

			w.WriteHeader(201)
			fmt.Fprintf(w, "Succeed.")
		}
	} else {
		w.WriteHeader(400)
		fmt.Fprintf(w, "Error: Expecting Profile Name.")
	}
}

// GetServers - GET /profiles/servers
func GetServers(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var profilesJSON bytes.Buffer
	profilesJSON.WriteString("{\"serverProfile\":[")

	ok := false
	for _, profile := range ServerProfiles {
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

// GetServer - GET /profiles/servers/:profilename
func GetServer(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	profile, ok := ServerProfiles[ps.ByName("profilename")]
	if ok {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		profileJSON, _ := json.Marshal(profile)
		fmt.Fprintf(w, "%s", profileJSON)
	} else {
		w.WriteHeader(404)
		fmt.Fprintf(w, "Error: Server profile does not exist.")
	}
}

// PutServer - PUT /profiles/servers/:profilename
func PutServer(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	profile := models.ServerProfile{}
	json.NewDecoder(r.Body).Decode(&profile)

	profileName := ps.ByName("profilename")
	_, ok := ServerProfiles[profileName]
	if ok {
		if profile.Profile.ProfileName == profileName {
			profile.Profile.LastUpdated = time.Now().Format("2006/01/02 15:04:05")
			ServerProfiles[profileName] = profile
			w.WriteHeader(200)
			fmt.Fprintf(w, "Succeed.")
		} else {
			w.WriteHeader(409)
			fmt.Fprintf(w, "Error: Server profile name does not match.")
		}
	} else {
		w.WriteHeader(404)
		fmt.Fprintf(w, "Error: Server profile does not exist.")
	}
}

// DeleteServer - DELETE /profiles/servers/:profilename
func DeleteServer(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	profileName := ps.ByName("profilename")
	_, ok := ServerProfiles[profileName]
	if ok {
		delete(ServerProfiles, profileName)
		w.WriteHeader(200)
		fmt.Fprintf(w, "Succeed.")
	} else {
		w.WriteHeader(404)
		fmt.Fprintf(w, "Error: Server profile does not exist.")
	}
}