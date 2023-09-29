package teams

type TeamData struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Role        string `json:"role"`
}

type TeamMembership struct {
	Role      string `json:"role"`
	CreatedAt string `json:"created_at"`
}

type TeamConnectorMembership struct {
	ConnectorId string `json:"id"`
	TeamMembership
}

type TeamGroupMembership struct {
	GroupId string `json:"id"`
	TeamMembership
}

type TeamUserMembership struct {
	UserId string `json:"user_id"`
	TeamMembership
}
