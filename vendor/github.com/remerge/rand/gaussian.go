package rand

import "math"

// Gaussian returns a random number of gaussian distribution Gauss(mean, stddev^2)
func Gaussian(mean, stddev float64) float64 {
	return mean + stddev*_gaussian()
}

// the probability density function, which describes the probability
// of a random variable taking on the value x
func GaussianPdf(mean, stddev, x float64) float64 {
	m := stddev * math.Sqrt(2*math.Pi)
	e := math.Exp(-math.Pow(x-mean, 2) / (2 * stddev * stddev))
	return e / m
}

func _gaussian() float64 {
	// Box-Muller Transform
	var r, x, y float64
	for r >= 1 || r == 0 {
		x = Float64Range(-1.0, 1.0)
		y = Float64Range(-1.0, 1.0)
		r = x*x + y*y
	}
	return x * math.Sqrt(-2*math.Log(r)/r)
}
