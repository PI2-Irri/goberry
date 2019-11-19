package controller

import (
	"log"

	"github.com/PI2-Irri/goberry/api"
	"github.com/PI2-Irri/goberry/common"
)

// Controller holds all data and methods to
// dealing with the controller in the API
type Controller struct {
	ID       int
	Name     string
	IsActive bool
	Token    string
	api      *api.API
}

// Create creates and configurates properly a controller
// with a given API object
func Create(api *api.API) *Controller {
	ctr := &Controller{
		api:   api,
		Token: common.Pin,
	}
	ctr.fetchController()
	return ctr
}

// TODO: Change to fetch from /controllers?token=xxxx
func (c *Controller) fetchController() {
	var res []map[string]interface{}
	c.api.Get("controllers", &res)

	for _, ctr := range res {
		var token string
		token = ctr["token"].(string)
		if token != c.Token {
			continue
		}
		c.initialize(ctr)
		log.Println("Fetched controller succesfully")
		return
	}
	c.registerController()
}

func (c *Controller) registerController() {
	var res map[string]interface{}
	data := map[string]interface{}{
		"name":      "Berry-" + c.Token,
		"token":     c.Token,
		"is_active": true,
	}
	c.api.Post("controllers", data, &res)
	c.initialize(res)
}

func (c *Controller) initialize(ctr map[string]interface{}) {
	c.ID = int(ctr["id"].(float64))
	c.Name = ctr["name"].(string)
	c.IsActive = ctr["is_active"].(bool)
}
