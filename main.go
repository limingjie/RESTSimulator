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

	// CG Bootstrap
	router.GET("/siebel/v1.0/cginfo", restapis.GetCGInfo)
	router.POST("/siebel/v1.0/cginfo", restapis.PostCGInfo)
	router.GET("/siebel/v1.0/cloudgateway/bootstrapCG", restapis.GetBootstrapCG)
	router.POST("/siebel/v1.0/cloudgateway/bootstrapCG", restapis.PostBootstrapCG)

	// List APIs
	router.GET("/siebel/v1.0/appicon", restapis.GetAppIcon)
	router.GET("/siebel/v1.0/cloudgateway/discovery/services", restapis.GetDiscoveryServices)

	// Security Profile APIs
	router.POST("/siebel/v1.0/cloudgateway/profiles/security", restapis.PostSecurityProfile)
	router.GET("/siebel/v1.0/cloudgateway/profiles/security", restapis.GetSecurityProfiles)
	router.GET("/siebel/v1.0/cloudgateway/profiles/security/:profilename", restapis.GetSecurityProfile)
	router.PUT("/siebel/v1.0/cloudgateway/profiles/security/:profilename", restapis.PutSecurityProfile)
	router.DELETE("/siebel/v1.0/cloudgateway/profiles/security/:profilename", restapis.DeleteSecurityProfile)

	// Enterprise Profile APIs
	router.POST("/siebel/v1.0/cloudgateway/profiles/enterprises", restapis.PostEnterpriseProfile)
	router.GET("/siebel/v1.0/cloudgateway/profiles/enterprises", restapis.GetEnterpriseProfiles)
	router.GET("/siebel/v1.0/cloudgateway/profiles/enterprises/:profilename", restapis.GetEnterpriseProfile)
	router.PUT("/siebel/v1.0/cloudgateway/profiles/enterprises/:profilename", restapis.PutEnterpriseProfile)
	router.DELETE("/siebel/v1.0/cloudgateway/profiles/enterprises/:profilename", restapis.DeleteEnterpriseProfile)

	// Enterprise Deployment APIs
	router.POST("/siebel/v1.0/cloudgateway/deployments/enterprises", restapis.PostEnterpriseDeployment)
	router.GET("/siebel/v1.0/cloudgateway/deployments/enterprises", restapis.GetEnterpriseDeployments)
	router.GET("/siebel/v1.0/cloudgateway/deployments/enterprises/:deploymentname", restapis.GetEnterpriseDeployment)
	router.PUT("/siebel/v1.0/cloudgateway/deployments/enterprises/:deploymentname", restapis.PutEnterpriseDeployment)
	router.DELETE("/siebel/v1.0/cloudgateway/deployments/enterprises/:deploymentname", restapis.DeleteEnterpriseDeployment)

	// Server Profile APIs
	router.POST("/siebel/v1.0/cloudgateway/profiles/servers", restapis.PostServerProfile)
	router.GET("/siebel/v1.0/cloudgateway/profiles/servers", restapis.GetServerProfiles)
	router.GET("/siebel/v1.0/cloudgateway/profiles/servers/:profilename", restapis.GetServerProfile)
	router.PUT("/siebel/v1.0/cloudgateway/profiles/servers/:profilename", restapis.PutServerProfile)
	router.DELETE("/siebel/v1.0/cloudgateway/profiles/servers/:profilename", restapis.DeleteServerProfile)

	// Server Deployment APIs
	router.POST("/siebel/v1.0/cloudgateway/deployments/servers", restapis.PostServerDeployment)
	router.GET("/siebel/v1.0/cloudgateway/deployments/servers", restapis.GetServerDeployments)
	router.GET("/siebel/v1.0/cloudgateway/deployments/servers/:deploymentname", restapis.GetServerDeployment)
	router.PUT("/siebel/v1.0/cloudgateway/deployments/servers/:deploymentname", restapis.PutServerDeployment)
	router.DELETE("/siebel/v1.0/cloudgateway/deployments/servers/:deploymentname", restapis.DeleteServerDeployment)

	// SWSM Profile APIs
	router.POST("/siebel/v1.0/cloudgateway/profiles/swsm", restapis.PostSWSMProfile)
	router.GET("/siebel/v1.0/cloudgateway/profiles/swsm", restapis.GetSWSMProfiles)
	router.GET("/siebel/v1.0/cloudgateway/profiles/swsm/:profilename", restapis.GetSWSMProfile)
	router.PUT("/siebel/v1.0/cloudgateway/profiles/swsm/:profilename", restapis.PutSWSMProfile)
	router.DELETE("/siebel/v1.0/cloudgateway/profiles/swsm/:profilename", restapis.DeleteSWSMProfile)

	// SWSM Deployment APIs
	router.POST("/siebel/v1.0/cloudgateway/deployments/swsm", restapis.PostSWSMDeployment)
	router.GET("/siebel/v1.0/cloudgateway/deployments/swsm", restapis.GetSWSMDeployments)
	router.GET("/siebel/v1.0/cloudgateway/deployments/swsm/:deploymentname", restapis.GetSWSMDeployment)
	router.PUT("/siebel/v1.0/cloudgateway/deployments/swsm/:deploymentname", restapis.PutSWSMDeployment)
	router.DELETE("/siebel/v1.0/cloudgateway/deployments/swsm/:deploymentname", restapis.DeleteSWSMDeployment)

	// CacheServer Profile APIs
	router.POST("/siebel/v1.0/cloudgateway/profiles/cacheserver", restapis.PostCacheServerProfile)
	router.GET("/siebel/v1.0/cloudgateway/profiles/cacheserver", restapis.GetCacheServerProfiles)
	router.GET("/siebel/v1.0/cloudgateway/profiles/cacheserver/:profilename", restapis.GetCacheServerProfile)
	router.PUT("/siebel/v1.0/cloudgateway/profiles/cacheserver/:profilename", restapis.PutCacheServerProfile)
	router.DELETE("/siebel/v1.0/cloudgateway/profiles/cacheserver/:profilename", restapis.DeleteCacheServerProfile)

	// CacheServer Deployment APIs
	router.POST("/siebel/v1.0/cloudgateway/deployments/cacheserver", restapis.PostCacheServerDeployment)
	router.GET("/siebel/v1.0/cloudgateway/deployments/cacheserver", restapis.GetCacheServerDeployments)
	router.GET("/siebel/v1.0/cloudgateway/deployments/cacheserver/:profilename", restapis.GetCacheServerDeployment)
	router.PUT("/siebel/v1.0/cloudgateway/deployments/cacheserver/:profilename", restapis.PutCacheServerDeployment)
	router.DELETE("/siebel/v1.0/cloudgateway/deployments/cacheserver/:profilename", restapis.DeleteCacheServerDeployment)

	// CacheClient Profile APIs
	router.POST("/siebel/v1.0/cloudgateway/profiles/cacheclient", restapis.PostCacheClientProfile)
	router.GET("/siebel/v1.0/cloudgateway/profiles/cacheclient", restapis.GetCacheClientProfiles)
	router.GET("/siebel/v1.0/cloudgateway/profiles/cacheclient/:profilename", restapis.GetCacheClientProfile)
	router.PUT("/siebel/v1.0/cloudgateway/profiles/cacheclient/:profilename", restapis.PutCacheClientProfile)
	router.DELETE("/siebel/v1.0/cloudgateway/profiles/cacheclient/:profilename", restapis.DeleteCacheClientProfile)

	// Migration Profile APIs
	router.POST("/siebel/v1.0/cloudgateway/profiles/migrations", restapis.PostMigrationProfile)
	router.GET("/siebel/v1.0/cloudgateway/profiles/migrations", restapis.GetMigrationProfiles)
	router.GET("/siebel/v1.0/cloudgateway/profiles/migrations/:profilename", restapis.GetMigrationProfile)
	router.PUT("/siebel/v1.0/cloudgateway/profiles/migrations/:profilename", restapis.PutMigrationProfile)
	router.DELETE("/siebel/v1.0/cloudgateway/profiles/migrations/:profilename", restapis.DeleteMigrationProfile)

	// Migration Deployment APIs
	router.POST("/siebel/v1.0/cloudgateway/deployments/migrations", restapis.PostMigrationDeployment)
	router.GET("/siebel/v1.0/cloudgateway/deployments/migrations", restapis.GetMigrationDeployments)
	router.GET("/siebel/v1.0/cloudgateway/deployments/migrations/:profilename", restapis.GetMigrationDeployment)
	router.PUT("/siebel/v1.0/cloudgateway/deployments/migrations/:profilename", restapis.PutMigrationDeployment)
	router.DELETE("/siebel/v1.0/cloudgateway/deployments/migrations/:profilename", restapis.DeleteMigrationDeployment)

	fmt.Printf("RESTful API URL - http://localhost" + port + "/siebel/v1.0/\n")
	log.Fatal(http.ListenAndServe(port, router))
}
