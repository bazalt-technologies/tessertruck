package models

import (
	"fmt"
	"math/rand"
	"time"
)

func Randomize(id int) Tractor {
	rand.Seed(time.Now().UnixNano())
	return Tractor{
		ID:   id,
		Name: fmt.Sprintf("TEST#%d", id),
		Teledata: Info{
			SpeedRT:         float64(rand.Intn(700) / 10),
			EngineRPS:       rand.Intn(6000),
			FuelLevel:       float64(rand.Intn(500) / 10),
			FuelSpent:       float64(rand.Intn(500) / 10),
			FuelConsumption: float64(rand.Intn(1000) / 10),
			EngineRegime:    rand.Intn(5),
			OilTemperature:  float64(rand.Intn(100) / 10),
			EnvTemperature:  float64(rand.Intn(400) / 10),
			RedLamp:         false,
			WarnLamp:        false,
		},
	}

}
