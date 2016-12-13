package main

import (
	"fmt"
	"log"
	"net/http"

	"./restapis"
	"github.com/julienschmidt/httprouter"
)

// Index - For /
func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}

func main() {
	router := httprouter.New()
	router.GET("/", Index)

	// Enterprise APIs
	router.POST("/profiles/enterprises", restapis.PostEnterprise)
	router.GET("/profiles/enterprises", restapis.GetEnterprises)
	router.GET("/profiles/enterprises/:profilename", restapis.GetEnterprise)
	router.PUT("/profiles/enterprises/:profilename", restapis.PutEnterprise)
	router.DELETE("/profiles/enterprises/:profilename", restapis.DeleteEnterprise)

	// Server APIs
	router.POST("/profiles/servers", restapis.PostServer)
	router.GET("/profiles/servers", restapis.GetServers)
	router.GET("/profiles/servers/:profilename", restapis.GetServer)
	router.PUT("/profiles/servers/:profilename", restapis.PutServer)
	router.DELETE("/profiles/servers/:profilename", restapis.DeleteServer)

	log.Fatal(http.ListenAndServe(":8888", router))
}
