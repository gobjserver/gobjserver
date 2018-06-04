package interactor

import (
	"github.com/gobjserver/gobjserver/core/entity"
	"github.com/gobjserver/gobjserver/core/factory"
	"github.com/gobjserver/gobjserver/core/gateway"
)

// AddObjectInteractor .
type AddObjectInteractor interface {
	Add(objectName string, instance *interface{}) (*entity.Object, error)
}

// AddObjectInteractorImpl is the implementation of AddObjectInteractor interface.
type AddObjectInteractorImpl struct {
	Factory factory.ObjectFactory
	Gateway gateway.ObjectGateway
}

// Add inserts new object.
func (interactor AddObjectInteractorImpl) Add(objectName string, instance *interface{}) (*entity.Object, error) {
	insertInstance := interactor.Factory.CreateWithDateTime(instance)
	return interactor.Gateway.Insert(objectName, insertInstance)
}
