package restapis

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"../logger"
	"../models"
	"github.com/julienschmidt/httprouter"
)

// CacheInfo map
var CacheInfo []models.CacheInfo
var index = 0

// GetCacheInfo - GET /cache
func GetCacheInfo(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var profilesJSON bytes.Buffer
	profilesJSON.WriteString("{\"cacheInfo\":[")

	if len(CacheInfo) < 5 {
		for i := 0; i < 20; i++ {
			cache := models.CacheInfo{
				Name:   "Cache" + strconv.Itoa(index),
				Status: "Good",
				Size:   "100",
			}

			index++

			CacheInfo = append(CacheInfo, cache)
		}
	}

	logger.Logger("GetCacheInfo", strconv.Itoa(len(CacheInfo)))

	ok := false
	for _, profile := range CacheInfo {
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
