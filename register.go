package dockertest

import (
	"sync"

	"dockertest/driver"
)

var (
	driversMu sync.RWMutex
	drivers   = make(map[string]driver.Driver)
)

func Register(name string, driver driver.Driver) {
	driversMu.Lock()
	defer driversMu.Unlock()
	if driver == nil {
		panic("dockertest: Register driver is nil")
	}
	if _, dup := drivers[name]; dup {
		panic("dockertest: Register called twice for driver " + name)
	}
	drivers[name] = driver
}
