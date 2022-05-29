package golang_united_school_homework

import (
	"errors"
)

// box contains list of shapes and able to perform operations on them
type box struct {
	shapes         []Shape
	shapesCapacity int // Maximum quantity of shapes that can be inside the box.
}

// NewBox creates new instance of box
func NewBox(shapesCapacity int) *box {
	return &box{
		shapesCapacity: shapesCapacity,
	}
}

// AddShape adds shape to the box
// returns the error in case it goes out of the shapesCapacity range.
func (b *box) AddShape(shape Shape) error {
	if len(b.shapes) == b.shapesCapacity {
		return errors.New("out of shapes capacity")
	}
	b.shapes = append(b.shapes, shape)
	return nil
}

// GetByIndex allows getting shape by index
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) GetByIndex(i int) (Shape, error) {
	if i >= len(b.shapes) {
		return nil, errors.New("out of shapes length")
	}
	if i > b.shapesCapacity {
		return nil, errors.New("out of shapes capacity")
	}
	return b.shapes[i], nil
}

// ExtractByIndex allows getting shape by index and removes this shape from the list.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ExtractByIndex(i int) (Shape, error) {
	if i >= len(b.shapes) {
		return nil, errors.New("out of shapes length")
	}
	if i > b.shapesCapacity {
		return nil, errors.New("out of shapes capacity")
	}

	var res = b.shapes[i]
	newShape := b.shapes[:i]

	if i == len(b.shapes)-1 {
		b.shapes = newShape
		return res, nil
	}

	newShape = append(newShape, b.shapes[i+1:]...)
	b.shapes = newShape
	return res, nil
}

// ReplaceByIndex allows replacing shape by index and returns removed shape.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ReplaceByIndex(i int, shape Shape) (Shape, error) {
	if i >= len(b.shapes) {
		return nil, errors.New("out of shapes length")
	}
	if i > b.shapesCapacity {
		return nil, errors.New("out of shapes capacity")
	}

	var res = b.shapes[i]
	newShape := b.shapes[:i]

	if i == len(b.shapes)-1 {
		b.shapes = newShape
		return res, nil
	}

	newShape = append(newShape, b.shapes[i+1:]...)
	b.shapes = newShape
	return res, nil
}

// SumPerimeter provides sum perimeter of all shapes in the list.
func (b *box) SumPerimeter() float64 {
	var sum float64
	for i := 0; i < len(b.shapes); i++ {
		sum += b.shapes[i].CalcPerimeter()
	}
	return sum
}

// SumArea provides sum area of all shapes in the list.
func (b *box) SumArea() float64 {
	var sum float64
	for i := 0; i < len(b.shapes); i++ {
		sum += b.shapes[i].CalcArea()
	}
	return sum

}

// RemoveAllCircles removes all circles in the list
// whether circles are not exist in the list, then returns an error
func (b *box) RemoveAllCircles() error {
	newShapesWithoutCircle := []Shape{}
	counterOfCircle := 0
	for i := 0; i < len(b.shapes); i++ {
		switch b.shapes[i].(type) {
		case Circle:
			counterOfCircle++
		default:
			newShapesWithoutCircle = append(newShapesWithoutCircle, b.shapes[i])
		}
	}
	if counterOfCircle == 0 {
		return errors.New("no circle in list")
	}
	return nil
}
