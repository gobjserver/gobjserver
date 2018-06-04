package config

import (
	"github.com/gobjserver/gobjserver/core/factory"
)

// CreateObjectFactory .
func CreateObjectFactory() factory.ObjectFactory {
	return factory.ObjectFactoryImpl{}
}
