package restapis

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"../models"
	"github.com/julienschmidt/httprouter"
)

// Caches map
var Caches = make(map[string]models.CacheInfo)

var cacheStatus = []string{"Active", "Inactive"}

// randInt - Generate integer in [min, max)
func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

// ramdomString - Generate random string
func ramdomString(len int) string {
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		bytes[i] = byte(randInt(97, 122))
	}
	return string(bytes)
}

// createCaches - Create Caches
func createCaches(count int) {
	for i := 0; i < count; i++ {
		name := "cache-" + ramdomString(10)
		Caches[name] = models.CacheInfo{
			Name:   name,
			Status: cacheStatus[rand.Intn(4)/3],
			Size:   strconv.Itoa(randInt(1, 100)),
		}
	}
}

// cacheGenerator - Generate and update caches
func cacheGenerator() {
	// Update cache size.
	for _, cache := range Caches {
		v, err := strconv.Atoi(cache.Size)
		if err != nil {
			v = 0
		}
		cache.Size = strconv.Itoa(v + randInt(50, 100))
		Caches[cache.Name] = cache
	}

	// Generate new caches
	if len(Caches) < 5 {
		createCaches(randInt(5, 21))
	} else {
		createCaches(randInt(1, 3))
	}
}

// GetCaches - GET /cache
func GetCaches(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var cachesJSON bytes.Buffer
	cachesJSON.WriteString("{\"cacheInfo\":[")

	log.Println("GetCaches", strconv.Itoa(len(Caches)))

	ok := false
	for _, cache := range Caches {
		if ok {
			cachesJSON.WriteString(",")
		} else {
			ok = true
		}
		cacheJSON, _ := json.Marshal(cache)
		cachesJSON.Write(cacheJSON)
	}

	cachesJSON.WriteString("]}")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s", cachesJSON.String())

	// update cache
	cacheGenerator()
}

// GetCache - GET /cache/:cacheName
func GetCache(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	cache, ok := Caches[ps.ByName("cachename")]

	log.Println("GetCache", ps.ByName("cachename"))

	if ok {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		cacheJSON, _ := json.Marshal(cache)
		fmt.Fprintf(w, "%s", cacheJSON)
	} else {
		w.WriteHeader(404)
		fmt.Fprintf(w, "Error: Cache does not exist.")
	}
}

// DeleteCaches - DELETE /cache
func DeleteCaches(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	log.Println("DeleteCaches", strconv.Itoa(len(Caches)))

	// Let GC do the rest :)
	Caches = make(map[string]models.CacheInfo)

	w.WriteHeader(200)
	fmt.Fprintf(w, "Succeed.")
}

// DeleteCache - DELETE /cache/:cachename
func DeleteCache(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	cacheName := ps.ByName("cachename")

	log.Println("DeleteCache", cacheName)

	_, ok := Caches[cacheName]
	if ok {
		delete(Caches, cacheName)
		w.WriteHeader(200)
		fmt.Fprintf(w, "Succeed.")
	} else {
		w.WriteHeader(404)
		fmt.Fprintf(w, "Error: Cache does not exist.")
	}
}

// ClearCaches - DELETE /clearcache
func ClearCaches(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	log.Println("ClearCaches", strconv.Itoa(len(Caches)))

	// Update cache size.
	for _, cache := range Caches {
		cache.Size = "0"
		Caches[cache.Name] = cache
	}

	w.WriteHeader(200)
	fmt.Fprintf(w, "Succeed.")
}

// ClearCache - DELETE /clearcache/:cachename
func ClearCache(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	cacheName := ps.ByName("cachename")

	log.Println("ClearCache", cacheName)

	cache, ok := Caches[cacheName]
	if ok {
		cache.Size = "0"
		Caches[cache.Name] = cache
		w.WriteHeader(200)
		fmt.Fprintf(w, "Succeed.")
	} else {
		w.WriteHeader(404)
		fmt.Fprintf(w, "Error: Cache does not exist.")
	}
}
