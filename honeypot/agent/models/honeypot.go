package models

type (
	Policy struct {
		Id         string   `json:"Id"`
		WhiteIps   []string `json:"white_ips"`
		WhitePorts []string `json:"white_ports"`
	}

	BackendService struct {
		Id          string `bson:"Id"`
		ServiceName string `json:"service_name"`
		LocalPort   int    `json:"local_port"`
		BackendHost string `json:"backend_host"`
		BackendPort int    `json:"backend_port"`
	}

	PolicyData struct {
		Policy  []Policy         `json:"policy"`
		Service []BackendService `json:"service"`
	}
)
