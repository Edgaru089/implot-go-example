package itype

import (
	"math"
)

// Vec2i is a two-element int vector
type Vec2i [2]int

func (v Vec2i) ToFloat32() Vec2f { return Vec2f{float32(v[0]), float32(v[1])} }

func (v Vec2i) Add(add Vec2i) Vec2i        { return Vec2i{v[0] + add[0], v[1] + add[1]} }
func (v Vec2i) MultiplyInt(mult int) Vec2i { return Vec2i{v[0] * mult, v[1] * mult} }

// Vec3i is a three-element int vector
type Vec3i [3]int

func (v Vec3i) ToFloat32() Vec3f { return Vec3f{float32(v[0]), float32(v[1]), float32(v[2])} }
func (v Vec3i) ToFloat64() Vec3d { return Vec3d{float64(v[0]), float64(v[1]), float64(v[2])} }

func (v Vec3i) Add(add Vec3i) Vec3i         { return Vec3i{v[0] + add[0], v[1] + add[1], v[2] + add[2]} }
func (v Vec3i) Addv(x, y, z int) Vec3i      { return Vec3i{v[0] + x, v[1] + y, v[2] + z} }
func (v Vec3i) MultiplyInt(mult int) Vec3i  { return Vec3i{v[0] * mult, v[1] * mult, v[2] * mult} }
func (v Vec3i) Multiplyv(x, y, z int) Vec3i { return Vec3i{v[0] * x, v[1] * y, v[2] * z} }
func (v Vec3i) Multiply(mult Vec3i) Vec3i {
	return Vec3i{v[0] * mult[0], v[1] * mult[1], v[2] * mult[2]}
}

// Vec4i is a four-element int vector
type Vec4i [4]int

func Vec4iToFloat32(v Vec4i) Vec4f {
	return Vec4f{float32(v[0]), float32(v[1]), float32(v[2]), float32(v[3])}
}

func (v Vec4i) Add(add Vec4i) Vec4i {
	return Vec4i{
		v[0] + add[0],
		v[1] + add[1],
		v[2] + add[2],
		v[3] + add[3],
	}
}
func (v Vec4i) MultiplyInt(mult int) Vec4i {
	return Vec4i{
		v[0] * mult,
		v[1] * mult,
		v[2] * mult,
		v[3] * mult,
	}
}

// Vec2f is a two-element float vector
type Vec2f [2]float32

func (v Vec2f) Add(add Vec2f) Vec2f {
	return Vec2f{v[0] + add[0], v[1] + add[1]}
}

func (v Vec2f) Addv(x, y float32) Vec2f {
	return Vec2f{v[0] + x, v[1] + y}
}

func (v Vec2f) Multiply(x float32) Vec2f {
	return Vec2f{v[0] * x, v[1] * x}
}

func (v Vec2f) Floor() Vec2i {
	return Vec2i{
		int(math.Floor(float64(v[0]))),
		int(math.Floor(float64(v[1]))),
	}
}
func (v Vec2f) Length() float32 {
	return float32(math.Sqrt(float64(v[0]*v[0] + v[1]*v[1])))
}

func (v Vec2f) Normalize() Vec2f {
	l := v.Length()
	return Vec2f{v[0] / l, v[1] / l}
}

func (v Vec2f) ToFloat64() Vec2d {
	return Vec2d{float64(v[0]), float64(v[1])}
}

// Vec3f is a three-element float vector
type Vec3f [3]float32

func (v Vec3f) Negative() Vec3f {
	return Vec3f{-v[0], -v[1], -v[2]}
}

func (v Vec3f) Add(add Vec3f) Vec3f {
	return Vec3f{v[0] + add[0], v[1] + add[1], v[2] + add[2]}
}

func (v Vec3f) Addv(x, y, z float32) Vec3f {
	return Vec3f{v[0] + x, v[1] + y, v[2] + z}
}

func (v Vec3f) Multiply(mult float32) Vec3f {
	return Vec3f{v[0] * mult, v[1] * mult, v[2] * mult}
}

func (v Vec3f) Floor() Vec3i {
	return Vec3i{
		int(math.Floor(float64(v[0]))),
		int(math.Floor(float64(v[1]))),
		int(math.Floor(float64(v[2]))),
	}
}
func (v Vec3f) Length() float32 {
	v0, v1, v2 := float64(v[0]), float64(v[1]), float64(v[2])
	return float32(math.Sqrt(v0*v0 + v1*v1 + v2*v2))
}

func (v Vec3f) Normalize() Vec3f {
	l := v.Length()
	return Vec3f{v[0] / l, v[1] / l, v[2] / l}
}

func (v Vec3f) ToFloat64() Vec3d {
	return Vec3d{float64(v[0]), float64(v[1]), float64(v[2])}
}

// Vec4f is a four-element float vector
type Vec4f [4]float32

// Vec2d is a two-element float64 vector
type Vec2d [2]float64

func (v Vec2d) Add(add Vec2d) Vec2d {
	return Vec2d{v[0] + add[0], v[1] + add[1]}
}

func (v Vec2d) Addv(x, y float64) Vec2d {
	return Vec2d{v[0] + x, v[1] + y}
}

func (v Vec2d) Multiply(x float64) Vec2d {
	return Vec2d{v[0] * x, v[1] * x}
}

func (v Vec2d) Floor() Vec2i {
	return Vec2i{
		int(math.Floor(float64(v[0]))),
		int(math.Floor(float64(v[1]))),
	}
}
func (v Vec2d) Length() float64 {
	return math.Sqrt(v[0]*v[0] + v[1]*v[1])
}

func (v Vec2d) Normalize() Vec2d {
	l := v.Length()
	return Vec2d{v[0] / l, v[1] / l}
}

func (v Vec2d) ToFloat32() Vec2f {
	return Vec2f{float32(v[0]), float32(v[1])}
}

// Vec3d is a three-element float64 vector
type Vec3d [3]float64

func (v Vec3d) Negative() Vec3d {
	return Vec3d{-v[0], -v[1], -v[2]}
}

func (v Vec3d) Add(add Vec3d) Vec3d {
	return Vec3d{v[0] + add[0], v[1] + add[1], v[2] + add[2]}
}

func (v Vec3d) Addf(add Vec3f) Vec3d {
	return Vec3d{v[0] + float64(add[0]), v[1] + float64(add[1]), v[2] + float64(add[2])}
}

func (v Vec3d) Addv(x, y, z float64) Vec3d {
	return Vec3d{v[0] + x, v[1] + y, v[2] + z}
}

func (v Vec3d) Multiply(mult float64) Vec3d {
	return Vec3d{v[0] * mult, v[1] * mult, v[2] * mult}
}

func (v1 Vec3d) Cross(v2 Vec3d) Vec3d {
	return Vec3d{v1[1]*v2[2] - v1[2]*v2[1], v1[2]*v2[0] - v1[0]*v2[2], v1[0]*v2[1] - v1[1]*v2[0]}
}

func (v1 Vec3d) Dot(v2 Vec3d) float64 {
	return v1[0]*v2[0] + v1[1]*v2[1] + v1[2]*v2[2]
}

func (v Vec3d) Length() float64 {
	return math.Sqrt(v[0]*v[0] + v[1]*v[1] + v[2]*v[2])
}

func (v Vec3d) Normalize() Vec3d {
	l := v.Length()
	return Vec3d{v[0] / l, v[1] / l, v[2] / l}
}

func (v Vec3d) Floor() Vec3i {
	return Vec3i{
		int(math.Floor(v[0])),
		int(math.Floor(v[1])),
		int(math.Floor(v[2])),
	}
}

func (v Vec3d) Ceiling() Vec3i {
	return Vec3i{
		int(math.Ceil(v[0])),
		int(math.Ceil(v[1])),
		int(math.Ceil(v[2])),
	}
}

func (v Vec3d) ToFloat32() Vec3f {
	return Vec3f{float32(v[0]), float32(v[1]), float32(v[2])}
}

// Vec4d is a four-element float64 vector
type Vec4d [4]float64
