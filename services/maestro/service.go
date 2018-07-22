package maestro

import (
	"conductor/models"
	"conductor/providers"
	"errors"
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

// Maestro submits job to key/value store.
type Maestro interface {
	LoadJob(string, string) (models.Job, error)
	Deploy(string, models.Job) (string, error)
	Undeploy(string, models.Job) (bool, error)
}

type maestro struct{}

func (maestro) LoadJob(location string, sourceType string) (models.Job, error) {
	if location == "" {
		return models.Job{}, errors.New("location was empty")
	}
	if sourceType != "file" && sourceType != "url" {
		return models.Job{}, errors.New("source should be \"file\" or \"url\", but got: " + sourceType)
	}
	data, err := ioutil.ReadFile(location)
	if err != nil {
		return models.Job{}, err
	}

	job := models.Job{}
	err = yaml.Unmarshal([]byte(data), &job)
	if err != nil {
		return models.Job{}, err
	}

	return job, nil
}

func (maestro) Deploy(jobName string, job models.Job) (string, error) {
	if jobName == "" {
		return "", errors.New("Pod is empty string")
	}
	jobYaml, err := yaml.Marshal(&job)
	if err != nil {
		return "", err
	}

	etcd := providers.NewEtcd([]string{"localhost:2379"})
	jobContent, err := etcd.SubmitJob(jobName, string(jobYaml))
	if err != nil {
		return "", err
	}

	return jobContent, nil
}

func (maestro) Undeploy(pod string, job models.Job) (bool, error) {
	if pod == "" {
		return false, errors.New("Pod is empty string")
	}
	return true, nil
}

// NewService placeholder for dependencies.
func NewService() Maestro {
	return &maestro{}
}
