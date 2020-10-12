package daos

import (
	"awesomeProject/cmd/epsilon5000/config"
	"awesomeProject/cmd/epsilon5000/models"
	"fmt"
	"github.com/mitchellh/mapstructure"
	"time"
)

type InfoCamera struct {
	Id          int    `json:"id"`
	Host_       string `json:"host"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type ZabbixCameraDAO struct {
}

func NewZabbixDAO() *ZabbixCameraDAO {

	return &ZabbixCameraDAO{}
}

func (s ZabbixCameraDAO) Get(id models.Id) ([]models.СameraIncidents, error) {

	//var z zabbix.Context
	list := []models.СameraIncidents{}

	var z Context

	if err := z.Login(config.Config.ZbxHost, config.Config.ZbxLogin, config.Config.ZbxPassword); err != nil {
		fmt.Println("Login error:", err)
	}

	fmt.Println(z.sessionKey)

	unixTime := time.Now().Unix() - 86400

	//time := time.Unix(unixTime, 0)

	//fmt.Println(tt.Format("2006-01-02 15:04:05"))

	res, err := z.GetEvent(EventObject{
		Groupids:              []int{38, 160, 149},
		Time_from:             unixTime,
		Output:                "extend",
		Sortfield:             []string{"clock", "eventid"},
		Sortorder:             "DESC",
		SelectRelatedObject:   []string{"name", "description"},
		SelectHosts:           []string{"name", "description", "host", "interfaceids"},
		SelectAcknowledges:    "extend",
		SelectTags:            "extend",
		SelectSuppressionData: "extend",
		Recent:                "true",
	})

	if err != nil {
		fmt.Println(err)
	}

	//fmt.Println(res

	//test := res.([]EventResult)

	//fmt.Println(res["host"])

	var result []EventResult

	err = mapstructure.Decode(res, &result)
	if err != nil {
		panic(err)
	}

	//var p =  make(map[int][]InfoCamera)
	//
	//for i, obj := range result{
	//
	//	p[5] = append(p[5], InfoCamera{Id: 5,Host_: "", Name: "", Description: ""})
	//}
	//
	//var host Host
	//err1 := mapstructure.Decode(result[0].Hosts, &host)
	//if err1 != nil {
	//	fmt.Println(err1)
	//}
	//
	//
	fmt.Println(result[0].Hosts[0].Name)
	fmt.Println(result[0].Hosts[0].Host)

	list = append(list, models.СameraIncidents{"h", "", "h.Host", "", ""})

	return list, nil
}
