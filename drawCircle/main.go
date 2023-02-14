package main

import (
	"math"
	"math/rand"
)

func main() {}

type Solution struct {
	Radius   float64
	X_center float64
	Y_center float64
	X_start  float64
	Y_start  float64
}

func Constructor(radius float64, x_center float64, y_center float64) Solution {

	return Solution{
		Radius:   radius,
		X_center: x_center,
		Y_center: y_center,
		X_start:  x_center - radius,
		Y_start:  y_center - radius,
	}
}

func (this *Solution) RandPoint() []float64 {

	result := []float64{}
	for {
		x_point := this.X_start + rand.Float64()*2*this.Radius
		y_point := this.Y_start + rand.Float64()*2*this.Radius

		// check if the distance between center and point is equal or less than radius
		distance := math.Pow((x_point-this.X_center), 2) + math.Pow((y_point-this.Y_center), 2)

		if distance <= math.Pow(this.Radius, 2) {
			result = append(result, x_point)
			result = append(result, y_point)
			break
		}
	}

	return result
}
