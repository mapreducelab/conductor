package maestro

import (
	"conductor/models"
	"errors"
	"fmt"
	"io/ioutil"
	"log"

	yaml "gopkg.in/yaml.v2"
)

// Maestro submits job to key/value store.
type Maestro interface {
	LoadJob(string, string) (string, error)
	Deploy(string, string) (bool, error)
	Undeploy(string, string) (bool, error)
}

type maestro struct{}

func (maestro) LoadJob(location string, sourceType string) (string, error) {
	if location == "" {
		return "", errors.New("location was empty")
	}
	if sourceType != "file" && sourceType != "url" {
		return "", errors.New("source should be \"file\" or \"url\", but got: " + sourceType)
	}

	data, err := ioutil.ReadFile(location)
	check(err)

	job := models.Job{}

	err = yaml.Unmarshal([]byte(data), &job)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("%+v\n", job)

	return "test", nil
}

func (maestro) Deploy(region string, job string) (bool, error) {
	if region == "" {
		return false, errors.New("Pod is empty string")
	}
	fmt.Printf("Region: %s, File: %s", region, job)
	return true, nil
}

func (maestro) Undeploy(pod string, job string) (bool, error) {
	if pod == "" {
		return false, errors.New("Pod is empty string")
	}
	return true, nil
}

// NewService placeholder for dependencies.
func NewService() Maestro {
	return &maestro{}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
