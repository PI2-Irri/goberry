package measurement

import (
	"log"
	"strconv"

	"github.com/PI2-Irri/goberry/api"
	"github.com/PI2-Irri/goberry/common"
)

// Actuator holds all data to be sent to the api
// related to the metrics of a specific actuator
type Actuator struct {
	WaterConsumption float64
	ReservoirLevel   float64
}

// Send sends the actuator data to the API
func (a *Actuator) Send() {
	log.Println("Actuator send")

	data := make(map[string]interface{}, 3)
	res := make(map[string]interface{})

	data["water_consumption"] = a.WaterConsumption
	data["reservoir_level"] = a.ReservoirLevel
	data["token"] = common.Pin

	api := api.Instance()
	api.Post("actuator", data, &res)

	val, ok := res["error"]
	if ok {
		log.Fatal("Error while sending actuator metrics:", val)
	}
}

// CreateActuator creates an actuactor with the given map[string]string
func CreateActuator(data map[string]string) *Actuator {
	actuator := &Actuator{}

	waterComsumption, err := strconv.ParseFloat(data["water_consumption"], 64)
	if err != nil {
		log.Fatal(err)
	}
	actuator.WaterConsumption = waterComsumption

	reservoirLevel, err := strconv.ParseFloat(data["reservoir_level"], 64)
	if err != nil {
		log.Fatal(err)
	}
	actuator.ReservoirLevel = reservoirLevel

	return actuator
}
