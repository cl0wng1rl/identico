package identicon

import (
	"io/ioutil"
	"testing"
)

func TestQuadrantIdenticonCreate(t *testing.T) {
	id := QuadrantIdenticon{}
	svgElement := id.Create("test input", 8)
	if svgElement.ToString() != string(expectedSVG) {
		t.Errorf("created svg should be: %v , instead got: %v", expectedSVG, svgElement.ToString())
	}
}

var expectedSVG, _ = ioutil.ReadFile("./quadrant_snapshot.txt")
