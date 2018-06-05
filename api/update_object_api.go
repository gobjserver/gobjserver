package api

import (
	"encoding/json"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/gobjserver/gobjserver/core/interactor"
)

// UpdateObjectAPI .
type UpdateObjectAPI interface {
	Update(context *gin.Context)
}

// UpdateObjectAPIImpl .
type UpdateObjectAPIImpl struct {
	Interactor interactor.UpdateObjectInteractor
}

// Update .
func (api UpdateObjectAPIImpl) Update(context *gin.Context) {
	objectName := context.Param("objectName")
	objectID := context.Param("objectId")
	log.Println("request to update Object:", objectName, " with ID:", objectID)
	bytes, err := context.GetRawData()
	if err != nil {
		context.JSON(400, err.Error())
		return
	}
	channel := make(chan bool)
	go func() {
		var instance interface{}
		json.Unmarshal(bytes, &instance)
		body, err := api.Interactor.Update(objectName, objectID, &instance)
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
