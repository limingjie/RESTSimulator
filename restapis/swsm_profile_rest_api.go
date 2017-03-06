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

// SWSMProfiles map
var SWSMProfiles = make(map[string]models.SWSMProfile)

// PostSWSMProfile - POST /profiles/swsm
func PostSWSMProfile(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	profile := models.SWSMProfile{}

	json.NewDecoder(r.Body).Decode(&profile)

	msg, _ := json.Marshal(profile)
	logger.Logger("PostSWSMProfile", string(msg))

	profileName := profile.Profile.ProfileName
	if len(profileName) > 0 {
		_, ok := SWSMProfiles[profileName]
		if ok {
			w.WriteHeader(409)
			fmt.Fprintf(w, "Error: SWSM profile with same name already exists.")
		} else {
			profile.Profile.LastUpdated = time.Now().Format("2006/01/02 15:04:05")
			if len(profile.Profile.AccessPermission) == 0 {
				profile.Profile.AccessPermission = "ReadWrite"
			}
			SWSMProfiles[profileName] = profile

			w.WriteHeader(201)
			fmt.Fprintf(w, "Succeed.")
		}
	} else {
		w.WriteHeader(400)
		fmt.Fprintf(w, "Error: Expecting Profile Name.")
	}
}

// GetSWSMProfiles - GET /profiles/swsm
func GetSWSMProfiles(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var profilesJSON bytes.Buffer
	profilesJSON.WriteString("{\"SWSMProfile\":[")

	logger.Logger("GetSWSMProfiles", strconv.Itoa(len(SWSMProfiles)))

	ok := false
	for _, profile := range SWSMProfiles {
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

// GetSWSMProfile - GET /profiles/swsm/:profilename
func GetSWSMProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	profile, ok := SWSMProfiles[ps.ByName("profilename")]

	logger.Logger("GetSWSMProfile", ps.ByName("profilename"))

	if ok {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		profileJSON, _ := json.Marshal(profile)
		fmt.Fprintf(w, "%s", profileJSON)
	} else {
		w.WriteHeader(404)
		fmt.Fprintf(w, "Error: SWSM profile does not exist.")
	}
}

// PutSWSMProfile - PUT /profiles/swsm/:profilename
func PutSWSMProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	profile := models.SWSMProfile{}
	json.NewDecoder(r.Body).Decode(&profile)

	msg, _ := json.Marshal(profile)
	logger.Logger("PutSWSMProfile", string(msg))

	profileName := ps.ByName("profilename")
	_, ok := SWSMProfiles[profileName]
	if ok {
		if profile.Profile.ProfileName == profileName {
			profile.Profile.LastUpdated = time.Now().Format("2006/01/02 15:04:05")
			if len(profile.Profile.AccessPermission) == 0 {
				profile.Profile.AccessPermission = "ReadWrite"
			}
			SWSMProfiles[profileName] = profile
			w.WriteHeader(200)
			fmt.Fprintf(w, "Succeed.")
		} else {
			w.WriteHeader(409)
			fmt.Fprintf(w, "Error: SWSM profile name does not match.")
		}
	} else {
		w.WriteHeader(404)
		fmt.Fprintf(w, "Error: SWSM profile does not exist.")
	}
}

// DeleteSWSMProfile - DELETE /profiles/swsm/:profilename
func DeleteSWSMProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	profileName := ps.ByName("profilename")

	logger.Logger("DeleteSWSMProfile", profileName)

	_, ok := SWSMProfiles[profileName]
	if ok {
		delete(SWSMProfiles, profileName)
		w.WriteHeader(200)
		fmt.Fprintf(w, "Succeed.")
	} else {
		w.WriteHeader(404)
		fmt.Fprintf(w, "Error: SWSM profile does not exist.")
	}
}
