package model

type User struct {
	Id         string   `json:"_id,"`
	FirstName  string   `json:"firstName"`
	LastName   string   `json:"lastName"`
	Email      string   `json:"email"`
	Phone      string   `json:"phone"`
	AdProvider string   `json:"adProvider"`
	Access     []Access `json:"access"`
}

type Access struct {
	OrgId    string   `json:"orgId"`
	CenterId string   `json:"centerId"`
	Roles    []string `json:"roles"`
}
