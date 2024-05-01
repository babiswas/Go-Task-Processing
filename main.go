package main

import (
	"log"
	dbutil "notification/Database"
	"os"

	machinery "github.com/RichardKnop/machinery/v1"

	workerJob "notification/Task"

	config "github.com/RichardKnop/machinery/v1/config"
)

func init() {
	dbutil.LoadENVVar()
	dbutil.ConnectDB()
	dbutil.SyncDatabase()
}

func main() {
	var cnf = config.Config{
		Broker:        os.Getenv("REDIS_SERVER_URL"),
		ResultBackend: os.Getenv("REDIS_SERVER_URL"),
	}

	server, err := machinery.NewServer(&cnf)
	if err != nil {
		log.Println("Error occured")
		os.Exit(1)
	}

	server.RegisterTask("process_et_jenkins_notification", workerJob.Process_ET_Jenkins_Notification)

	worker := server.NewWorker("worker-1", 1)
	err = worker.Launch()
	if err != nil {
		log.Println("Error occured.")
		os.Exit(1)
	}

}
