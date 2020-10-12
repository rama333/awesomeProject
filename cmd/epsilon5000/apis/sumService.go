package apis

import (
	"awesomeProject/cmd/epsilon5000/daos"
	"awesomeProject/cmd/epsilon5000/models"
	"awesomeProject/cmd/epsilon5000/services"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
)

func GetSumService(c *gin.Context) {

	s := services.NewSumServiceDAO(daos.NewSumServiceDAO())

	dec := json.NewDecoder(c.Request.Body)

	var ip models.Ip

	err := dec.Decode(&ip)

	if len(ip.Ip) >= 500 {
		c.JSON(http.StatusOK, models.ResponseData{Code: 200, Message: "Error: count >= 500"})
		return
	}

	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		log.Println(err)
	}

	inc := strings.Join(ip.Ip, "','")

	if list, err := s.Get(inc); err != nil || len(list) < 1 {
		c.JSON(http.StatusOK, models.ResponseData{Code: 411, Message: "Status Not Found"})
		log.Println(err)
	} else {
		fmt.Println(len(list))
		c.JSON(http.StatusOK, models.ResponseData{Code: 200, Message: "ok", Data: list})
	}

}
