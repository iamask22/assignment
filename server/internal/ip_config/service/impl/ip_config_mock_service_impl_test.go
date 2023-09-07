package impl

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
	"mta-hosting-optimizer/server/internal/ip_config/service"
	"mta-hosting-optimizer/server/internal/ip_config/service/data_generator"
	"mta-hosting-optimizer/server/internal/ip_config/service/dtos"
	"testing"
)

type IPConfigMockServiceImplTestSuite struct {
	suite.Suite
	controller          *gomock.Controller
	ipConfigMockService service.IPConfigMockService
}

func (suite *IPConfigMockServiceImplTestSuite) SetupTest() {
}

func (suite *IPConfigMockServiceImplTestSuite) BeforeTest(string, string) {
	suite.controller = gomock.NewController(suite.T())
	suite.ipConfigMockService = NewIPConfigMockService()
}

func (suite *IPConfigMockServiceImplTestSuite) AfterTest(string, string) {
	suite.controller.Finish()
}

func TestIPConfigMockServiceImplTestSuite(t *testing.T) {
	suite.Run(t, new(IPConfigMockServiceImplTestSuite))
}

func (suite *IPConfigMockServiceImplTestSuite) TestGetIPConfigData_success() {
	ipConfigData := data_generator.IPConfigMockData

	actualIPConfigData := suite.ipConfigMockService.GetIPConfigData()
	suite.Equal(ipConfigData, actualIPConfigData)
}

func (suite *IPConfigMockServiceImplTestSuite) TestGetIPConfigData_failure() {
	var ipConfigData []dtos.IpConfig

	actualIPConfigData := suite.ipConfigMockService.GetIPConfigData()
	suite.NotEqual(ipConfigData, actualIPConfigData)
}
