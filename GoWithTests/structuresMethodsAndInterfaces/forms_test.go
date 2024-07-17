package structuresMethodsAndInterfaces

import (
	"strconv"
	"testing"
)

/*
	func TestArea(t *testing.T) {
		verifyArea := func(t *testing.T, form Form, expected float64) {
			t.Helper()
			result, _ := strconv.ParseFloat(strconv.FormatFloat(form.Area(), 'f', 2, 64), 64)
			if result != expected {
				t.Errorf("Rectangle.Area() = %v; expected %v", result, expected)
			}
		}
		t.Run("Should test Rectangle Area", func(t *testing.T) {
			rectangle := Rectangle{Width: 12.0, Height: 6.0}
			expected := 72.0
			verifyArea(t, rectangle, expected)
		})

		t.Run("Should test Circle Area", func(t *testing.T) {
			circle := Circle{Radius: 5}
			expected := 78.54
			verifyArea(t, circle, expected)
		})
	}
*/

func TestArea(t *testing.T) {
	verifyArea := func(t *testing.T, form Form, expected float64) {
		t.Helper()
		result, _ := strconv.ParseFloat(strconv.FormatFloat(form.Area(), 'f', 2, 64), 64)
		if result != expected {
			t.Errorf("%#v resultado %.2f, esperado %.2f", t, result, expected)

		}
	}
	testCases := []struct {
		name    string
		form    Form
		hasArea float64
	}{
		{"Rectangle", Rectangle{12, 6}, 72.0},
		{"Circle", Circle{10}, 314.16},
		{"Triangle", Triangle{12, 6}, 36},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			verifyArea(t, testCase.form, testCase.hasArea)
		})
	}

}

/*func TestPerimeter(t *testing.T) {
	t.Run("Should test Rectangle Perimeter", func(t *testing.T) {
		rectangle := Rectangle{Width: 10, Height: 10.0}
		result := rectangle.Perimeter()

		expected := 40.0

		if result != expected {
			t.Errorf("Rectangle.Perimeter() = %v; expected %v", result, expected)
		}
	})

}
*/
