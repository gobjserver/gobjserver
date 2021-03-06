package gateway

import "github.com/gobjserver/gobjserver/core/entity"

// ObjectGateway .
type ObjectGateway interface {
	FindAll() ([]string, error)
	Insert(objectName string, instance interface{}) (*entity.Object, error)
	Find(objectName string) []*entity.Object
	FindByID(objectName string, objectID string) (*entity.Object, error)
	Update(objectName string, objectID string, instance *entity.Object) (*entity.Object, error)
	Delete(objectName string, objectID string) (bool, error)
}
