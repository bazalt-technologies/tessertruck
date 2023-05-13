package models

type Info struct {
	ID              int
	Name            string
	SpeedRT         float32
	EngineRPS       int
	FuelLevel       float32
	FuelSpent       float32
	FuelConsumption float32
	EngineRegime    int
	OilTemperature  float32
	EnvTemperature  float32
	RedLamp         bool
	WarnLamp        bool
}
