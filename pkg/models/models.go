package models

type Tractor struct {
	ID         int
	Name       string
	CreateDate string
	UseDate    string
	UsePlace   string
	Teledata   Info
}

type Info struct {
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
	RulesWarnings   []string
}

type Rule struct {
	ID        int
	Name      string
	TractorID int
	FieldName string
	ValInt    int
	ValFloat  float64
}

type Note struct {
	ID        int
	TractorID int
	Message   string
	Time      string
}
