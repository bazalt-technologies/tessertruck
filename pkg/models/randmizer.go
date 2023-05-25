package models

import (
	"fmt"
	"math/rand"
	"time"
)

func Randomize(id int) Info {
	rand.Seed(time.Now().UnixNano())
	return Info{
		ID:              id,
		Name:            fmt.Sprintf("TEST#%d", rand.Intn(1000000)),
		SpeedRT:         float64(rand.Intn(700) / 10),
		EngineRPS:       rand.Intn(6000),
		FuelLevel:       float64(rand.Intn(100) / 10),
		FuelSpent:       float64(rand.Intn(100) / 10),
		FuelConsumption: float64(rand.Intn(100) / 10),
		EngineRegime:    rand.Intn(5),
		OilTemperature:  float64(rand.Intn(100) / 10),
		EnvTemperature:  float64(rand.Intn(400) / 10),
		RedLamp:         false,
		WarnLamp:        false,
	}

}
