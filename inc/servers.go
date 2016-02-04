package keydropd

import "net"

//Server definition type
type Server struct {
	IPAddress   net.IP `json:"ipAddress,omitempty"`
	SSHPort     int    `json:"sshPort"`
	FQDN        string `json:"fqdn"`
	Description string `json:"description,omitempty"`
	DeployKey   string `json:"deployKey,omitempty"`
}
