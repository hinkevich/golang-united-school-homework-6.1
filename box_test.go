package golang_united_school_homework

import (
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
		{shapes: []Shape{Triangle{Side: 5}, Triangle{Side: 5}, Triangle{Side: 5}, Triangle{Side: 5}}, wantmaxcapacityBox: 6, wantCapacity: 5, wantError: nil},
		{shapes: []Shape{Circle{Radius: 2}, Triangle{Side: 4}}, wantmaxcapacityBox: 3, wantCapacity: 2, wantError: nil},
		{shapes: []Shape{Circle{Radius: 2}, Triangle{Side: 4}, Rectangle{Height: 4, Weight: 2}}, wantmaxcapacityBox: 3, wantCapacity: 3, wantError: nil},
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

type testDataRemoveAllCircles struct {
	shapes     []Shape
	wantShapes int
}

func TestRemoveAllCircles(t *testing.T) {
	tests := []testDataRemoveAllCircles{
		{shapes: []Shape{Triangle{Side: 5}, Triangle{Side: 5}}, wantShapes: 2},
		{shapes: []Shape{Triangle{Side: 5}, Triangle{Side: 5}, Circle{Radius: 2}}, wantShapes: 2},
		{shapes: []Shape{Circle{Radius: 2}, Triangle{Side: 5}, Triangle{Side: 5}, Circle{Radius: 2}}, wantShapes: 2},
		{shapes: []Shape{Circle{Radius: 2}, Circle{Radius: 2}, Circle{Radius: 2}, Triangle{Side: 5}, Triangle{Side: 5}, Circle{Radius: 2}}, wantShapes: 2},
		{shapes: []Shape{Rectangle{Height: 4, Weight: 2}, Circle{Radius: 2}, Rectangle{Height: 4, Weight: 2}, Circle{Radius: 2}, Circle{Radius: 2}, Triangle{Side: 5}, Triangle{Side: 5}, Circle{Radius: 2}}, wantShapes: 4},
	}
	for i, test := range tests {
		box := NewBox(len(test.shapes))
		for _, shape := range test.shapes {
			box.AddShape(shape)
		}
		box.RemoveAllCircles()
		if len(box.shapes) != test.wantShapes {
			t.Errorf("test %d is fail, len want  %d, len got: %d", i, test.wantShapes, len(test.shapes))
		}

	}

}
func TestRemoveAllCirclesWithoutCircle(t *testing.T) {
	tests := []testDataRemoveAllCircles{
		{shapes: []Shape{Triangle{Side: 5}, Triangle{Side: 5}}, wantShapes: 2},
		{shapes: []Shape{Triangle{Side: 5}, Triangle{Side: 5}, Rectangle{Height: 4, Weight: 2}}, wantShapes: 2},
		{shapes: []Shape{}, wantShapes: 0},
	}
	for _, test := range tests {
		box := NewBox(len(test.shapes))
		for _, shape := range test.shapes {
			box.AddShape(shape)
		}
		err := box.RemoveAllCircles()
		if err == nil {
			t.Error("Circles are not exist in the list, then returns an error")
		}
		// if len(box.shapes) != test.wantShapes {
		// 	t.Errorf("test %d is fail, len want  %d, len got: %d", i, test.wantShapes, len(test.shapes))
		// }

	}

}
