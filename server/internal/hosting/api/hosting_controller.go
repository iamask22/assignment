package api

import (
	"github.com/gorilla/mux"
	"mta-hosting-optimizer/server/api_handler"
	"mta-hosting-optimizer/server/internal/hosting/service"
	"net/http"
)

type HostingController struct {
	HostingService service.HostingService
}

func NewHostingController(apiMux *mux.Router, hostingService service.HostingService) *HostingController {
	controller := &HostingController{
		HostingService: hostingService,
	}

	v1Mux := apiMux.PathPrefix(HostingControllerApiPathPrefixV1).Subrouter()

	v1Mux.HandleFunc("/hostnames", controller.GetHostNames).Methods(http.MethodGet)

	return controller
}

func (hc *HostingController) GetHostNames(w http.ResponseWriter, _ *http.Request) {
	hostNames := hc.HostingService.GetHostNames()

	response := api_handler.SuccessResponse{
		Status: http.StatusOK,
		Data:   hostNames,
	}

	response.Write(w)
}
