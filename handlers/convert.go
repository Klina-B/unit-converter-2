package handlers


var LengthUnitMap = map[string]float64{
	"m": 1.0,
	"km": 0.001,
	"cm": 10,
	"mm": 1000,
}

var WeightUnitMap = map[string]float64{
	"kg": 1.0,
	"t": 0.001,
	"g": 1000,
}

type Unit interface{
	Convert() float64
}
type UnitLength struct{
	Value float64
	Unit string
}
type UnitWeight struct{
	Value float64 
	Unit string
}

func (u UnitLength) Convert(to_unit string) float64{
	//fmt.Println(u.Value, LengthUnitMap[""])
	return u.Value / LengthUnitMap[u.Unit] * LengthUnitMap[to_unit]
}

func (u UnitWeight) Convert(to_unit string) float64{
	return u.Value / WeightUnitMap[u.Unit]* WeightUnitMap[to_unit]
}