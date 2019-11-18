package measurement

import (
	"log"
	"strconv"
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
