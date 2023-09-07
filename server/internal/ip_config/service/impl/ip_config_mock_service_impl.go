package impl

import (
	"mta-hosting-optimizer/server/internal/ip_config/service"
	"mta-hosting-optimizer/server/internal/ip_config/service/data_generator"
	"mta-hosting-optimizer/server/internal/ip_config/service/dtos"
	"sync"
)

type ipConfigMockServiceImpl struct {
	sync.RWMutex
}

// NewIPConfigMockService creates a new instance of IPConfigMockService,
// Factory pattern is used to create instances of this service.
func NewIPConfigMockService() service.IPConfigMockService {
	return &ipConfigMockServiceImpl{}
}

func (svc *ipConfigMockServiceImpl) GetIPConfigData() []dtos.IpConfig {
	// Acquiring a write lock here, to guarantee data consistency in the presence of concurrent reads and writes.
	svc.Lock()
	defer svc.Unlock()

	ipConfigData := data_generator.IPConfigMockData
	return ipConfigData
}
