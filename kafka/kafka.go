package kafka

import (
	dockertest ".."
)

func init() {
	dockertest.Register("kafka", nil)
}
