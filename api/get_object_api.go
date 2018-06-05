package api

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/gobjserver/gobjserver/core/interactor"
)

// GetObjectAPI .
type GetObjectAPI interface {
	GetAllObjectNames(context *gin.Context)
	Get(context *gin.Context)
	GetByObjectID(context *gin.Context)
}

// GetObjectAPIImpl .
type GetObjectAPIImpl struct {
	Interactor interactor.GetObjectInteractor
}

// GetAllObjectNames .
func (api GetObjectAPIImpl) GetAllObjectNames(context *gin.Context) {
	log.Println("request to get all object's names")
	channel := make(chan bool)
	go func() {
		body, err := api.Interactor.GetAll()
		if err != nil {
			context.JSON(400, err.Error())
			channel <- false
			return
		}
		context.JSON(200, body)
		channel <- true
	}()
	<-channel
}

// Get .
func (api GetObjectAPIImpl) Get(context *gin.Context) {
	objectName := context.Param("objectName")
	log.Println("request to get Object:", objectName)
	channel := make(chan bool)
	go func() {
		body := api.Interactor.Get(objectName)
		context.JSON(200, body)
		channel <- true
	}()
	<-channel
}

// GetByObjectID .
func (api GetObjectAPIImpl) GetByObjectID(context *gin.Context) {
	objectName := context.Param("objectName")
	objectID := context.Param("objectId")
	log.Println("request to get Object:", objectName, " with ID:", objectID)
	channel := make(chan bool)
	go func() {
		body, err := api.Interactor.GetByObjectID(objectName, objectID)
		if err != nil {
			context.JSON(400, err.Error())
			channel <- false
			return
		}
		context.JSON(200, body)
		channel <- true
	}()
	<-channel
}
