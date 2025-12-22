package models

type Info struct {
	Version              string   `json:"version"`
	CriteriaVersion      string   `json:"criteriaVersion"`
	MaxAssessments       int      `json:"maxAssessments"`
	CurrentAssessments   int      `json:"currentAssessments"`
	NewAssessmentCoolOff int64    `json:"NewAssessmentCoolOff"`
	Messages             []string `json:"messages"`
}
