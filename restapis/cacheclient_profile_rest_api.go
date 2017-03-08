package restapis

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"../models"
	"github.com/julienschmidt/httprouter"
)

// CacheClientProfiles map
var CacheClientProfiles = make(map[string]models.CacheClientProfile)

// PostCacheClientProfile - POST /profiles/cacheserver
func PostCacheClientProfile(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	profile := models.CacheClientProfile{}
	json.NewDecoder(r.Body).Decode(&profile)

	msg, _ := json.Marshal(profile)
	log.Println("PostCacheClientProfile", string(msg))

	profileName := profile.Profile.ProfileName
	if len(profileName) > 0 {
		_, ok := CacheClientProfiles[profileName]
		if ok {
			w.WriteHeader(409)
			fmt.Fprintf(w, "Error: CacheClient profile with same name already exists.")
		} else {
			profile.Profile.LastUpdated = time.Now().Format("2006/01/02 15:04:05")
			if len(profile.Profile.AccessPermission) == 0 {
				profile.Profile.AccessPermission = "ReadWrite"
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
	profilesJSON.WriteString("{\"CacheClientProfile\":[")

	log.Println("GetCacheClientProfiles", strconv.Itoa(len(CacheClientProfiles)))

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

	log.Println("GetCacheClientProfiles", ps.ByName("profilename"))

	if ok {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		profileJSON, _ := json.Marshal(profile)
		fmt.Fprintf(w, "%s", profileJSON)
	} else {
		w.WriteHeader(404)
		fmt.Fprintf(w, "Error: CacheClient profile does not exist.")
	}
}

// PutCacheClientProfile - PUT /profiles/cacheserver/:profilename
func PutCacheClientProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	profile := models.CacheClientProfile{}
	json.NewDecoder(r.Body).Decode(&profile)

	msg, _ := json.Marshal(profile)
	log.Println("PutCacheClientProfile", string(msg))

	profileName := ps.ByName("profilename")
	_, ok := CacheClientProfiles[profileName]
	if ok {
		if profile.Profile.ProfileName == profileName {
			profile.Profile.LastUpdated = time.Now().Format("2006/01/02 15:04:05")
			if len(profile.Profile.AccessPermission) == 0 {
				profile.Profile.AccessPermission = "ReadWrite"
			}
			CacheClientProfiles[profileName] = profile
			w.WriteHeader(200)
			fmt.Fprintf(w, "Succeed.")
		} else {
			w.WriteHeader(409)
			fmt.Fprintf(w, "Error: CacheClient profile name does not match.")
		}
	} else {
		w.WriteHeader(404)
		fmt.Fprintf(w, "Error: CacheClient profile does not exist.")
	}
}

// DeleteCacheClientProfile - DELETE /profiles/cacheserver/:profilename
func DeleteCacheClientProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	profileName := ps.ByName("profilename")

	log.Println("DeleteCacheClientProfile", profileName)

	_, ok := CacheClientProfiles[profileName]
	if ok {
		delete(CacheClientProfiles, profileName)
		w.WriteHeader(200)
		fmt.Fprintf(w, "Succeed.")
	} else {
		w.WriteHeader(404)
		fmt.Fprintf(w, "Error: CacheClient profile does not exist.")
	}
}

// PostCacheClientCacheConfigConversion - POST /profiles/cacheclient/tangoconfig
func PostCacheClientCacheConfigConversion(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s", "{\"cacheConfig\": {\"cacheConfigXml\": \"<?xml version='1.0'?><pre><tend><tobexml></tobexml></tend></pre>\"}}")
}
