package datasource

import (
	"github.com/go-xorm/xorm"
	"sync"
)

var (
	masterEngine *xorm.Engine
	slaveEngine *xorm.Engine
	lock sync.Locker
)

func InstanceMaster() *xorm.Engine {
	// todo
	return masterEngine
}

func InstanceSlave() *xorm.Engine {
	// todo
	return slaveEngine
}
