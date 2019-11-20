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

// APIPaths is a map where the keys are the paths names and
// the values are the url path
var APIPaths = map[string]string{
	"login":       "/login/",
	"controllers": "/controllers/",
	"controller":  "/controllers/",
	"actuator":    "/actuators_measurements/",
}

var singletonAPI *API

var apiConfig map[string]interface{}

func init() {
	config := common.LoadConfiguration()
	apiConfig = config.API
}

// Instance creates and initializes the API object if has not been initialized yet
func Instance() *API {
	if singletonAPI != nil {
		return singletonAPI
	}

	api := &API{}
	api.configurate()
	api.buildBaseURL()
	singletonAPI = api
	return singletonAPI
}

// Get makes a GET request to the API using data as the body response
// converted to an interface{} and the body as the json sent to the API
func (api *API) Get(pathName string, data interface{}) {
	url := api.buildURL(pathName)

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

	content, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(content, data)
	if err != nil {
		log.Fatal(err)
	}
}

// GetController fetches for a specific controller
func (api *API) GetController(token string, data interface{}) {
	url := api.buildURL("controller") + token + "/"

	api.Get(url, &data)
}

// Post makes a POST request to the API using data as the body response
// converted to an interface{} and body as the request body
func (api *API) Post(pathName string, body interface{}, data interface{}) {
	url := api.buildURL(pathName)

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

// Login get the client connected to the
// api as the goberry user
func (api *API) Login() {
	// TODO: maybe make this safe?
	data := map[string]string{
		"username": "goberry",
		"password": "goberry",
	}

	var res map[string]interface{}
	api.Post("login", data, &res)
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
	value, ok := APIPaths[pathName]

	if !ok {
		return pathName
	}

	ret := api.url + value
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
