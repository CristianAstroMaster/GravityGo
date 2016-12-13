package gravitygo

// Taken from https://en.wikipedia.org/wiki/Gravitational_constant
const GConstant float64 = 6.67408e-11

// Taken from http://nssdc.gsfc.nasa.gov/planetary/planetfact.html
// Data is in kg, all refered to the Earth's 10^24 kg order of magnitude
const SunMass float64 = 1988500e24
const EarthMass float64 = 5.9723e24
const MoonMass float64 = 0.07346e24

// Distances and radious are measured in kilometers
const EarthMeanRadius float64 = 6371.0e3
const MoonMeanRadius float64 = 1737.1e3
