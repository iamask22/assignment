package service

import "mta-hosting-optimizer/server/internal/ip_config/service/dtos"

type IPConfigMockService interface {
	GetIPConfigData() []dtos.IpConfig
}

//go:generate mockgen -source=ip_config_mock_service.go -destination=./mocks/ip_config_mock_service.go -package=mocks
