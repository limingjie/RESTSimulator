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

// SecurityProfile map
var SecurityProfile = make(map[string]models.SecurityProfile)

// PostSecurityProfile - POST /profiles/security
func PostSecurityProfile(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	profile := models.SecurityProfile{}
	json.NewDecoder(r.Body).Decode(&profile)

	msg, _ := json.Marshal(profile)
	logger.Logger("PostSecurityProfile", string(msg))

	profileName := profile.Profile.ProfileName
	if len(profileName) > 0 {
		_, ok := SecurityProfile[profileName]
		if ok {
			w.WriteHeader(409)
			fmt.Fprintf(w, "Error: Enterprise profile with same name already exists.")
		} else {
			profile.Profile.LastUpdated = time.Now().Format("2006/01/02 15:04:05")
			if len(profile.Profile.AccessPermission) == 0 {
				profile.Profile.AccessPermission = "READWRITE"
			}
			SecurityProfile[profileName] = profile

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
	profilesJSON.WriteString("{\"SecurityProfile\":[")

	logger.Logger("GetSecurityProfiles", strconv.Itoa(len(SecurityProfile)))

	ok := false
	for _, profile := range SecurityProfile {
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

// GetSecurityProfile - GET /profiles/security/:profilename
func GetSecurityProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	profile, ok := SecurityProfile[ps.ByName("profilename")]

	logger.Logger("GetSecurityProfile", ps.ByName("profilename"))

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

// PutSecurityProfile - PUT /profiles/security/:profilename
func PutSecurityProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	profile := models.SecurityProfile{}
	json.NewDecoder(r.Body).Decode(&profile)

	msg, _ := json.Marshal(profile)
	logger.Logger("PutSecurityProfile", string(msg))

	profileName := ps.ByName("profilename")
	_, ok := SecurityProfile[profileName]
	if ok {
		if profile.Profile.ProfileName == profileName {
			profile.Profile.LastUpdated = time.Now().Format("2006/01/02 15:04:05")
			if len(profile.Profile.AccessPermission) == 0 {
				profile.Profile.AccessPermission = "READWRITE"
			}
			SecurityProfile[profileName] = profile
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

// DeleteSecurityProfile - DELETE /profiles/security/:profilename
func DeleteSecurityProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	profileName := ps.ByName("profilename")

	logger.Logger("DeleteSecurityProfile", profileName)

	_, ok := SecurityProfile[profileName]
	if ok {
		delete(SecurityProfile, profileName)
		w.WriteHeader(200)
		fmt.Fprintf(w, "Succeed.")
	} else {
		w.WriteHeader(404)
		fmt.Fprintf(w, "Error: Enterprise profile does not exist.")
	}
}
