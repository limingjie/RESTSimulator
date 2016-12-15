package main

import (
	"fmt"
	"log"
	"net/http"

	"./restapis"
	"github.com/julienschmidt/httprouter"
)

var port = ":8889"

// Index - For /
func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}

func main() {
	router := httprouter.New()
	router.GET("/", Index)

	// Enterprise Profile APIs
	router.POST("/cloudgateway/v1.0/profiles/enterprises", restapis.PostEnterpriseProfile)
	router.GET("/cloudgateway/v1.0/profiles/enterprises", restapis.GetEnterpriseProfiles)
	router.GET("/cloudgateway/v1.0/profiles/enterprises/:profilename", restapis.GetEnterpriseProfile)
	router.PUT("/cloudgateway/v1.0/profiles/enterprises/:profilename", restapis.PutEnterpriseProfile)
	router.DELETE("/cloudgateway/v1.0/profiles/enterprises/:profilename", restapis.DeleteEnterpriseProfile)

	// Enterprise Deployment APIs
	router.POST("/cloudgateway/v1.0/deployments/enterprises", restapis.PostEnterpriseDeployment)
	router.GET("/cloudgateway/v1.0/deployments/enterprises", restapis.GetEnterpriseDeployments)
	router.GET("/cloudgateway/v1.0/deployments/enterprises/:deploymentname", restapis.GetEnterpriseDeployment)
	router.PUT("/cloudgateway/v1.0/deployments/enterprises/:deploymentname", restapis.PutEnterpriseDeployment)
	router.DELETE("/cloudgateway/v1.0/deployments/enterprises/:deploymentname", restapis.DeleteEnterpriseDeployment)

	// Server Profile APIs
	router.POST("/cloudgateway/v1.0/profiles/servers", restapis.PostServerProfile)
	router.GET("/cloudgateway/v1.0/profiles/servers", restapis.GetServerProfiles)
	router.GET("/cloudgateway/v1.0/profiles/servers/:profilename", restapis.GetServerProfile)
	router.PUT("/cloudgateway/v1.0/profiles/servers/:profilename", restapis.PutServerProfile)
	router.DELETE("/cloudgateway/v1.0/profiles/servers/:profilename", restapis.DeleteServerProfile)

	// Server Deployment APIs
	router.POST("/cloudgateway/v1.0/deployments/servers", restapis.PostServerDeployment)
	router.GET("/cloudgateway/v1.0/deployments/servers", restapis.GetServerDeployments)
	router.GET("/cloudgateway/v1.0/deployments/servers/:deploymentname", restapis.GetServerDeployment)
	router.PUT("/cloudgateway/v1.0/deployments/servers/:deploymentname", restapis.PutServerDeployment)
	router.DELETE("/cloudgateway/v1.0/deployments/servers/:deploymentname", restapis.DeleteServerDeployment)

	// SWSM Profile APIs
	router.POST("/cloudgateway/v1.0/profiles/swsm", restapis.PostSWSMProfile)
	router.GET("/cloudgateway/v1.0/profiles/swsm", restapis.GetSWSMProfiles)
	router.GET("/cloudgateway/v1.0/profiles/swsm/:profilename", restapis.GetSWSMProfile)
	router.PUT("/cloudgateway/v1.0/profiles/swsm/:profilename", restapis.PutSWSMProfile)
	router.DELETE("/cloudgateway/v1.0/profiles/swsm/:profilename", restapis.DeleteSWSMProfile)

	// SWSM Deployment APIs
	router.POST("/cloudgateway/v1.0/deployments/swsm", restapis.PostSWSMDeployment)
	router.GET("/cloudgateway/v1.0/deployments/swsm", restapis.GetSWSMDeployments)
	router.GET("/cloudgateway/v1.0/deployments/swsm/:deploymentname", restapis.GetSWSMDeployment)
	router.PUT("/cloudgateway/v1.0/deployments/swsm/:deploymentname", restapis.PutSWSMDeployment)
	router.DELETE("/cloudgateway/v1.0/deployments/swsm/:deploymentname", restapis.DeleteSWSMDeployment)

	// CacheServer Profile APIs
	router.POST("/cloudgateway/v1.0/profiles/cacheserver", restapis.PostCacheServerProfile)
	router.GET("/cloudgateway/v1.0/profiles/cacheserver", restapis.GetCacheServerProfiles)
	router.GET("/cloudgateway/v1.0/profiles/cacheserver/:profilename", restapis.GetCacheServerProfile)
	router.PUT("/cloudgateway/v1.0/profiles/cacheserver/:profilename", restapis.PutCacheServerProfile)
	router.DELETE("/cloudgateway/v1.0/profiles/cacheserver/:profilename", restapis.DeleteCacheServerProfile)

	// CacheServer Deployment APIs
	router.POST("/cloudgateway/v1.0/deployments/cacheserver", restapis.PostCacheServerDeployment)
	router.GET("/cloudgateway/v1.0/deployments/cacheserver", restapis.GetCacheServerDeployments)
	router.GET("/cloudgateway/v1.0/deployments/cacheserver/:profilename", restapis.GetCacheServerDeployment)
	router.PUT("/cloudgateway/v1.0/deployments/cacheserver/:profilename", restapis.PutCacheServerDeployment)
	router.DELETE("/cloudgateway/v1.0/deployments/cacheserver/:profilename", restapis.DeleteCacheServerDeployment)

	// CacheClient Profile APIs
	router.POST("/cloudgateway/v1.0/profiles/cacheclient", restapis.PostCacheClientProfile)
	router.GET("/cloudgateway/v1.0/profiles/cacheclient", restapis.GetCacheClientProfiles)
	router.GET("/cloudgateway/v1.0/profiles/cacheclient/:profilename", restapis.GetCacheClientProfile)
	router.PUT("/cloudgateway/v1.0/profiles/cacheclient/:profilename", restapis.PutCacheClientProfile)
	router.DELETE("/cloudgateway/v1.0/profiles/cacheclient/:profilename", restapis.DeleteCacheClientProfile)

	fmt.Printf("RESTful API URL - http://localhost" + port + "/cloudgateway/v1.0/\n")
	log.Fatal(http.ListenAndServe(port, router))
}
