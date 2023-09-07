package data_generator

import "mta-hosting-optimizer/server/internal/ip_config/service/dtos"

var IPConfigMockData = []dtos.IpConfig{
	{IP: "127.0.0.1", HostName: "mta-prod-1", Active: true},
	{IP: "127.0.0.2", HostName: "mta-prod-1", Active: false},
	{IP: "127.0.0.3", HostName: "mta-prod-2", Active: true},
	{IP: "127.0.0.4", HostName: "mta-prod-2", Active: true},
	{IP: "127.0.0.5", HostName: "mta-prod-2", Active: false},
	{IP: "127.0.0.6", HostName: "mta-prod-3", Active: false},
}
