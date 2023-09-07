package webserver

import (
	"fmt"
	"github.com/gorilla/mux"
	hostingAPI "mta-hosting-optimizer/server/internal/hosting/api"
	hostingServiceImpl "mta-hosting-optimizer/server/internal/hosting/service/impl"
	ipConfigServiceImpl "mta-hosting-optimizer/server/internal/ip_config/service/impl"
	"os"
	"strconv"
)

func InitializeApplicationConfig(apiMux *mux.Router) {
	// Get the threshold value from the environment variable, defaulting to 1.
	thresholdStr := os.Getenv(ThresholdEnvVar)
	if thresholdStr == "" {
		thresholdStr = "1"
	}

	var err error
	threshold, err := strconv.Atoi(thresholdStr)
	if err != nil {
		panic(fmt.Sprintf("Invalid threshold value: %s", err.Error()))
	}

	iPConfigMockService := ipConfigServiceImpl.NewIPConfigMockService()

	hostingServImpl := hostingServiceImpl.NewHostingService(threshold, iPConfigMockService)
	// Abstract the hosting service behind an interface, so that it can be easily mocked in unit tests.
	// Abstract factory pattern is used to create the hosting controller.
	hostingAPI.NewHostingController(apiMux, hostingServImpl)
}
