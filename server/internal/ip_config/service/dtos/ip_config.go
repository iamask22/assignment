package dtos

// IpConfig represents the data model for IP configurations.
type IpConfig struct {
	IP       string `json:"IP"`
	Hostname string `json:"Hostname"`
	Active   bool   `json:"Active"`
}
