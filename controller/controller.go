package controller

import (
	"log"
	"time"

	"github.com/PI2-Irri/goberry/api"
	"github.com/PI2-Irri/goberry/command"
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
}

var instance *Controller

// Instance creates and configurates properly a controller
// with a given API object
func Instance() *Controller {
	if instance != nil {
		return instance
	}

	ctr := &Controller{
		Token: common.Pin,
	}
	ctr.fetchController()
	instance = ctr
	return instance
}

func (c *Controller) fetchController() {
	var res map[string]interface{}
	api := api.Instance()
	api.GetController(c.Token, &res)

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
	api := api.Instance()
	api.Post("controllers", data, &res)
	c.initialize(res)
}

func (c *Controller) initialize(ctr map[string]interface{}) {
	c.Name = ctr["name"].(string)
	c.IsActive = ctr["is_active"].(bool)
	c.Status = ctr["status"].(bool)
	c.Timer = int(ctr["timer"].(float64))
	c.Read = ctr["read"].(bool)
}

// Poll starts an http polling for changes in the controller state
func (c *Controller) Poll() {
	var res map[string]interface{}
	var resPatch map[string]interface{}
	api := api.Instance()
	for {
		log.Println("Polling controller")
		api.GetController(c.Token, &res)
		if !res["read"].(bool) {
			data := map[string]bool{"read": true}
			api.PatchController(common.Pin, data, &resPatch)
			// TODO: send new command
			cmd := &command.Command{
				Status: res["status"].(bool),
				Timer:  int(res["timer"].(float64)),
			}
			cmd.Send()
		}
		time.Sleep(time.Second * 5)
	}
}
