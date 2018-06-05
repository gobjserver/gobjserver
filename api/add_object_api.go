package api

import (
	"encoding/json"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/gobjserver/gobjserver/core/interactor"
)

// AddObjectAPI .
type AddObjectAPI interface {
	Add(context *gin.Context)
}

// AddObjectAPIImpl .
type AddObjectAPIImpl struct {
	Interactor interactor.AddObjectInteractor
}

// Add .
func (api AddObjectAPIImpl) Add(context *gin.Context) {
	objectName := context.Param("objectName")
	log.Println("request to add Object:", objectName)
	bytes, err := context.GetRawData()
	if err != nil {
		context.JSON(400, err.Error())
		return
	}
	channel := make(chan bool)
	go func() {
		var instance interface{}
		json.Unmarshal(bytes, &instance)
		body, err := api.Interactor.Add(objectName, &instance)
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
