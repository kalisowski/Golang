package structures

import "math"

type Addable interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | uintptr | float32 | float64
}

type Point[T Addable] struct {
	X T
	Y T
}

func (p Point[T]) Add(q Point[T]) Point[T] {
	return Point[T]{p.X + q.X, p.Y + q.Y}
}

func (p Point[T]) Sub(q Point[T]) Point[T] {
	return Point[T]{p.X - q.X, p.Y - q.Y}
}

func (p Point[T]) Move(q Point[T]) Point[T] {
	return Point[T]{q.X, q.Y}
}

func (p Point[T]) Distance(q Point[T]) T {
	val := (p.X-q.X)*(p.X-q.X) + (p.Y-q.Y)*(p.Y-q.Y)
	sqrt := math.Sqrt(float64(val))
	return T(sqrt)
}

type Rect[T Addable] struct {
	Min Point[T]
	Max Point[T]
}

func (r Rect[T]) Corners() [4]Point[T] {
	return [4]Point[T]{
		r.Min,
		{r.Max.X, r.Min.Y},
		r.Max,
		{r.Min.X, r.Max.Y},
	}
}

func (r Rect[T]) Area() T {
	return (r.Max.X - r.Min.X) * (r.Max.Y - r.Min.Y)
}

func (r Rect[T]) Perimeter() T {
	return 2 * (r.Max.X - r.Min.X + r.Max.Y - r.Min.Y)
}

func (r Rect[T]) Contains(p Point[T]) bool {
	return r.Min.X <= p.X && p.X <= r.Max.X && r.Min.Y <= p.Y && p.Y <= r.Max.Y
}

func (r Rect[T]) Move(p Point[T]) Rect[T] {
	return Rect[T]{
		Min: r.Min.Move(p),
		Max: r.Max.Move(p),
	}
}
