package factory_test

import (
	"testing"

	"github.com/gobjserver/gobjserver/core/entity"

	"github.com/gobjserver/gobjserver/core/factory"
	"github.com/stretchr/testify/assert"
)

// SetUp .
func SetupFactory() factory.ObjectFactory {
	return factory.ObjectFactoryImpl{}
}

// TestCreateWithDateTime .
func TestCreateWithDateTime(test *testing.T) {
	factory := SetupFactory()
	assert.NotNil(test, factory)

	result := factory.CreateWithDateTime(nil)
	assert.NotNil(test, result)

	sampleEntity := &entity.Object{}
	result = factory.CreateWithDateTime(sampleEntity)
	assert.NotNil(test, result)
}

// TestUpdateWithTime .
func TestUpdateWithTime(test *testing.T) {
	factory := SetupFactory()
	assert.NotNil(test, factory)

	sampleEntity := &entity.Object{}
	newEntity := &entity.Object{}

	result := factory.UpdateWithTime(sampleEntity, newEntity)
	assert.NotNil(test, result)
}
