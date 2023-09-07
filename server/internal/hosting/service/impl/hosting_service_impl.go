package impl

import (
	"mta-hosting-optimizer/server/internal/hosting/service"
	service2 "mta-hosting-optimizer/server/internal/ip_config/service"
	"sync"
)

type hostingServiceImpl struct {
	threshold int
	sync.RWMutex
	iPConfigMockService service2.IPConfigMockService
}

func NewHostingService(threshold int, iPConfigMockService service2.IPConfigMockService) service.HostingService {
	return &hostingServiceImpl{
		threshold:           threshold,
		iPConfigMockService: iPConfigMockService,
	}
}

func (svc *hostingServiceImpl) GetHostNames() []string {
	ipConfigData := svc.iPConfigMockService.GetIPConfigData()

	hostnames := make(map[string]int)
	for _, ipConfig := range ipConfigData {
		if ipConfig.Active {
			hostnames[ipConfig.Hostname]++
		} else {
			if hostnames[ipConfig.Hostname] == 0 {
				// If the hostname is not in the map, add it with a value of 0.
				hostnames[ipConfig.Hostname] = 0
			}
		}
	}

	var result []string
	for hostname, count := range hostnames {
		if count <= svc.threshold {
			result = append(result, hostname)
		}
	}

	return result
}
