package api

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/gobjserver/gobjserver/core/interactor"
)

// DeleteObjectAPI .
type DeleteObjectAPI interface {
	Delete(context *gin.Context)
}

// DeleteObjectAPIImpl .
type DeleteObjectAPIImpl struct {
	Interactor interactor.DeleteObjectInteractor
}

// Delete .
func (api DeleteObjectAPIImpl) Delete(context *gin.Context) {
	objectName := context.Param("objectName")
	objectID := context.Param("objectId")
	log.Println("request to delete Object:", objectName, " with ID:", objectID)
	channel := make(chan bool)
	go func() {
		body, err := api.Interactor.Delete(objectName, objectID)
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
