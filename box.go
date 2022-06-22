package golang_united_school_homework

import "errors"

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
	if len(b.shapes) >= b.shapesCapacity {
		return errors.New("box is full")
	} else {
		b.shapes = append(b.shapes, shape)
		return nil
	}
}

// GetByIndex allows getting shape by index
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) GetByIndex(i int) (Shape, error) {
	if len(b.shapes) > i && i >= 0 {
		return b.shapes[i], nil
	} else {
		return nil, errors.New("dsf")
	}

}

// ExtractByIndex allows getting shape by index and removes this shape from the list.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ExtractByIndex(i int) (Shape, error) {
	if len(b.shapes) > i && i >= 0 {
		tempShapes := b.shapes[i]
		if i == 0 {
			b.shapes = b.shapes[1:]
			return tempShapes, nil
		} else if i == len(b.shapes)-1 {
			b.shapes = b.shapes[:len(b.shapes)-1]
			return tempShapes, nil
		} else {
			sliceOne := b.shapes[0:i]
			sliceTwo := b.shapes[i+1:]
			b.shapes = sliceOne
			b.shapes = append(b.shapes, sliceTwo...)
			return tempShapes, nil
		}
	} else {
		return nil, errors.New("index went out of the range, func: ExtractByIndex")
	}

}

// ReplaceByIndex allows replacing shape by index and returns removed shape.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ReplaceByIndex(i int, shape Shape) (Shape, error) {
	if len(b.shapes) > i && i >= 0 {
		tempShapes := b.shapes[i]
		b.shapes[i] = shape
		return tempShapes, nil
	} else {
		return nil, errors.New("index went out of the range, func: ReplaceByIndex")
	}

}

// SumPerimeter provides sum perimeter of all shapes in the list.
func (b *box) SumPerimeter() float64 {
	var summ float64 = 0
	for _, shape := range b.shapes {
		summ += shape.CalcPerimeter()
	}
	return summ

}

// SumArea provides sum area of all shapes in the list.
func (b *box) SumArea() float64 {
	var summ float64 = 0
	for _, shape := range b.shapes {
		summ += shape.CalcArea()
	}
	return summ

}

// RemoveAllCircles removes all circles in the list
// whether circles are not exist in the list, then returns an error
func (b *box) RemoveAllCircles() error {
	var numberOfRemove int = 0
	for i, shape := range b.shapes {
		_, ok := shape.(Circle)
		if ok {
			_, _ = b.ExtractByIndex(i)
			numberOfRemove++
		}
	}
	if numberOfRemove > 0 {
		return nil
	} else {
		return errors.New("ttt Circles are not exist in the list, func: RemoveAllCircles")
	}

}
