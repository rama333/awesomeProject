package daos

import (
	"awesomeProject/cmd/epsilon5000/config"
	"awesomeProject/cmd/epsilon5000/models"
	"encoding/json"
	"fmt"
	"github.com/mitchellh/mapstructure"
	"io/ioutil"
	"net/http"
	"strings"
)

type ZabbixCameraDAO struct {

}

func NewZabbixDAO() *ZabbixCameraDAO {

	return &ZabbixCameraDAO{}
}

type requestData struct {
	JSONRPC string      `json:"jsonrpc"`
	Method  string      `json:"method"`
	Params  interface{} `json:"params,omitempty"`
	Auth    string      `json:"auth,omitempty"`
	ID      int         `json:"id"`
}

type responseData struct {
	JSONRPC string      `json:"jsonrpc"`
	Result  interface{} `json:"result"`
	Error   struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
		Data    string `json:"data"`
	} `json:"error"`
	ID int `json:"id"`
}


func Request(method string, params interface{}, result interface{}){

	resp := responseData{
		Result: result,
	}

	req := requestData{
		JSONRPC: "2.0",
		Method:  method,
		Params:  params,
		Auth:    "null",
		ID:      1,
	}

	s, _ := json.Marshal(req)


	re, _ := http.NewRequest("POST", config.Config.ZbxHost, strings.NewReader(string(s)))


	// Set headers
	re.Header.Add("Content-Type", "application/json-rpc")

	// Make request
	res, _ := http.DefaultClient.Do(re)

	defer res.Body.Close()

	if res.StatusCode != 200 {
		if bodyBytes, err := ioutil.ReadAll(res.Body); err == nil {
			fmt.Println(bodyBytes)
		}
	} else {
		if result != nil {

			rawConf := make(map[string]interface{})

			dJ := json.NewDecoder(res.Body)
			if err := dJ.Decode(&rawConf); err != nil {
				fmt.Errorf("json decode error: %v", err)
			}

			dM, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
				WeaklyTypedInput: true,
				Result:           resp,
				TagName:          "json",
			})
			if err != nil {
				fmt.Errorf("mapstructure create decoder error: %v", err)
			}

			if err := dM.Decode(rawConf); err != nil {
				fmt.Errorf("mapstructure decode error: %v", err)
			}

			fmt.Println(rawConf)
		}
	}

	fmt.Println(resp)

}


func (s ZabbixCameraDAO) Get(id models.Id) ([]models.СameraIncidents, error)  {

	//var z zabbix.Context
	list := []models.СameraIncidents{}








	Request()



	list = append(list, models.СameraIncidents{"h","","h.Host","",""})

	return list, nil
}