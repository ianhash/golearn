package exercise

import "math"

// Vertex 坐标点
type Vertex struct {
	X, Y float64
}

// Abs abs of Vertex
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Scale scale Vertex
func (v *Vertex) Scale(f float64) {
	v.X, v.Y = v.X*f, v.Y*f
}
