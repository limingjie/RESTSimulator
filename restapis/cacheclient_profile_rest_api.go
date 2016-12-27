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

// CacheClientProfiles map
var CacheClientProfiles = make(map[string]models.CacheProfile)

// PostCacheClientProfile - POST /profiles/cacheserver
func PostCacheClientProfile(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	profile := models.CacheProfile{}
	json.NewDecoder(r.Body).Decode(&profile)

	msg, _ := json.Marshal(profile)
	logger.Logger("PostCacheClientProfile", string(msg))

	profileName := profile.Profile.ProfileName
	if len(profileName) > 0 {
		_, ok := CacheClientProfiles[profileName]
		if ok {
			w.WriteHeader(409)
			fmt.Fprintf(w, "Error: Enterprise profile with same name already exists.")
		} else {
			profile.Profile.LastUpdated = time.Now().Format("2006/01/02 15:04:05")
			if len(profile.Profile.AccessPermission) == 0 {
				profile.Profile.AccessPermission = "READWRITE"
			}
			CacheClientProfiles[profileName] = profile

			w.WriteHeader(201)
			fmt.Fprintf(w, "Succeed.")
		}
	} else {
		w.WriteHeader(400)
		fmt.Fprintf(w, "Error: Expecting Profile Name.")
	}
}

// GetCacheClientProfiles - GET /profiles/cacheserver
func GetCacheClientProfiles(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var profilesJSON bytes.Buffer
	profilesJSON.WriteString("{\"cacheServerProfile\":[")

	logger.Logger("GetCacheClientProfiles", strconv.Itoa(len(CacheClientProfiles)))

	ok := false
	for _, profile := range CacheClientProfiles {
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

// GetCacheClientProfile - GET /profiles/cacheserver/:profilename
func GetCacheClientProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	profile, ok := CacheClientProfiles[ps.ByName("profilename")]

	logger.Logger("GetCacheClientProfiles", ps.ByName("profilename"))

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

// PutCacheClientProfile - PUT /profiles/cacheserver/:profilename
func PutCacheClientProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	profile := models.CacheProfile{}
	json.NewDecoder(r.Body).Decode(&profile)

	msg, _ := json.Marshal(profile)
	logger.Logger("PutCacheClientProfile", string(msg))

	profileName := ps.ByName("profilename")
	_, ok := CacheClientProfiles[profileName]
	if ok {
		if profile.Profile.ProfileName == profileName {
			profile.Profile.LastUpdated = time.Now().Format("2006/01/02 15:04:05")
			if len(profile.Profile.AccessPermission) == 0 {
				profile.Profile.AccessPermission = "READWRITE"
			}
			CacheClientProfiles[profileName] = profile
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

// DeleteCacheClientProfile - DELETE /profiles/cacheserver/:profilename
func DeleteCacheClientProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	profileName := ps.ByName("profilename")

	logger.Logger("DeleteCacheClientProfile", profileName)

	_, ok := CacheClientProfiles[profileName]
	if ok {
		delete(CacheClientProfiles, profileName)
		w.WriteHeader(200)
		fmt.Fprintf(w, "Succeed.")
	} else {
		w.WriteHeader(404)
		fmt.Fprintf(w, "Error: Enterprise profile does not exist.")
	}
}