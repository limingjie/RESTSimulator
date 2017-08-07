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

// ServerProfiles map
var ServerProfiles = make(map[string]models.ServerProfile)

func saveServerProfiles() {
	WriteFile(filepath.Clean("GatewayData/ServerProfiles.json"), ServerProfiles)
}

// PostServerProfile - POST /profiles/servers
func PostServerProfile(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	profile := models.ServerProfile{}
	json.NewDecoder(r.Body).Decode(&profile)

	msg, _ := json.Marshal(profile)
	log.Println(r.RemoteAddr, "PostServerProfile", string(msg))

	profileName := profile.Profile.ProfileName
	if len(profileName) > 0 {
		_, ok := ServerProfiles[profileName]
		if ok {
			w.WriteHeader(409)
			fmt.Fprintf(w, "Error: Server profile with same name already exists.")
		} else {
			profile.Profile.LastUpdated = time.Now().Format("2006/01/02 15:04:05")
			if len(profile.Profile.AccessPermission) == 0 {
				profile.Profile.AccessPermission = "ReadWrite"
			}
			ServerProfiles[profileName] = profile

			w.WriteHeader(201)
			fmt.Fprintf(w, "Succeed.")

			saveServerProfiles()
		}
	} else {
		w.WriteHeader(400)
		fmt.Fprintf(w, "Error: Expecting Profile Name.")
	}
}

// GetServerProfiles - GET /profiles/servers
func GetServerProfiles(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var profilesJSON bytes.Buffer
	profilesJSON.WriteString("{\"ServerProfile\":[")

	log.Println(r.RemoteAddr, "GetServerProfiles", strconv.Itoa(len(ServerProfiles)))

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

// GetServerProfile - GET /profiles/servers/:name
func GetServerProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	profile, ok := ServerProfiles[ps.ByName("name")]

	log.Println(r.RemoteAddr, "GetServerProfile", ps.ByName("name"))

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

// PutServerProfile - PUT /profiles/servers/:name
func PutServerProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	profile := models.ServerProfile{}
	json.NewDecoder(r.Body).Decode(&profile)

	msg, _ := json.Marshal(profile)
	log.Println(r.RemoteAddr, "PutServerProfile", string(msg))

	profileName := ps.ByName("name")
	_, ok := ServerProfiles[profileName]
	if ok {
		if profile.Profile.ProfileName == profileName {
			profile.Profile.LastUpdated = time.Now().Format("2006/01/02 15:04:05")
			if len(profile.Profile.AccessPermission) == 0 {
				profile.Profile.AccessPermission = "ReadWrite"
			}
			ServerProfiles[profileName] = profile
			w.WriteHeader(200)
			fmt.Fprintf(w, "Succeed.")

			saveServerProfiles()
		} else {
			w.WriteHeader(409)
			fmt.Fprintf(w, "Error: Server profile name does not match.")
		}
	} else {
		w.WriteHeader(404)
		fmt.Fprintf(w, "Error: Server profile does not exist.")
	}
}

// DeleteServerProfile - DELETE /profiles/servers/:name
func DeleteServerProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	profileName := ps.ByName("name")

	log.Println(r.RemoteAddr, "DeleteServerProfile", profileName)

	_, ok := ServerProfiles[profileName]
	if ok {
		delete(ServerProfiles, profileName)
		w.WriteHeader(200)
		fmt.Fprintf(w, "Succeed.")

		saveServerProfiles()
	} else {
		w.WriteHeader(404)
		fmt.Fprintf(w, "Error: Server profile does not exist.")
	}
}
