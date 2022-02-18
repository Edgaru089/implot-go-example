package itype

type Direction int

const (
	XPlus  Direction = iota // X+
	XMinus                  // X-
	YPlus                   // Y+
	YMinus                  // Y-
	ZPlus                   // Z+
	ZMinus                  // Z-
)

var DirectionVeci = [6]Vec3i{
	XPlus:  {1, 0, 0},
	XMinus: {-1, 0, 0},
	YPlus:  {0, 1, 0},
	YMinus: {0, -1, 0},
	ZPlus:  {0, 0, 1},
	ZMinus: {0, 0, -1},
}
var DirectionVecf = [6]Vec3f{
	XPlus:  {1, 0, 0},
	XMinus: {-1, 0, 0},
	YPlus:  {0, 1, 0},
	YMinus: {0, -1, 0},
	ZPlus:  {0, 0, 1},
	ZMinus: {0, 0, -1},
}
var DirectionVecd = [6]Vec3d{
	XPlus:  {1, 0, 0},
	XMinus: {-1, 0, 0},
	YPlus:  {0, 1, 0},
	YMinus: {0, -1, 0},
	ZPlus:  {0, 0, 1},
	ZMinus: {0, 0, -1},
}
