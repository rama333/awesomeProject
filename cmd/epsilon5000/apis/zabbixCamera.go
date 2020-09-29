package apis

import (
	"awesomeProject/cmd/epsilon5000/daos"
	"awesomeProject/cmd/epsilon5000/models"
	"awesomeProject/cmd/epsilon5000/services"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func GetCameraOfZabbix(c *gin.Context){

	s:= services.NewZabbixCameraDAO(daos.NewZabbixDAO())

	dec := json.NewDecoder(c.Request.Body)

	var id models.Id

	err := dec.Decode(&id)

	if err != nil{
		c.JSON(http.StatusOK, models.ResponseError{Status: "error", Description: "Status Bad Request"})
		return
	}

	if list, err := s.Get(id); err != nil{
		log.Println(err)
		c.JSON(http.StatusOK, models.ResponseError{Status: "error", Description: "Уrror while executing the request"})
	} else {
		c.JSON(http.StatusOK, list)
	}

}
