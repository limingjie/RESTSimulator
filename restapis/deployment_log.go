package restapis

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// GetDeploymentLog - Get deployment log
func GetDeploymentLog(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	log.Println("GetDeploymentLog", "log")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "{\"Log\":\"Line 01: Warning - there should be a new line\\nLine 02: Fatal Error - if there was not a new line.\\n\"}")
}
