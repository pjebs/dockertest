package dockertest

import (
	"fmt"

	"github.com/ory/dockertest/v3"
)

type Pool struct {
	Pool *dockertest.Pool

	Resources []*dockertest.Resource
}

// DockerStartUp creates a conncetion to the Docker API.
//
// See: https://pkg.go.dev/github.com/ory/dockertest#NewPool
func DockerStartUp(endpoint string, certpath ...string) (*Pool, error) {
	pool := &Pool{}
	var err error
	if len(certpath) > 0 {
		pool.Pool, err = dockertest.NewTLSPool(endpoint, certpath[0])
		if err != nil {
			return nil, fmt.Errorf("could not connect to docker: %w", err)
		}
	} else {
		pool.Pool, err = dockertest.NewPool(endpoint)
		if err != nil {
			return nil, fmt.Errorf("could not connect to docker: %w", err)
		}
	}

	err = pool.Pool.Client.Ping()
	if err != nil {
		return nil, err
	}

	return pool, nil
}

func (p *Pool) CleanUp() error {
	return nil
}
