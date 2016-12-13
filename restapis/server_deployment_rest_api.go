package restapis

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"../models"
	"github.com/julienschmidt/httprouter"
)

// ServerDeployments map
var ServerDeployments = make(map[string]models.ServerDeployment)

// PostServerDeployment - POST /deployments/servers
func PostServerDeployment(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	deployment := models.ServerDeployment{}
	json.NewDecoder(r.Body).Decode(&deployment)

	siebelServer := deployment.ServerDeployParams.SiebelServer
	if len(siebelServer) > 0 {
		_, ok := ServerDeployments[siebelServer]
		if ok {
			w.WriteHeader(409)
			fmt.Fprintf(w, "Error: Server deployment with same name already exists.")
		} else {
			ServerDeployments[siebelServer] = deployment

			w.WriteHeader(201)
			fmt.Fprintf(w, "Succeed.")
		}
	} else {
		w.WriteHeader(400)
		fmt.Fprintf(w, "Error: Expecting Deployment Name.")
	}
}

// GetServerDeployments - GET /deployments/servers
func GetServerDeployments(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var deploymentsJSON bytes.Buffer
	deploymentsJSON.WriteString("{\"serverDeployment\":[")

	ok := false
	for _, deployment := range ServerDeployments {
		if ok {
			deploymentsJSON.WriteString(",")
		} else {
			ok = true
		}
		deploymentJSON, _ := json.Marshal(deployment)
		deploymentsJSON.Write(deploymentJSON)
	}

	deploymentsJSON.WriteString("]}")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s", deploymentsJSON.String())
}

// GetServerDeployment - GET /deployments/servers/:deploymentname
func GetServerDeployment(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	deployment, ok := ServerDeployments[ps.ByName("deploymentname")]
	if ok {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		deploymentJSON, _ := json.Marshal(deployment)
		fmt.Fprintf(w, "%s", deploymentJSON)
	} else {
		w.WriteHeader(404)
		fmt.Fprintf(w, "Error: Server deployment does not exist.")
	}
}

// PutServerDeployment - PUT /deployments/servers/:deploymentname
func PutServerDeployment(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	deployment := models.ServerDeployment{}
	json.NewDecoder(r.Body).Decode(&deployment)

	siebelServer := ps.ByName("deploymentname")
	_, ok := ServerDeployments[siebelServer]
	if ok {
		if deployment.ServerDeployParams.SiebelServer == siebelServer {
			ServerDeployments[siebelServer] = deployment
			w.WriteHeader(200)
			fmt.Fprintf(w, "Succeed.")
		} else {
			w.WriteHeader(409)
			fmt.Fprintf(w, "Error: Server deployment name does not match.")
		}
	} else {
		w.WriteHeader(404)
		fmt.Fprintf(w, "Error: Server deployment does not exist.")
	}
}

// DeleteServerDeployment - DELETE /deployments/servers/:deploymentname
func DeleteServerDeployment(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	siebelServer := ps.ByName("deploymentname")
	_, ok := ServerDeployments[siebelServer]
	if ok {
		delete(ServerDeployments, siebelServer)
		w.WriteHeader(200)
		fmt.Fprintf(w, "Succeed.")
	} else {
		w.WriteHeader(404)
		fmt.Fprintf(w, "Error: Server deployment does not exist.")
	}
}
