package gravitygo

import "math"

type Planet struct {
	name         string
	number       int
	mass         float64
	position     Vector
	velocity     Vector
	acceleration Vector
}

// Getter methods
func (p *Planet) GetName() string {
	return p.name
}

func (p *Planet) GetNumber() int {
	return p.number
}

func (p *Planet) GetMass() float64 {
	return p.mass
}

func (p *Planet) GetPosition() Vector {
	return p.position
}

func (p *Planet) GetVelocity() Vector {
	return p.velocity
}

func (p *Planet) GetAcceleration() Vector {
	return p.acceleration
}

// Setter methods
func (p *Planet) SetName(name_ string) {
	p.name = name_
}

func (p *Planet) SetNumber(number_ int) {
	p.number = number_
}

func (p *Planet) SetPosition(position_ Vector) {
	p.position = position_
}

func (p *Planet) SetVelocity(velocity_ Vector) {
	p.velocity = velocity_
}

func (p *Planet) SetAcceleration(acceleration_ Vector) {
	p.acceleration = acceleration_
}

// Pseudo constructor methods

func (p *Planet) MakePlanetV(number_ int, name_ string, mass_ float64, position_, velocity_ Vector) {
	p.name = name_
	p.mass = mass_
	p.number = number_
	p.position = position_
	p.velocity = velocity_
}

func (p *Planet) MakePlanet(number_ int, name_ string, mass_ float64, position_, velocity_ [3]float64) {
	p.name = name_
	p.mass = mass_
	p.number = number_
	p.position.MakeVector(position_)
	p.velocity.MakeVector(velocity_)
}

// Other operators
func (p *Planet) FieldMagnitudeAtV(vector_ Vector) float64 {
	var vectorV Vector = VectorSubstract(p.GetPosition(), vector_)
	var distance float64 = vectorV.Norm()
	if distance != 0 {
		return GConstant * p.GetMass() / distance
	} else {
		return math.NaN()
	}
}

func (p *Planet) FieldMagnitudeAt(vector_ [3]float64) float64 {
	var vectorV Vector
	vectorV.MakeVector(vector_)
	vectorV = VectorSubstract(p.GetPosition(), vectorV)
	var distance float64 = vectorV.NormSquared()
	if distance != 0 {
		return GConstant * p.GetMass() / distance
	} else {
		return math.NaN()
	}
}

func (p *Planet) ForceAtV(vector_ Vector) Vector {
	var vectorV Vector = VectorSubstract(p.GetPosition(), vector_)
	var fieldMagnitude float64 = FieldMagnitudeAtV(vector_)
	if fieldMagnitude != math.NaN() {
		vectorV.MakeVector([3]float64{0, 0, 0})
		return vectorV
	} else {
		vectorV.MakeVector([3]float64{0, 0, 0})
		return vectorV
	}
}

func (p *Planet) NetFieldOver(planets ...Planet) float64 {
	var netField float64

	for _, planet := range planets {
		if *p != planet {
			netField += planet.FieldMagnitudeAtV(p.GetPosition())
		}
	}

	return netField
}

func (p *Planet) NetForceOver(planets ...Planet) Vector {
	var netForce Vector

	for _, planet := range planets {
		if *p != planet {
			netField = VectorAdd(planet.ForceAtV(p.GetPosition()), netField)
		}
	}

	return netField
}
