package service

type HostingService interface {
	GetHostNames() []string
}

//go:generate mockgen -source=hosting_service.go -destination=./mocks/hosting_service.go -package=mocks
