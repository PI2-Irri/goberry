package measurement

import (
	"log"
	"strconv"
	"strings"

	"github.com/PI2-Irri/goberry/api"
	"github.com/PI2-Irri/goberry/common"
)

// Module holds all data to be sent to the api
// related to the metrics of a specific module
type Module struct {
	Temperature    float64
	GroundHumidity float64
	BatteryLevel   int64
	RFAddress      int64
}

// Send sends the module data to the API
func (m *Module) Send() {
	log.Println("MOD:", m.Temperature, m.GroundHumidity, m.BatteryLevel, m.RFAddress)

	data := make(map[string]interface{}, 4)
	res := make(map[string]interface{})

	data["soil_temperature"] = m.Temperature
	data["ground_humidity"] = m.GroundHumidity
	data["battery_level"] = m.BatteryLevel
	data["rf_address"] = m.RFAddress
	data["token"] = common.Pin

	api := api.Instance()
	api.Post("module", data, &res)

	val, ok := res["error"].(string)

	if ok && strings.Contains(val, "Zone not found") {
		log.Println("No zone found, ignoring message")
		return
	}

	if ok {
		log.Fatal("Error ocurred during module measurement send:", val)
	}
}

// CreateModule creates a module with the given map[string]string
func CreateModule(data map[string]string) *Module {
	module := &Module{}

	temperature, err := strconv.ParseFloat(data["temperature"], 64)
	if err != nil {
		log.Fatal(err)
	}
	module.Temperature = temperature

	groundHumidity, err := strconv.ParseFloat(data["ground_humidity"], 64)
	if err != nil {
		log.Fatal(err)
	}
	module.GroundHumidity = groundHumidity

	batteryLevel, err := strconv.ParseInt(data["battery_level"], 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	module.BatteryLevel = batteryLevel

	rfAdress, err := strconv.ParseInt(data["rf_address"], 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	module.RFAddress = rfAdress

	return module
}
