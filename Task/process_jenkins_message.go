package Task

import (
	"log"
	jnkn_util "notification/Common/JenkinsUtil"
	"notification/Database"
	model "notification/Model"
	"strings"
)

func Process_ET_Jenkins_Notification(JobName, ProjectName, FeatureName, BuildNumber string, Status string, childJob string, jobParamKeys string, jobParamVals string) (string, error) {
	var job_status bool

	log.Println("Executing task.")

	if Status == "true" {
		job_status = true
	} else if Status == "false" {
		job_status = false
	}

	jenkins_job := model.JenkinsJobStatus{JobName: JobName, ProjectName: ProjectName, FeatureName: FeatureName, Status: job_status, BuildNumber: BuildNumber, ChildJob: childJob, JobParamsKey: jobParamKeys, JobParamsValue: jobParamVals}
	result := Database.DB.Create(&jenkins_job)
	if result.Error != nil {
		log.Println("Error occured :", result.Error)
		return "FAILURE", result.Error
	}
	if job_status {
		jnkn_obj := jnkn_util.GetJenkinsObject()
		job_params := make(map[string]string)
		jparams := strings.Split(jobParamKeys, ",")
		jvalues := strings.Split(jobParamVals, ",")
		for index, value := range jparams {
			job_params[value] = jvalues[index]
		}
		jnkn_obj.TriggerJenkinsJob(childJob, job_params)
	} else {
		return "REMOTE_JOB_FAILURE", nil
	}
	return "SUCCESS", nil
}
