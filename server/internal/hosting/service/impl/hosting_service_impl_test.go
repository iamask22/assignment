package impl

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
	hostingService "mta-hosting-optimizer/server/internal/hosting/service"
	"mta-hosting-optimizer/server/internal/ip_config/service/dtos"
	"mta-hosting-optimizer/server/internal/ip_config/service/mocks"
	"sort"
	"testing"
)

type HostingServiceImplTestSuite struct {
	suite.Suite
	controller          *gomock.Controller
	hostingService      hostingService.HostingService
	ipConfigMockService *mocks.MockIPConfigMockService
}

func (suite *HostingServiceImplTestSuite) SetupTest() {
}

func (suite *HostingServiceImplTestSuite) BeforeTest(string, string) {
	suite.controller = gomock.NewController(suite.T())
	suite.ipConfigMockService = mocks.NewMockIPConfigMockService(suite.controller)
	suite.hostingService = NewHostingService(1, suite.ipConfigMockService)
}

func (suite *HostingServiceImplTestSuite) AfterTest(string, string) {
	suite.controller.Finish()
}

func TestHostingServiceImplTestSuite(t *testing.T) {
	suite.Run(t, new(HostingServiceImplTestSuite))
}

func (suite *HostingServiceImplTestSuite) TestGetHostNames_success() {

	ipConfigData := []dtos.IpConfig{
		{IP: "192.168.1.1", HostName: "server1", Active: true},
		{IP: "192.168.1.2", HostName: "server2", Active: true},
		{IP: "192.168.1.3", HostName: "server3", Active: false},
	}

	testCases := []struct {
		name          string
		x             int
		expectedHosts []string
	}{
		{
			name:          "X=2",
			x:             2,
			expectedHosts: []string{"server1", "server2", "server3"},
		},
		{
			name:          "X=1 (default)",
			x:             1,
			expectedHosts: []string{"server1", "server2", "server3"},
		},
		{
			name:          "X=0",
			x:             0,
			expectedHosts: []string{"server3"},
		},
	}

	for _, tc := range testCases {
		suite.Run(tc.name, func() {
			suite.hostingService = NewHostingService(tc.x, suite.ipConfigMockService)
			suite.ipConfigMockService.EXPECT().GetIPConfigData().Return(ipConfigData)
			actualResponse := suite.hostingService.GetHostNames()
			sort.Strings(actualResponse)
			sort.Strings(tc.expectedHosts)
			suite.Equal(tc.expectedHosts, actualResponse)
		})
	}
}

func (suite *HostingServiceImplTestSuite) TestGetHostNames_nil() {

	testCases := []struct {
		name          string
		x             int
		expectedHosts []string
	}{
		{
			name:          "X=2",
			x:             2,
			expectedHosts: []string(nil),
		},
		{
			name:          "X=1 (default)",
			x:             1,
			expectedHosts: []string(nil),
		},
		{
			name:          "X=0",
			x:             0,
			expectedHosts: []string(nil),
		},
	}

	for _, tc := range testCases {
		suite.Run(tc.name, func() {
			suite.hostingService = NewHostingService(tc.x, suite.ipConfigMockService)
			suite.ipConfigMockService.EXPECT().GetIPConfigData().Return(nil)
			actualResponse := suite.hostingService.GetHostNames()
			sort.Strings(actualResponse)
			sort.Strings(tc.expectedHosts)
			suite.Equal(tc.expectedHosts, actualResponse)
		})
	}
}
