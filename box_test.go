package golang_united_school_homework

import (
	"errors"
	"testing"
)

type testDataAddShapes struct {
	shapes             []Shape
	wantmaxcapacityBox int
	wantCapacity       int
	wantError          error
}

func Test_box_AddShape(t *testing.T) {

	tests := []testDataAddShapes{
		testDataAddShapes{shapes: []Shape{Triangle{Side: 5}, Triangle{Side: 5}, Triangle{Side: 5}, Triangle{Side: 5}}, wantmaxcapacityBox: 6, wantCapacity: 5, wantError: nil},
		testDataAddShapes{shapes: []Shape{Circle{Radius: 2}, Triangle{Side: 4}}, wantmaxcapacityBox: 3, wantCapacity: 2, wantError: nil},
		testDataAddShapes{shapes: []Shape{Circle{Radius: 2}, Triangle{Side: 4}, Rectangle{Height: 4, Weight: 2}}, wantmaxcapacityBox: 2, wantCapacity: 2, wantError: errors.New("box is full")},
	}
	for iTest, test := range tests {
		box := NewBox(test.wantmaxcapacityBox)
		for _, shape := range test.shapes {
			gotError := box.AddShape(shape)
			if gotError != test.wantError {
				t.Errorf("Test number %d - failed", iTest)
			}
		}
	}

	box := NewBox(5)
	var triangleOne Triangle = Triangle{Side: 5}
	box.AddShape(triangleOne)
	if len(box.shapes) != 1 {
		t.Error("box is empty")
	}
}
