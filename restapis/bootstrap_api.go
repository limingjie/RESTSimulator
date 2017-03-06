package restapis

import (
	"encoding/json"
	"fmt"
	"net/http"

	"../logger"
	"../models"
	"github.com/julienschmidt/httprouter"
)

// CGInfo struct
var CGInfo = models.CGInfo{}

// BootstrapCG struct
var BootstrapCG = models.BootstrapCG{}

// GetCGInfo - GET /cginfo
func GetCGInfo(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	cgInfoJSON, _ := json.Marshal(CGInfo)
	logger.Logger("GetCGInfo", string(cgInfoJSON))

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	// fmt.Fprintf(w, "%s", cgInfoJSON)
	fmt.Fprintf(w, "{\"CGHostURI\":\"https://localhost:8889/siebel/v1.0/cloudgateway\"}")
}

// PostCGInfo - POST /cginfo
func PostCGInfo(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	json.NewDecoder(r.Body).Decode(&CGInfo)

	CGInfoJSON, _ := json.Marshal(CGInfo)
	logger.Logger("PostCGInfo", string(CGInfoJSON))

	w.WriteHeader(200)
}

// GetBootstrapCG - GET /bootstrapCG
func GetBootstrapCG(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	BootstrapCGJSON, _ := json.Marshal(BootstrapCG)
	logger.Logger("GetBootstrapCG", string(BootstrapCGJSON))

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	// fmt.Fprintf(w, "%s", BootstrapCGJSON)
	fmt.Fprintf(w, "{\"registryPort\":\"4330\", \"registryPassword\":\"sadmin\", \"registryUserName\":\"sadmin\", \"SecurityProfile\": \"Gateway\", \"PrimaryLanguage\":\"enu\"}")
}

// PostBootstrapCG - POST /bootstrapCG
func PostBootstrapCG(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	json.NewDecoder(r.Body).Decode(&BootstrapCG)

	BootstrapCGJSON, _ := json.Marshal(BootstrapCG)
	logger.Logger("PostBootstrapCG", string(BootstrapCGJSON))

	w.WriteHeader(200)
}
