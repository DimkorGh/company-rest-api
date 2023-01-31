package model

type CompanyEntity struct {
	Id          string `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Amount      int    `json:"amount,omitempty"`
	Registered  bool   `json:"registered,omitempty"`
	Type        string `json:"type,omitempty"`
}
