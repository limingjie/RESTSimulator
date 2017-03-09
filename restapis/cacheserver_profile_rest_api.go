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

// CacheServerProfiles map
var CacheServerProfiles = make(map[string]models.CacheServerProfile)

// PostCacheServerProfile - POST /profiles/cacheserver
func PostCacheServerProfile(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	profile := models.CacheServerProfile{}
	json.NewDecoder(r.Body).Decode(&profile)

	msg, _ := json.Marshal(profile)
	log.Println("PostCacheServerProfile", string(msg))

	profileName := profile.Profile.ProfileName
	if len(profileName) > 0 {
		_, ok := CacheServerProfiles[profileName]
		if ok {
			w.WriteHeader(409)
			fmt.Fprintf(w, "Error: CacheServer profile with same name already exists.")
		} else {
			profile.Profile.LastUpdated = time.Now().Format("2006/01/02 15:04:05")
			if len(profile.Profile.AccessPermission) == 0 {
				profile.Profile.AccessPermission = "ReadWrite"
			}
			CacheServerProfiles[profileName] = profile

			w.WriteHeader(201)
			fmt.Fprintf(w, "Succeed.")
		}
	} else {
		w.WriteHeader(400)
		fmt.Fprintf(w, "Error: Expecting Profile Name.")
	}
}

// GetCacheServerProfiles - GET /profiles/cacheserver
func GetCacheServerProfiles(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var profilesJSON bytes.Buffer
	profilesJSON.WriteString("{\"CacheServerProfile\":[")

	log.Println("GetCacheServerProfiles", strconv.Itoa(len(CacheServerProfiles)))

	ok := false
	for _, profile := range CacheServerProfiles {
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

// GetCacheServerProfile - GET /profiles/cacheserver/:name
func GetCacheServerProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	profile, ok := CacheServerProfiles[ps.ByName("name")]

	log.Println("GetCacheServerProfiles", ps.ByName("name"))

	if ok {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		profileJSON, _ := json.Marshal(profile)
		fmt.Fprintf(w, "%s", profileJSON)
	} else {
		w.WriteHeader(404)
		fmt.Fprintf(w, "Error: CacheServer profile does not exist.")
	}
}

// PutCacheServerProfile - PUT /profiles/cacheserver/:name
func PutCacheServerProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	profile := models.CacheServerProfile{}
	json.NewDecoder(r.Body).Decode(&profile)

	msg, _ := json.Marshal(profile)
	log.Println("PutCacheServerProfile", string(msg))

	profileName := ps.ByName("name")
	_, ok := CacheServerProfiles[profileName]
	if ok {
		if profile.Profile.ProfileName == profileName {
			profile.Profile.LastUpdated = time.Now().Format("2006/01/02 15:04:05")
			if len(profile.Profile.AccessPermission) == 0 {
				profile.Profile.AccessPermission = "ReadWrite"
			}
			CacheServerProfiles[profileName] = profile
			w.WriteHeader(200)
			fmt.Fprintf(w, "Succeed.")
		} else {
			w.WriteHeader(409)
			fmt.Fprintf(w, "Error: CacheServer profile name does not match.")
		}
	} else {
		w.WriteHeader(404)
		fmt.Fprintf(w, "Error: CacheServer profile does not exist.")
	}
}

// DeleteCacheServerProfile - DELETE /profiles/cacheserver/:name
func DeleteCacheServerProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	profileName := ps.ByName("name")

	log.Println("DeleteCacheServerProfile", profileName)

	_, ok := CacheServerProfiles[profileName]
	if ok {
		delete(CacheServerProfiles, profileName)
		w.WriteHeader(200)
		fmt.Fprintf(w, "Succeed.")
	} else {
		w.WriteHeader(404)
		fmt.Fprintf(w, "Error: CacheServer profile does not exist.")
	}
}

// PostCacheServerTangoConfigConversion - POST /profiles/cacheserver/tangoconfig
func PostCacheServerTangoConfigConversion(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s", "{\"tangoConfig\": \"<?xml version='1.0'?><pre><tend><tobexml></tobexml></tend></pre>\"}")
}

// PostCacheServerCacheConfigConversion - POST /profiles/cacheserver/cacheconfig
func PostCacheServerCacheConfigConversion(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s", "{\"cacheConfigXml\": \"<?xml version='1.0'?><pre><tend><tobexml></tobexml></tend></pre>\"}")
}
