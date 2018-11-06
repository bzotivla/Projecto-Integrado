package space

type Planeta = string

var segundosUmAnoTerrestre = 31557600.0
var orbitaEmAnosTerrestres = map[Planeta]float64{
	"Terra":   1.0,
	"Mercurio": 0.2408467,
	"Venus":   0.61519726,
	"Marte":    1.8808158,
	"Jupiter": 11.862615,
	"Saturno":  29.447498,
	"Urano":  84.016846,
	"Neptuno": 164.79132,
}

func round(n float64) float64 {
	return float64(int64(n/0.01+0.5)) * 0.01
}

func Age(seconds float64, planet string) float64 {
	segundosUmAnoTerrestre := segundosUmAnoTerrestre * orbitaEmAnosTerrestres[planet]
	return round(seconds / segundosUmAnoTerrestre)
}
