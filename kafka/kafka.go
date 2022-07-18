package kafka

import (
	"github.com/pjebs/dockertest"
)

func init() {
	dockertest.Register("kafka", nil)
}
