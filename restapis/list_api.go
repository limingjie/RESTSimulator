package restapis

import (
	"fmt"
	"net/http"

	"../logger"
	"github.com/julienschmidt/httprouter"
)

// GetAppIcon - GET /AppIcon
func GetAppIcon(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	logger.Logger("GetAppIcon", "icons")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "{\"AppIcon\":[{\"name\":\"sales.png\"},{\"name\":\"sales1.png\"},{\"name\":\"sales2.png\"},{\"name\":\"sales3.png\"},{\"name\":\"sales4.png\"},{\"name\":\"service.png\"},{\"name\":\"service2.png\"},{\"name\":\"callcenter.png\"},{\"name\":\"LifeSciences1.png\"},{\"name\":\"LifeSciences3.png\"},{\"name\":\"LifeSciences4.png\"},{\"name\":\"ConsumerGoods.png\"},{\"name\":\"epharma.png\"},{\"name\":\"fins.png\"}]}")
}

// GetDiscoveryServices - GET /discovery/services
func GetDiscoveryServices(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	logger.Logger("GetDiscoveryServices", "services")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "{\"Service\":[{\"displayName\":\"EAIOutboundServer\",\"id\":\"infraeaioutbound\",\"lang\":\"ENU\",\"load\":0,\"capacity\":20,\"connect_string\":\"/siebel/v1.0/cloudgateway/discovery/services/infraeaioutbound/connectstring\"},{\"displayName\":\"CustomApplicationObjectManager(ENU)\",\"id\":\"customappobjmgr_enu\",\"lang\":\"ENU\",\"load\":0,\"capacity\":20,\"connect_string\":\"/siebel/v1.0/cloudgateway/discovery/services/customappobjmgr_enu/connectstring\"},{\"displayName\":\"CallCenterObjectManager(ENU)\",\"id\":\"sccobjmgr_enu\",\"lang\":\"ENU\",\"load\":0,\"capacity\":20,\"connect_string\":\"/siebel/v1.0/cloudgateway/discovery/services/sccobjmgr_enu/connectstring\"},{\"displayName\":\"EAIObjectManager(ENU)\",\"id\":\"eaiobjmgr_enu\",\"lang\":\"ENU\",\"load\":0,\"capacity\":20,\"connect_string\":\"/siebel/v1.0/cloudgateway/discovery/services/eaiobjmgr_enu/connectstring\"},{\"displayName\":\"ServerRequestProcessor\",\"id\":\"srproc\",\"lang\":\"ENU\",\"load\":2,\"capacity\":20,\"connect_string\":\"/siebel/v1.0/cloudgateway/discovery/services/srproc/connectstring\"},{\"displayName\":\"BusinessIntegrationManager\",\"id\":\"busintmgr\",\"lang\":\"ENU\",\"load\":0,\"capacity\":20,\"connect_string\":\"/siebel/v1.0/cloudgateway/discovery/services/busintmgr/connectstring\"},{\"displayName\":\"SiebelAdministratorNotificationComponent\",\"id\":\"adminnotify\",\"lang\":\"ENU\",\"load\":0,\"capacity\":10,\"connect_string\":\"/siebel/v1.0/cloudgateway/discovery/services/adminnotify/connectstring\"},{\"displayName\":\"SelfServiceObjectManager(ENU)\",\"id\":\"sserviceobjmgr_enu\",\"lang\":\"ENU\",\"load\":0,\"capacity\":20,\"connect_string\":\"/siebel/v1.0/cloudgateway/discovery/services/sserviceobjmgr_enu/connectstring\"},{\"displayName\":\"FileSystemManager\",\"id\":\"fsmsrvr\",\"lang\":\"ENU\",\"load\":0,\"capacity\":20,\"connect_string\":\"/siebel/v1.0/cloudgateway/discovery/services/fsmsrvr/connectstring\"},{\"displayName\":\"eServiceObjectManager(ENU)\",\"id\":\"eserviceobjmgr_enu\",\"lang\":\"ENU\",\"load\":0,\"capacity\":20,\"connect_string\":\"/siebel/v1.0/cloudgateway/discovery/services/eserviceobjmgr_enu/connectstring\"},{\"displayName\":\"BusinessIntegrationBatchManager\",\"id\":\"busintbatchmgr\",\"lang\":\"ENU\",\"load\":0,\"capacity\":20,\"connect_string\":\"/siebel/v1.0/cloudgateway/discovery/services/busintbatchmgr/connectstring\"},{\"displayName\":\"ServerRequestBroker\",\"id\":\"srbroker\",\"lang\":\"ENU\",\"load\":16,\"capacity\":100,\"connect_string\":\"/siebel/v1.0/cloudgateway/discovery/services/srbroker/connectstring\"},{\"displayName\":\"JMSReceiver\",\"id\":\"jmsreceiver\",\"lang\":\"ENU\",\"load\":0,\"capacity\":20,\"connect_string\":\"/siebel/v1.0/cloudgateway/discovery/services/jmsreceiver/connectstring\"}]}")
}
