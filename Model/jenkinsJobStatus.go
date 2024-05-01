package Model

import (
	"gorm.io/gorm"
)

type JenkinsJobStatus struct {
	gorm.Model
	JobName        string `json:"jenkins_job_name"`
	ProjectName    string `json:"project_name"`
	FeatureName    string `json:"feature_name"`
	Status         bool   `json:"job_status"`
	BuildNumber    string `json:"build_number"`
	ChildJob       string `json:"child_job_list"`
	JobParamsKey   string `json:"job_params_key"`
	JobParamsValue string `json:"job_params_value"`
}
