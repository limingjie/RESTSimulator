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

// GatewayClusterProfiles map
var GatewayClusterProfiles = make(map[string]models.GatewayClusterProfile)

func saveGatewayClusterProfiles() {
	WriteFile(filepath.Clean("GatewayData/GatewayClusterProfiles.json"), GatewayClusterProfiles)
}

// PostGatewayClusterProfile - POST /profiles/gatewaycluster
func PostGatewayClusterProfile(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	profile := models.GatewayClusterProfile{}
	json.NewDecoder(r.Body).Decode(&profile)

	msg, _ := json.Marshal(profile)
	log.Println(r.RemoteAddr, "PostGatewayClusterProfile", string(msg))

	profileName := profile.Profile.ProfileName
	if len(profileName) > 0 {
		_, ok := GatewayClusterProfiles[profileName]
		if ok {
			w.WriteHeader(409)
			fmt.Fprintf(w, "Error: Gateway Cluster profile with same name already exists.")
		} else {
			profile.Profile.LastUpdated = time.Now().Format("2006/01/02 15:04:05")
			if len(profile.Profile.AccessPermission) == 0 {
				profile.Profile.AccessPermission = "ReadWrite"
			}
			GatewayClusterProfiles[profileName] = profile

			w.WriteHeader(201)
			fmt.Fprintf(w, "Succeed.")

			saveGatewayClusterProfiles()
		}
	} else {
		w.WriteHeader(400)
		fmt.Fprintf(w, "Error: Expecting Profile Name.")
	}
}

// GetGatewayClusterProfiles - GET /profiles/gatewaycluster
func GetGatewayClusterProfiles(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var profilesJSON bytes.Buffer
	profilesJSON.WriteString("{\"GatewayClusterProfile\":[")

	log.Println(r.RemoteAddr, "GetGatewayClusterProfiles", strconv.Itoa(len(GatewayClusterProfiles)))

	ok := false
	for _, profile := range GatewayClusterProfiles {
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

// GetGatewayClusterProfile - GET /profiles/gatewaycluster/:name
func GetGatewayClusterProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	profile, ok := GatewayClusterProfiles[ps.ByName("name")]

	log.Println(r.RemoteAddr, "GetGatewayClusterProfile", ps.ByName("name"))

	if ok {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		profileJSON, _ := json.Marshal(profile)
		fmt.Fprintf(w, "%s", profileJSON)
	} else {
		w.WriteHeader(404)
		fmt.Fprintf(w, "Error: Gateway Cluster profile does not exist.")
	}
}

// PutGatewayClusterProfile - PUT /profiles/gatewaycluster/:name
func PutGatewayClusterProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	profile := models.GatewayClusterProfile{}
	json.NewDecoder(r.Body).Decode(&profile)

	msg, _ := json.Marshal(profile)
	log.Println(r.RemoteAddr, "PutGatewayClusterProfile", string(msg))

	profileName := ps.ByName("name")
	_, ok := GatewayClusterProfiles[profileName]
	if ok {
		if profile.Profile.ProfileName == profileName {
			profile.Profile.LastUpdated = time.Now().Format("2006/01/02 15:04:05")
			if len(profile.Profile.AccessPermission) == 0 {
				profile.Profile.AccessPermission = "ReadWrite"
			}
			GatewayClusterProfiles[profileName] = profile
			w.WriteHeader(200)
			fmt.Fprintf(w, "Succeed.")

			saveGatewayClusterProfiles()
		} else {
			w.WriteHeader(409)
			fmt.Fprintf(w, "Error: Gateway Cluster profile name does not match.")
		}
	} else {
		w.WriteHeader(404)
		fmt.Fprintf(w, "Error: Gateway Cluster profile does not exist.")
	}
}

// DeleteGatewayClusterProfile - DELETE /profiles/gatewaycluster/:name
func DeleteGatewayClusterProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	profileName := ps.ByName("name")

	log.Println(r.RemoteAddr, "DeleteGatewayClusterProfile", profileName)

	_, ok := GatewayClusterProfiles[profileName]
	if ok {
		delete(GatewayClusterProfiles, profileName)
		w.WriteHeader(200)
		fmt.Fprintf(w, "Succeed.")

		saveGatewayClusterProfiles()
	} else {
		w.WriteHeader(404)
		fmt.Fprintf(w, "Error: Gateway Cluster profile does not exist.")
	}
}
