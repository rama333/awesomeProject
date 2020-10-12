package daos

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/mitchellh/mapstructure"
	"io/ioutil"
	"net/http"
	"strings"
)

type Context struct {
	sessionKey string
	host       string
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
type SelectQuery interface{}
type GetParameters struct {
	CountOutput            bool                   `json:"countOutput,omitempty"`
	Editable               bool                   `json:"editable,omitempty"`
	ExcludeSearch          bool                   `json:"excludeSearch,omitempty"`
	Filter                 map[string]interface{} `json:"filter,omitempty"`
	Limit                  int                    `json:"limit,omitempty"`
	Output                 SelectQuery            `json:"output,omitempty"`
	PreserveKeys           bool                   `json:"preservekeys,omitempty"`
	Search                 map[string]string      `json:"search,omitempty"`
	SearchByAny            bool                   `json:"searchByAny,omitempty"`
	SearchWildcardsEnabled bool                   `json:"searchWildcardsEnabled,omitempty"`
	SortField              []string               `json:"sortfield,omitempty"`
	SortOrder              []string               `json:"sortorder,omitempty"` // has defined consts, see above
	StartSearch            bool                   `json:"startSearch,omitempty"`
}

func (c *Context) Login(host, login, password string) error {

	c.host = host
	var err error

	r := UserLoginParams{
		User:     login,
		Password: password,
	}

	if c.sessionKey, _, err = c.userLogin(r); err != nil {
		return err
	}

	return nil

}

func (z *Context) request(method string, params interface{}, result interface{}) (int, error) {

	resp := responseData{
		Result: result,
	}

	req := requestData{
		JSONRPC: "2.0",
		Method:  method,
		Params:  params,
		Auth:    z.sessionKey,
		ID:      1,
	}

	status, err := z.httpPost(req, &resp)
	if err != nil {
		return status, err
	}

	if resp.Error.Code != 0 {
		return status, errors.New(resp.Error.Data + " " + resp.Error.Message)
	}

	return status, nil
}

func (z *Context) httpPost(in interface{}, out interface{}) (int, error) {

	s, err := json.Marshal(in)
	if err != nil {
		return 0, err
	}
	req, err := http.NewRequest("POST", z.host, strings.NewReader(string(s)))
	if err != nil {
		return 0, err
	}

	req.Header.Add("Content-Type", "application/json-rpc")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return 0, err
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		if bodyBytes, err := ioutil.ReadAll(res.Body); err == nil {
			return res.StatusCode, errors.New(string(bodyBytes))
		}
	} else {
		if out != nil {

			rawConf := make(map[string]interface{})

			dJ := json.NewDecoder(res.Body)

			if err := dJ.Decode(&rawConf); err != nil {
				return res.StatusCode, fmt.Errorf("json decode error: %v", err)
			}

			dM, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
				WeaklyTypedInput: true,
				Result:           out,
				TagName:          "json",
			})
			if err != nil {
				return res.StatusCode, fmt.Errorf("mapstructure create decoder error: %v", err)
			}

			if err := dM.Decode(rawConf); err != nil {
				return res.StatusCode, fmt.Errorf("mapstructure decode error: %v", err)
			}
		}
	}

	return res.StatusCode, nil
}
