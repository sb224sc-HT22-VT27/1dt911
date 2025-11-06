package mathmetics

import "math"

func Sin(angle float64) float64 {
	angle = wrapAngle(angle)

	return math.Sin(angle)
}

func WrapAngle_wrapper(angle float64) float64 {
	return wrapAngle(angle)
}

func wrapAngle(angle float64) float64 {
	wrapped_angle := math.Mod(angle, 360)
	if wrapped_angle < 0 {
		wrapped_angle += 360
	}

	return wrapped_angle
}
