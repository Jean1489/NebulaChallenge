package models

type Endpoint struct {
	IPAddress            string           `json:"ipAddress"`
	ServerName           string           `json:"serverName,omitempty"`
	StatusMessage        string           `json:"statusMessage"`
	StatusDetails        string           `json:"statusDetails,omitempty"`
	StatusDetailsMessage string           `json:"statusDetailsMessage,omitempty"`
	Grade                string           `json:"grade,omitempty"`
	GradeTrustIgnored    string           `json:"gradeTrustIgnored,omitempty"`
	HasWarnings          bool             `json:"hasWarnings"`
	IsExceptional        bool             `json:"isExceptional"`
	Progress             int              `json:"progress"`
	Duration             int              `json:"duration"`
	ETA                  int              `json:"eta,omitempty"`
	Delegation           int              `json:"delegation,omitempty"`
	Details              *EndpointDetails `json:"details,omitempty"`
}
