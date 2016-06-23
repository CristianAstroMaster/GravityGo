package main

import (
	"fmt"
	"github.com/Daniel-M/GravityGo/GravityGo"
)

func main() {
	fmt.Println("Earth Mass is ", gravitygo.EarthMass)

	var tierra, luna gravitygo.Planet
	var pos, vel gravitygo.Vector

	pos.MakeVector([3]float64{1, 1, 1})
	vel.MakeVector([3]float64{1, 1, 1})

	fmt.Println("Substract", gravitygo.VectorSubstract(pos, vel))

	tierra.MakePlanet(1, "Tierra", gravitygo.EarthMass, [3]float64{0, 0, 0}, [3]float64{0, 0, 0})
	luna.MakePlanet(2, "Luna", gravitygo.MoonMass, [3]float64{0.3844e6, 0, 0}, [3]float64{0, 0, 0})

	fmt.Println("Planet", tierra.GetName(), "created with mass", tierra.GetMass(), "at position", tierra.GetPosition())

	fmt.Println("Field at origin", tierra.FieldMagnitudeAt([3]float64{6400e3, 0, 0}))

	fmt.Println(tierra.NetFieldOver(tierra, luna))
	fmt.Println(tierra.NetFieldOver(luna))

	fmt.Println(luna.NetFieldOver(tierra, luna))
	fmt.Println(luna.NetFieldOver(tierra))
}
