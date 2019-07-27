package models

// Playbook represent the playbook model
type Playbook struct {
	ID          int        `json:"id"`
	Name        string     `json:"name"`
	AddBy       string     `json:"addBy"`
	Description NullString `json:"description"`
	Location    string     `json:"location"`
}
