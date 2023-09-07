package api

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"mta-hosting-optimizer/server/internal/hosting/service"
	"net/http"
)

type HostingController struct {
	hostingService service.HostingService
}

func NewHostingController(apiMux *mux.Router, hostingService service.HostingService) *HostingController {
	controller := &HostingController{
		hostingService: hostingService,
	}

	v1Mux := apiMux.PathPrefix(HostingControllerApiPathPrefixV1).Subrouter()

	v1Mux.HandleFunc("/hostnames", controller.getHostNames).Methods(http.MethodGet)

	return controller
}

func (hc *HostingController) getHostNames(w http.ResponseWriter, _ *http.Request) {
	hostNames := hc.hostingService.GetHostNames()
	outputJsonResp := make(map[string]interface{})
	outputJsonResp["output"] = hostNames

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(outputJsonResp)
	if err != nil {
		log.Printf("Error while encoding response: %v", err)
		return
	}
}
