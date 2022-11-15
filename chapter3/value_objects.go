package chapter3

type Point struct {
	x int
	y int
}

func NewPoint(x, y int) Point {
	return Point{
		x: x,
		y: y,
	}
}

const (
	DirectionUnknown = iota
	DirectionNorth
	DirectionSouth
	DirectionEast
	DirectionWest
)

func TrackPlayer() {
	currLocation := NewPoint(3, 4)
	currLocation = Move(currLocation, DirectionNorth)
}

func Move(currLocation Point, direction int) Point {
	switch direction {
	case DirectionNorth:
		return NewPoint(currLocation.x, currLocation.y+1)
	case DirectionSouth:
		return NewPoint(currLocation.x, currLocation.y-1)
	case DirectionEast:
		return NewPoint(currLocation.x+1, currLocation.y)
	case DirectionWest:
		return NewPoint(currLocation.x-1, currLocation.x)
	default:
		//do a barrel roll
	}
	return currLocation
}
