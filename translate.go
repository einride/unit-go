package unit

import "math"

// translate a number x from one interval to another.
//
//         (b-a)(x - min)
//  f(x) = -------------- + a
//            max - min
func translate(x, min, max, a, b float64) float64 {
	x = math.Min(max, math.Max(min, x))
	return (((b - a) * (x - min)) / (max - min)) + a
}
