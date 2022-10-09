package main

import "fmt"

func main() {}

type Size int

const (
	Big Size = iota + 1
	Medium
	Small
)

type ParkingSystem struct {
	Slots map[Size]int
}

func Constructor(big int, medium int, small int) ParkingSystem {
	return ParkingSystem{
		Slots: map[Size]int{
			Big:    big,
			Medium: medium,
			Small:  small,
		},
	}
}

func (this *ParkingSystem) AddCar(carType int) bool {
	fmt.Println(this.Slots)
	if this.Slots[Size(carType)] > 0 {
		this.Slots[Size(carType)]--
		return true
	}
	return false
}
