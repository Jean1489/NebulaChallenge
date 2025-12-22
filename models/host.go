package models

type Host struct {
	Host            string     `json:"host"`
	Port            int        `json:"port"`
	Protocol        string     `json:"protocol"`
	IsPublic        bool       `json:"isPublic"`
	Status          string     `json:"status"`
	StatusMessage   string     `json:"statusMessage,omitempty"`
	StartTime       int64      `json:"startTime"`
	TestTime        int64      `json:"testTime,omitempty"`
	EngineVersion   string     `json:"engineVersion,omitempty"`
	CriteriaVersion string     `json:"criteriaVersion,omitempty"`
	CacheExpiryTime int64      `json:"cacheExpiryTime,omitempty"`
	Endpoints       []Endpoint `json:"endPoints"`
	CertHostnames   []string   `json:"certHostnames,omitempty"`
}
