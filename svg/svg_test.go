// svg/svg_test.go

package svg

import "testing"

func TestToString(t *testing.T) {
	svgElement := New()
	expectedSVG := `<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 15 16" shape-rendering="crispEdges"></svg>`
	if svgElement.ToString() != expectedSVG {
		t.Errorf("default svg ToString should be: %v , instead got: %v", expectedSVG, svgElement.ToString())
	}
}

func TestAddElementAndToString(t *testing.T) {
	svgElement := New()
	svgElement.AddElement(newRectElement("testColor", 0, 0, 1, 1))

	expectedSVG := `<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 15 16" shape-rendering="crispEdges"><rect fill="testColor" x="0" y="0" width="1" height="1" /></svg>`
	if svgElement.ToString() != expectedSVG {
		t.Errorf("svg with inner element ToString should be: %v , instead got: %v", expectedSVG, svgElement.ToString())
	}
}
