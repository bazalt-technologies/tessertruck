package models

type Info struct {
	ID              int
	Name            string
	SpeedRT         float64
	EngineRPS       int
	FuelLevel       float64
	FuelSpent       float64
	FuelConsumption float64
	EngineRegime    int
	OilTemperature  float64
	EnvTemperature  float64
	RedLamp         bool
	WarnLamp        bool
}

type Rule struct {
	ID        int
	Name      string
	TruckID   int
	FieldName string
	ValInt    int
	ValFloat  float64
}

type Note struct {
	ID      int
	TruckID int
	Message string
	Time    string
}
