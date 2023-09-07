package webserver

import (
	"github.com/gorilla/mux"
	"log"
	hostingAPI "mta-hosting-optimizer/server/internal/hosting/api"
	hostingServiceImpl "mta-hosting-optimizer/server/internal/hosting/service/impl"
	ipConfigServiceImpl "mta-hosting-optimizer/server/internal/ip_config/service/impl"
	"os"
	"strconv"
)

func InitializeApplicationConfig(apiMux *mux.Router) {
	// Get the threshold value from the environment variable, defaulting to 1.
	thresholdStr := os.Getenv("THRESHOLD")
	if thresholdStr == "" {
		thresholdStr = "1"
	}

	var err error
	threshold, err := strconv.Atoi(thresholdStr)
	if err != nil {
		log.Fatalf("Invalid threshold value: %v", err)
	}

	iPConfigMockService := ipConfigServiceImpl.NewIPConfigMockService()

	hostingServiceImpl := hostingServiceImpl.NewHostingService(threshold, iPConfigMockService)
	hostingAPI.NewHostingController(apiMux, hostingServiceImpl)
}
