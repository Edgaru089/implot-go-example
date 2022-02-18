package itype

import "sort"

// Recti is a 2D rectangle with int coordinates.
type Recti struct {
	Left, Top     int
	Width, Height int
}

// Rectf is a 2D rectangle with float32 coordinates.
type Rectf struct {
	Left, Top     float32
	Width, Height float32
}

func (r Rectf) MinPoint() Vec2f {
	return Vec2f{r.Left, r.Top}
}

func (r Rectf) MaxPoint() Vec2f {
	return Vec2f{r.Left + r.Width, r.Top + r.Height}
}

func (r Rectf) Size() Vec2f {
	return Vec2f{r.Width, r.Height}
}

// Rectd is a 2D rectangle with float64 coordinates.
type Rectd struct {
	Left, Top     float64
	Width, Height float64
}

func (r Rectd) MinPoint() Vec2d {
	return Vec2d{r.Left, r.Top}
}

func (r Rectd) MaxPoint() Vec2d {
	return Vec2d{r.Left + r.Width, r.Top + r.Height}
}

func (r Rectd) Size() Vec2d {
	return Vec2d{r.Width, r.Height}
}

// Boxi is a 3D box with int coordinates.
type Boxi struct {
	OffX, OffY, OffZ    int
	SizeX, SizeY, SizeZ int
}

// Boxf is a 3D box with float32 coordinates.
type Boxf struct {
	OffX, OffY, OffZ    float32
	SizeX, SizeY, SizeZ float32
}

func (b Boxf) Offset(offset Vec3f) Boxf {
	return Boxf{
		OffX:  b.OffX + offset[0],
		OffY:  b.OffY + offset[1],
		OffZ:  b.OffZ + offset[2],
		SizeX: b.SizeX,
		SizeY: b.SizeY,
		SizeZ: b.SizeZ,
	}
}

func (b Boxf) MinPoint() Vec3f {
	return Vec3f{b.OffX, b.OffY, b.OffZ}
}
func (b Boxf) MaxPoint() Vec3f {
	return Vec3f{b.OffX + b.SizeX, b.OffY + b.SizeY, b.OffZ + b.SizeZ}
}

// Boxd is a 3D box with float64 coordinates.
type Boxd struct {
	OffX, OffY, OffZ    float64
	SizeX, SizeY, SizeZ float64
}

func (b Boxd) ToFloat32() Boxf {
	return Boxf{
		OffX:  float32(b.OffX),
		OffY:  float32(b.OffY),
		OffZ:  float32(b.OffZ),
		SizeX: float32(b.SizeX),
		SizeY: float32(b.SizeY),
		SizeZ: float32(b.SizeZ),
	}
}

func (b Boxd) Offset(offset Vec3d) Boxd {
	return Boxd{
		OffX:  b.OffX + offset[0],
		OffY:  b.OffY + offset[1],
		OffZ:  b.OffZ + offset[2],
		SizeX: b.SizeX,
		SizeY: b.SizeY,
		SizeZ: b.SizeZ,
	}
}

func (b Boxd) Offsetv(x, y, z float64) Boxd {
	return Boxd{
		OffX:  b.OffX + x,
		OffY:  b.OffY + y,
		OffZ:  b.OffZ + z,
		SizeX: b.SizeX,
		SizeY: b.SizeY,
		SizeZ: b.SizeZ,
	}
}

func (b Boxd) MinPoint() Vec3d {
	return Vec3d{b.OffX, b.OffY, b.OffZ}
}
func (b Boxd) MaxPoint() Vec3d {
	return Vec3d{b.OffX + b.SizeX, b.OffY + b.SizeY, b.OffZ + b.SizeZ}
}

func (b Boxd) Contains(point Vec3d) bool {
	return point[0] >= b.OffX && point[0] <= b.OffX+b.SizeX &&
		point[1] >= b.OffY && point[1] <= b.OffY+b.SizeY &&
		point[2] >= b.OffZ && point[2] <= b.OffZ+b.SizeZ
}

func pointIntersect(n, m, p, q float64) (min, len float64) {
	if m < p || q < n { // no intersection
		return 0, 0
	}

	arr := []float64{n, m, p, q}
	sort.Float64s(arr)

	return arr[1], arr[2] - arr[1]
}

func (box1 Boxd) Intersect(box2 Boxd) (ok bool, intersect Boxd) {
	a, b := pointIntersect(box1.OffX, box1.OffX+box1.SizeX, box2.OffX, box2.OffX+box2.SizeX)
	c, d := pointIntersect(box1.OffY, box1.OffY+box1.SizeY, box2.OffY, box2.OffY+box2.SizeY)
	e, f := pointIntersect(box1.OffZ, box1.OffZ+box1.SizeZ, box2.OffZ, box2.OffZ+box2.SizeZ)

	if b == 0 || d == 0 || f == 0 {
		return false, Boxd{}
	} else {
		return true, Boxd{
			OffX:  a,
			SizeX: b,
			OffY:  c,
			SizeY: d,
			OffZ:  e,
			SizeZ: f,
		}
	}
}
