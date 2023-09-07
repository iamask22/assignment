package dtos

// IpConfig represents the data model for IP configurations.
type IpConfig struct {
	IP       string `json:"ip"`
	HostName string `json:"host_name"`
	Active   bool   `json:"active"`
}
