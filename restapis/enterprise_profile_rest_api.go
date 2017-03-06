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

// EnterpriseProfiles map
var EnterpriseProfiles = make(map[string]models.EnterpriseProfile)

// PostEnterpriseProfile - POST /profiles/enterprises
func PostEnterpriseProfile(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	profile := models.EnterpriseProfile{}
	json.NewDecoder(r.Body).Decode(&profile)

	msg, _ := json.Marshal(profile)
	logger.Logger("PostEnterpriseProfile", string(msg))

	profileName := profile.Profile.ProfileName
	if len(profileName) > 0 {
		_, ok := EnterpriseProfiles[profileName]
		if ok {
			w.WriteHeader(409)
			fmt.Fprintf(w, "Error: Enterprise profile with same name already exists.")
		} else {
			profile.Profile.LastUpdated = time.Now().Format("2006/01/02 15:04:05")
			if len(profile.Profile.AccessPermission) == 0 {
				profile.Profile.AccessPermission = "ReadWrite"
			}
			EnterpriseProfiles[profileName] = profile

			w.WriteHeader(201)
			fmt.Fprintf(w, "Succeed.")
		}
	} else {
		w.WriteHeader(400)
		fmt.Fprintf(w, "Error: Expecting Profile Name.")
	}
}

// GetEnterpriseProfiles - GET /profiles/enterprises
func GetEnterpriseProfiles(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var profilesJSON bytes.Buffer
	profilesJSON.WriteString("{\"EnterpriseProfile\":[")

	logger.Logger("GetEnterpriseProfiles", strconv.Itoa(len(EnterpriseProfiles)))

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

// GetEnterpriseProfile - GET /profiles/enterprises/:profilename
func GetEnterpriseProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	profile, ok := EnterpriseProfiles[ps.ByName("profilename")]

	logger.Logger("GetEnterpriseProfile", ps.ByName("profilename"))

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

// PutEnterpriseProfile - PUT /profiles/enterprises/:profilename
func PutEnterpriseProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	profile := models.EnterpriseProfile{}
	json.NewDecoder(r.Body).Decode(&profile)

	msg, _ := json.Marshal(profile)
	logger.Logger("PutEnterpriseProfile", string(msg))

	profileName := ps.ByName("profilename")
	_, ok := EnterpriseProfiles[profileName]
	if ok {
		if profile.Profile.ProfileName == profileName {
			profile.Profile.LastUpdated = time.Now().Format("2006/01/02 15:04:05")
			if len(profile.Profile.AccessPermission) == 0 {
				profile.Profile.AccessPermission = "ReadWrite"
			}
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

// DeleteEnterpriseProfile - DELETE /profiles/enterprises/:profilename
func DeleteEnterpriseProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	profileName := ps.ByName("profilename")

	logger.Logger("DeleteEnterpriseProfile", profileName)

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
