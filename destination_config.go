package fivetran

type DestinationConfig struct {
	Fhost     string `json:"host,omitempty"`
	Fport     int    `json:"port,omitempty"`
	Fdatabase string `json:"database,omitempty"`
	Fauth     string `json:"auth,omitempty"`
	Fuser     string `json:"user,omitempty"`
	Fpassword string `json:"password,omitempty"`
}

// DestinationConfigTemp is a temporary type used to accomodate the field Fport as type string instead of type int.
// The field should be int but the REST API is returning string.
// When https://fivetran.height.app/T-97508 is fixed, this temporary type should be removed.
type DestinationConfigTemp struct {
	Fhost     string `json:"host,omitempty"`
	Fport     string `json:"port,omitempty"` // Can use interface{} ? to check...
	Fdatabase string `json:"database,omitempty"`
	Fauth     string `json:"auth,omitempty"`
	Fuser     string `json:"user,omitempty"`
	Fpassword string `json:"password,omitempty"`
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
