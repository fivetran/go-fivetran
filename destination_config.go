package fivetran

type DestinationConfig struct {
	Fhost     string      `json:"host,omitempty"`
	Fport     interface{} `json:"port,omitempty"` // Type should change to int when https://fivetran.height.app/T-97508 fixed.
	Fdatabase string      `json:"database,omitempty"`
	Fauth     string      `json:"auth,omitempty"`
	Fuser     string      `json:"user,omitempty"`
	Fpassword string      `json:"password,omitempty"`
}

func NewDestinationConfig() *DestinationConfig {
	return &DestinationConfig{}
}

func (dc *DestinationConfig) Host(host string) *DestinationConfig {
	dc.Fhost = host
	return dc
}

func (dc *DestinationConfig) Port(port int) *DestinationConfig {
	dc.Fport = port
	return dc
}

func (dc *DestinationConfig) Database(database string) *DestinationConfig {
	dc.Fdatabase = database
	return dc
}

func (dc *DestinationConfig) Auth(auth string) *DestinationConfig {
	dc.Fauth = auth
	return dc
}

func (dc *DestinationConfig) User(user string) *DestinationConfig {
	dc.Fuser = user
	return dc
}

func (dc *DestinationConfig) Password(password string) *DestinationConfig {
	dc.Fpassword = password
	return dc
}
