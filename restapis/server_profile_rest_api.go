package restapis

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"../logger"
	"../models"
	"github.com/julienschmidt/httprouter"
)

// ServerProfiles map
var ServerProfiles = make(map[string]models.ServerProfile)

// PostServerProfile - POST /profiles/servers
func PostServerProfile(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	profile := models.ServerProfile{}
	json.NewDecoder(r.Body).Decode(&profile)

	msg, _ := json.Marshal(profile)
	logger.Logger("PostServerProfile", string(msg))

	profileName := profile.Profile.ProfileName
	if len(profileName) > 0 {
		_, ok := ServerProfiles[profileName]
		if ok {
			w.WriteHeader(409)
			fmt.Fprintf(w, "Error: Server profile with same name already exists.")
		} else {
			profile.Profile.LastUpdated = time.Now().Format("2006/01/02 15:04:05")
			if len(profile.Profile.AccessPermission) == 0 {
				profile.Profile.AccessPermission = "READWRITE"
			}
			ServerProfiles[profileName] = profile

			w.WriteHeader(201)
			fmt.Fprintf(w, "Succeed.")
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

	logger.Logger("GetServerProfiles", strconv.Itoa(len(ServerProfiles)))

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

// GetServerProfile - GET /profiles/servers/:profilename
func GetServerProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	profile, ok := ServerProfiles[ps.ByName("profilename")]

	logger.Logger("GetServerProfile", ps.ByName("profilename"))

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

// PutServerProfile - PUT /profiles/servers/:profilename
func PutServerProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	profile := models.ServerProfile{}
	json.NewDecoder(r.Body).Decode(&profile)

	msg, _ := json.Marshal(profile)
	logger.Logger("PutServerProfile", string(msg))

	profileName := ps.ByName("profilename")
	_, ok := ServerProfiles[profileName]
	if ok {
		if profile.Profile.ProfileName == profileName {
			profile.Profile.LastUpdated = time.Now().Format("2006/01/02 15:04:05")
			if len(profile.Profile.AccessPermission) == 0 {
				profile.Profile.AccessPermission = "READWRITE"
			}
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

// DeleteServerProfile - DELETE /profiles/servers/:profilename
func DeleteServerProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	profileName := ps.ByName("profilename")

	logger.Logger("DeleteServerProfile", profileName)

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
