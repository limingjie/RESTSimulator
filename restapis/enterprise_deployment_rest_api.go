package restapis

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"../models"
	"github.com/julienschmidt/httprouter"
)

// EnterpriseDeployments map
var EnterpriseDeployments = make(map[string]models.EnterpriseDeployment)

// PostEnterpriseDeployment - POST /deployments/enterprises
func PostEnterpriseDeployment(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	deployment := models.EnterpriseDeployment{}
	json.NewDecoder(r.Body).Decode(&deployment)

	enterpriseServer := deployment.EnterpriseDeployParams.EnterpriseServer
	if len(enterpriseServer) > 0 {
		_, ok := EnterpriseDeployments[enterpriseServer]
		if ok {
			w.WriteHeader(409)
			fmt.Fprintf(w, "Error: Enterprise deployment with same name already exists.")
		} else {
			EnterpriseDeployments[enterpriseServer] = deployment

			w.WriteHeader(201)
			fmt.Fprintf(w, "Succeed.")
		}
	} else {
		w.WriteHeader(400)
		fmt.Fprintf(w, "Error: Expecting Deployment Name.")
	}
}

// GetEnterpriseDeployments - GET /deployments/enterprises
func GetEnterpriseDeployments(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var deploymentsJSON bytes.Buffer
	deploymentsJSON.WriteString("{\"enterpriseDeployment\":[")

	ok := false
	for _, deployment := range EnterpriseDeployments {
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

// GetEnterpriseDeployment - GET /deployments/enterprises/:deploymentname
func GetEnterpriseDeployment(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	deployment, ok := EnterpriseDeployments[ps.ByName("deploymentname")]
	if ok {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		deploymentJSON, _ := json.Marshal(deployment)
		fmt.Fprintf(w, "%s", deploymentJSON)
	} else {
		w.WriteHeader(404)
		fmt.Fprintf(w, "Error: Enterprise deployment does not exist.")
	}
}

// PutEnterpriseDeployment - PUT /deployments/enterprises/:deploymentname
func PutEnterpriseDeployment(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	deployment := models.EnterpriseDeployment{}
	json.NewDecoder(r.Body).Decode(&deployment)

	enterpriseServer := ps.ByName("deploymentname")
	_, ok := EnterpriseDeployments[enterpriseServer]
	if ok {
		if deployment.EnterpriseDeployParams.EnterpriseServer == enterpriseServer {
			EnterpriseDeployments[enterpriseServer] = deployment
			w.WriteHeader(200)
			fmt.Fprintf(w, "Succeed.")
		} else {
			w.WriteHeader(409)
			fmt.Fprintf(w, "Error: Enterprise deployment name does not match.")
		}
	} else {
		w.WriteHeader(404)
		fmt.Fprintf(w, "Error: Enterprise deployment does not exist.")
	}
}

// DeleteEnterpriseDeployment - DELETE /deployments/enterprises/:deploymentname
func DeleteEnterpriseDeployment(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	enterpriseServer := ps.ByName("deploymentname")
	_, ok := EnterpriseDeployments[enterpriseServer]
	if ok {
		delete(EnterpriseDeployments, enterpriseServer)
		w.WriteHeader(200)
		fmt.Fprintf(w, "Succeed.")
	} else {
		w.WriteHeader(404)
		fmt.Fprintf(w, "Error: Enterprise deployment does not exist.")
	}
}
