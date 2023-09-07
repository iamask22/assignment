package api

import (
	"encoding/json"
	"github.com/gorilla/mux"
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

func (hc *HostingController) getHostNames(w http.ResponseWriter, r *http.Request) {
	resp := hc.hostingService.GetHostNames()
	outputResp := make(map[string]interface{})
	outputResp["output"] = resp

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(outputResp)
	if err != nil {
		return
	}
}
