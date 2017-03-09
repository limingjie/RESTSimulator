package restapis

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"../models"
	"github.com/julienschmidt/httprouter"
)

// ServerDeployments map
var ServerDeployments = make(map[string]models.ServerDeployment)

func deployServerProfile(profileName string) {
	profile, ok := ServerProfiles[profileName]
	if ok {
		profile.Profile.Deploy()
		ServerProfiles[profileName] = profile
	}
}

func undeployServerProfile(profileName string) {
	profile, ok := ServerProfiles[profileName]
	if ok {
		profile.Profile.Undeploy()
		ServerProfiles[profileName] = profile
	}
}

// PostServerDeployment - POST /deployments/servers
func PostServerDeployment(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	deployment := models.ServerDeployment{}
	json.NewDecoder(r.Body).Decode(&deployment)

	msg, _ := json.Marshal(deployment)
	log.Println("PostServerDeployment", string(msg))

	_, ok := ServerProfiles[deployment.Deployment.ProfileName]
	if !ok {
		w.WriteHeader(400)
		fmt.Fprintf(w, "Error: Inexist Deployment Profile.")
		return
	}

	deploymentName := deployment.ServerDeployParams.SiebelServer
	if len(deploymentName) > 0 {
		_, ok := ServerDeployments[deploymentName]
		if ok {
			w.WriteHeader(409)
			fmt.Fprintf(w, "Error: Server deployment with same name already exists.")
		} else {
			deployment.Deployment.Check()
			ServerDeployments[deploymentName] = deployment

			deployServerProfile(deployment.Deployment.ProfileName)

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
	deploymentsJSON.WriteString("{\"ServerDeployment\":[")

	log.Println("GetServerDeployments", strconv.Itoa(len(ServerDeployments)))

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

// GetServerDeployment - GET /deployments/servers/:name
func GetServerDeployment(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	deployment, ok := ServerDeployments[ps.ByName("name")]

	log.Println("GetServerDeployment", ps.ByName("name"))

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

// PutServerDeployment - PUT /deployments/servers/:name
func PutServerDeployment(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	deployment := models.ServerDeployment{}
	json.NewDecoder(r.Body).Decode(&deployment)

	msg, _ := json.Marshal(deployment)
	log.Println("PutServerDeployment", string(msg))

	_, ok := ServerProfiles[deployment.Deployment.ProfileName]
	if !ok {
		w.WriteHeader(400)
		fmt.Fprintf(w, "Error: Expecting Deployment Profile.")
		return
	}

	deploymentName := ps.ByName("name")
	_, ok = ServerDeployments[deploymentName]
	if ok {
		if deployment.ServerDeployParams.SiebelServer == deploymentName {
			oldProfileName := ServerDeployments[deploymentName].Deployment.ProfileName
			if deployment.Deployment.ProfileName != oldProfileName {
				deployServerProfile(deployment.Deployment.ProfileName)
				undeployServerProfile(oldProfileName)
			}

			ServerDeployments[deploymentName] = deployment

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

// DeleteServerDeployment - DELETE /deployments/servers/:name
func DeleteServerDeployment(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	deploymentName := ps.ByName("name")

	log.Println("DeleteServerDeployment", deploymentName)

	deployment, ok := ServerDeployments[deploymentName]
	if ok {
		_, ok := ServerProfiles[deployment.Deployment.ProfileName]
		if ok {
			undeployServerProfile(deployment.Deployment.ProfileName)
		}

		delete(ServerDeployments, deploymentName)

		w.WriteHeader(200)
		fmt.Fprintf(w, "Succeed.")
	} else {
		w.WriteHeader(404)
		fmt.Fprintf(w, "Error: Server deployment does not exist.")
	}
}
