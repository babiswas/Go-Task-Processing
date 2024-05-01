package Database

import "notification/Model"

func SyncDatabase() {
	DB.AutoMigrate(&Model.JenkinsJobStatus{})
}
