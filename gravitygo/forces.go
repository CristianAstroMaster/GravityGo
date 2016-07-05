package gravitygo

func NetForceV(forces ...Vector) Vector {
	var netForce Vector

	for _, force := range forces {
		netForce = VectorAdd(netForce, force)
	}

	return netForce
}

func NetForce(forces ...[3]float64) Vector {
	var netForce Vector
	var bufferForce Vector

	for _, force := range forces {
		bufferForce.MakeVector(force)
		netForce = VectorAdd(netForce, bufferForce)
	}

	return netForce
}
