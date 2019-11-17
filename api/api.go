package api

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/PI2-Irri/goberry/common"
)

// API holds information and methods for connecting
// with the central service
type API struct {
	Protocol     string
	Host         string
	Port         int
	PollInterval int
	url          string
	token        string
	client       http.Client
}

var paths = map[string]string{
	"login":       "/login/",
	"controllers": "/controllers/",
}

var apiConfig map[string]interface{}

func init() {
	config := common.LoadConfiguration()
	apiConfig = config.API
}

// Create creates and initializes the API object
func Create() *API {
	api := &API{}
	api.configurate()
	api.buildBaseURL()
	return api
}

func (api *API) get(url string, data interface{}) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}
	api.insertHeaders(req)
	res, err := api.client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(body, data)
	if err != nil {
		log.Fatal(err)
	}
}

func (api *API) post(url string, body map[string]string, data interface{}) {
	jsonData, err := json.Marshal(body)
	if err != nil {
		log.Fatal(err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatal(err)
	}
	api.insertHeaders(req)

	res, err := api.client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	resContent, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(resContent, data)
	if err != nil {
		log.Fatal(err)
	}
}

// GetControllers gets all the controllers for the
// logged in user
func (api *API) GetControllers() {
	url := api.buildURL("controllers")
	var data []map[string]interface{}
	api.get(url, &data)
}

// Login get the client connected to the
// api as the goberry user
func (api *API) Login() {
	url := api.buildURL("login")
	// TODO: maybe make this safe?
	data := map[string]string{
		"username": "goberry",
		"password": "goberry",
	}

	var res map[string]interface{}
	api.post(url, data, &res)
	if res["token"] == nil {
		log.Fatal("User not registered")
	}
	api.token = res["token"].(string)
}

func (api *API) configurate() {
	api.Protocol = apiConfig["protocol"].(string)
	api.Port = int(apiConfig["port"].(float64))
	api.Host = apiConfig["host"].(string)
	api.PollInterval = int(apiConfig["pollInterval"].(float64))
}

func (api *API) buildURL(pathName string) string {
	ret := api.url + paths[pathName]
	return ret
}

func (api *API) buildBaseURL() {
	schema := api.Protocol + "://"
	url := schema + api.Host
	url = url + ":" + strconv.FormatInt(int64(api.Port), 10)

	api.url = url
}

func (api *API) insertHeaders(req *http.Request) {
	if api.token != "" {
		req.Header.Set("Authorization", "Token "+api.token)
	}
	req.Header.Set("content-type", "application/json")
}
