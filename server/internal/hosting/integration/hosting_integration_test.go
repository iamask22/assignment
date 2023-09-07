package integration

import (
	"fmt"
	"github.com/gorilla/mux"
	"mta-hosting-optimizer/server/config/webserver"
	hostingAPI "mta-hosting-optimizer/server/internal/hosting/api"
	hostingServiceImpl "mta-hosting-optimizer/server/internal/hosting/service/impl"
	ipConfigServiceImpl "mta-hosting-optimizer/server/internal/ip_config/service/impl"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"strings"
	"testing"
)

func TestGetHostNames(t *testing.T) {
	expected := `{"status":200,"data":["mta-prod-1","mta-prod-2","mta-prod-3"]}`

	err := os.Setenv(webserver.ThresholdEnvVar, "2")
	if err != nil {
		panic(fmt.Sprintf("Invalid threshold value: %s", err.Error()))
	}
	defer os.Unsetenv(webserver.ThresholdEnvVar)

	threshold, err := strconv.Atoi(os.Getenv(webserver.ThresholdEnvVar))
	if err != nil {
		t.Errorf("Handler returned unexpected error: got %v want %v", err.Error(), expected)
	}

	router := mux.NewRouter()

	hostingAPI.NewHostingController(router,
		hostingServiceImpl.NewHostingService(threshold, ipConfigServiceImpl.NewIPConfigMockService()))

	req, err := http.NewRequest(http.MethodGet, GetHostNamesAPIPath, nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
	rrBodyStr := strings.TrimSpace(rr.Body.String())

	if rrBodyStr != expected {
		t.Errorf("Handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

func TestGetHostNamesDefaultThreshold(t *testing.T) {
	expected := `{"status":200,"data":["mta-prod-1","mta-prod-3"]}`

	// Set the environment variable X to a specific value for testing
	err := os.Setenv(webserver.ThresholdEnvVar, "1")
	if err != nil {
		panic(fmt.Sprintf("Invalid threshold value: %s", err.Error()))
	}
	defer os.Unsetenv(webserver.ThresholdEnvVar)

	threshold, err := strconv.Atoi(os.Getenv(webserver.ThresholdEnvVar))
	if err != nil {
		t.Errorf("Handler returned unexpected error: got %v want %v", err.Error(), expected)
	}

	router := mux.NewRouter()

	hostingAPI.NewHostingController(router,
		hostingServiceImpl.NewHostingService(threshold, ipConfigServiceImpl.NewIPConfigMockService()))

	req, err := http.NewRequest(http.MethodGet, GetHostNamesAPIPath, nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
	arrBodyInStr := strings.TrimSpace(rr.Body.String())

	if arrBodyInStr != expected {
		t.Errorf("Handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}
