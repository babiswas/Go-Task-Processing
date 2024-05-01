package JenkinsUtil

import (
	"context"
	"fmt"
	"os"

	"github.com/bndr/gojenkins"
)

type JenkinsObject struct {
	JenkinsURL    string
	AdminUserName string
	AdminPassword string
	Jobname       string
	Params        map[string]string
}

func GetJenkinsObject() JenkinsObject {
	jnkn := JenkinsObject{JenkinsURL: os.Getenv("JENKINS_URL"), AdminUserName: os.Getenv("JENKINS_USERNAME"), AdminPassword: os.Getenv("JENKINS_PASSWORD"), Jobname: "", Params: make(map[string]string)}
	return jnkn
}

func (jnkn *JenkinsObject) GetJenkinsContext() (*gojenkins.Jenkins, context.Context) {
	ctx := context.Background()
	jenkins := gojenkins.CreateJenkins(nil, jnkn.JenkinsURL, jnkn.AdminUserName, jnkn.AdminPassword)
	_, err := jenkins.Init(ctx)

	if err != nil {
		fmt.Println(err)
		panic("Unable to connect to remote jenkins server.")
	}
	return jenkins, ctx

}

func (jnkn *JenkinsObject) SetJobName(JobName string) {
	jnkn.Jobname = JobName
}

func (jnkn *JenkinsObject) SetJobParams(mp map[string]string) {
	for k := range jnkn.Params {
		delete(jnkn.Params, k)
	}
	for key, value := range mp {
		jnkn.Params[key] = value
	}
}

func (jnkn *JenkinsObject) TriggerBuild(ctx context.Context, jenkins *gojenkins.Jenkins) int64 {
	queueid, err := jenkins.BuildJob(ctx, jnkn.Jobname, jnkn.Params)
	if err != nil {
		panic(err)
	}
	build, err := jenkins.GetBuildFromQueueID(ctx, queueid)
	if err != nil {
		fmt.Println(err)
	}
	return build.GetBuildNumber()
}

func (jnkn *JenkinsObject) TriggerJenkinsJob(jobName string, mp map[string]string) int64 {
	jenkin, ctx := jnkn.GetJenkinsContext()
	jnkn.SetJobName(jobName)
	jnkn.SetJobParams(mp)
	build_number := jnkn.TriggerBuild(ctx, jenkin)
	return build_number
}
