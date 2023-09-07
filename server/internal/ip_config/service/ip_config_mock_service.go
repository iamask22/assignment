package service

import "mta-hosting-optimizer/server/internal/ip_config/service/dtos"

type IPConfigMockService interface {
	GetIPConfigData() []dtos.IpConfig
}
