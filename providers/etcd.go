package providers

import (
	"context"
	"time"

	"github.com/coreos/etcd/clientv3"
)

// Etcd represent etcd client
type Etcd interface {
	SubmitJob(string, string) (string, error)
}

type etcd struct {
	EndPoints []string
}

func (e *etcd) SubmitJob(jobName string, jobContent string) (string, error) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   e.EndPoints,
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		return "", err
	}
	defer cli.Close()
	timeout := 10 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	_, err = cli.Put(ctx, jobName, jobContent)
	cancel()
	if err != nil {
		return "", err
	}

	return jobContent, nil
}

// NewEtcd returns etcd client
func NewEtcd(endpoints []string) Etcd {
	return &etcd{
		EndPoints: endpoints,
	}
}
