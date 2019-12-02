package fuelcounterupper

func CalculateFuelByMass(mass int) int {
	ret := mass/3 - 2
	if ret > 0 {
		return ret
	}

	return 0
}
