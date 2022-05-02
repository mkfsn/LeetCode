package design_parking_system

type ParkingSystem struct {
	place map[int]int
}

func Constructor(big int, medium int, small int) ParkingSystem {
	place := make(map[int]int)
	place[1] = big
	place[2] = medium
	place[3] = small

	return ParkingSystem{
		place: place,
	}
}

func (this *ParkingSystem) AddCar(carType int) bool {
	count := this.place[carType]
	if count > 0 {
		this.place[carType] = count - 1
		return true
	}

	return false
}

/**
 * Your ParkingSystem object will be instantiated and called as such:
 * obj := Constructor(big, medium, small);
 * param_1 := obj.AddCar(carType);
 */
