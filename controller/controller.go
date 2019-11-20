package controller

import (
	"log"

	"github.com/PI2-Irri/goberry/api"
	"github.com/PI2-Irri/goberry/common"
)

// Controller holds all data and methods to
// dealing with the controller in the API
type Controller struct {
	Name     string
	IsActive bool
	Token    string
	Timer    int
	Read     bool
	Status   bool
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

func (c *Controller) fetchController() {
	var res map[string]interface{}
	c.api.GetController(c.Token, &res)

	// if controller doesnt exist
	value, ok := res["detail"]
	if ok && value == "Not found." {
		c.registerController()
		return
	}

	// controller already exists
	log.Println("Controller already exists")
	c.initialize(res)
}

func (c *Controller) registerController() {
	log.Println("Registering controller")

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
	c.Name = ctr["name"].(string)
	c.IsActive = ctr["is_active"].(bool)
	c.Status = ctr["permit_irrigation"].(bool)
	c.Timer = int(ctr["time_to_irrigate"].(float64))
	c.Read = ctr["requested"].(bool)
}
