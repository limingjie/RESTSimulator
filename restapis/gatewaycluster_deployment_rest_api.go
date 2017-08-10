package restapis

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"strconv"

	"../models"
	"github.com/julienschmidt/httprouter"
)

// GatewayClusterDeployments map
var GatewayClusterDeployments = make(map[string]models.GatewayClusterDeployment)

func saveGatewayClusterDeployments() {
	WriteFile(filepath.Clean("GatewayData/GatewayClusterDeployments.json"), GatewayClusterDeployments)
}

func deployGatewayClusterProfile(profileName string) {
	profile, ok := GatewayClusterProfiles[profileName]
	if ok {
		profile.Profile.Deploy()
		GatewayClusterProfiles[profileName] = profile
	}
}

func undeployGatewayClusterProfile(profileName string) {
	profile, ok := GatewayClusterProfiles[profileName]
	if ok {
		profile.Profile.Undeploy()
		GatewayClusterProfiles[profileName] = profile
	}
}

// PostGatewayClusterDeployment - POST /deployments/gatewaycluster
func PostGatewayClusterDeployment(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	deployment := models.GatewayClusterDeployment{}
	json.NewDecoder(r.Body).Decode(&deployment)

	msg, _ := json.Marshal(deployment)
	log.Println(r.RemoteAddr, "PostGatewayClusterDeployment", string(msg))

	_, ok := GatewayClusterProfiles[deployment.Deployment.ProfileName]
	if !ok {
		w.WriteHeader(400)
		fmt.Fprintf(w, "Error: Inexist Deployment Profile.")
		return
	}

	deploymentName := "GatewayCluster"
	_, ok = GatewayClusterDeployments[deploymentName]
	if ok {
		w.WriteHeader(409)
		fmt.Fprintf(w, "Error: Gateway Cluster deployment with same name already exists.")
	} else {
		deployment.Deployment.Check()
		GatewayClusterDeployments[deploymentName] = deployment

		deployGatewayClusterProfile(deployment.Deployment.ProfileName)

		w.WriteHeader(201)
		fmt.Fprintf(w, "Succeed.")

		saveGatewayClusterDeployments()
	}
}

// GetGatewayClusterDeployments - GET /deployments/gatewaycluster
func GetGatewayClusterDeployments(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var deploymentsJSON bytes.Buffer
	deploymentsJSON.WriteString("{\"GatewayClusterDeployment\":[")

	log.Println(r.RemoteAddr, "GetGatewayClusterDeployments", strconv.Itoa(len(GatewayClusterDeployments)))

	ok := false
	for _, deployment := range GatewayClusterDeployments {
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

// PutGatewayClusterDeployment - PUT /deployments/gatewaycluster
func PutGatewayClusterDeployment(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	deployment := models.GatewayClusterDeployment{}
	json.NewDecoder(r.Body).Decode(&deployment)

	msg, _ := json.Marshal(deployment)
	log.Println(r.RemoteAddr, "PutGatewayClusterDeployment", string(msg))

	_, ok := GatewayClusterProfiles[deployment.Deployment.ProfileName]
	if !ok {
		w.WriteHeader(400)
		fmt.Fprintf(w, "Error: Expecting Deployment Profile.")
		return
	}

	deploymentName := "GatewayCluster"
	_, ok = GatewayClusterDeployments[deploymentName]
	if ok {
		oldProfileName := GatewayClusterDeployments[deploymentName].Deployment.ProfileName
		if deployment.Deployment.ProfileName != oldProfileName {
			deployGatewayClusterProfile(deployment.Deployment.ProfileName)
			undeployGatewayClusterProfile(oldProfileName)
		}

		deployment.Deployment.Check()
		GatewayClusterDeployments[deploymentName] = deployment

		w.WriteHeader(200)
		fmt.Fprintf(w, "Succeed.")

		saveGatewayClusterDeployments()
	}
}

// DeleteGatewayClusterDeployment - DELETE /deployments/gatewaycluster
func DeleteGatewayClusterDeployment(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	deploymentName := "GatewayCluster"

	log.Println(r.RemoteAddr, "DeleteGatewayClusterDeployment", deploymentName)

	deployment, ok := GatewayClusterDeployments[deploymentName]
	if ok {
		_, ok := GatewayClusterProfiles[deployment.Deployment.ProfileName]
		if ok {
			undeployGatewayClusterProfile(deployment.Deployment.ProfileName)
		}

		delete(GatewayClusterDeployments, deploymentName)

		w.WriteHeader(200)
		fmt.Fprintf(w, "Succeed.")

		saveGatewayClusterDeployments()
	} else {
		w.WriteHeader(404)
		fmt.Fprintf(w, "Error: Gateway Cluster deployment does not exist.")
	}
}
