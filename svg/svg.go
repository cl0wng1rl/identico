// svg/svg.go

package svg

import (
	"fmt"
	"math/rand"
)

type SVG struct {
	header        string
	innerElements []string
}

func New() *SVG {
	return &SVG{header: defaultHeader(), innerElements: []string{}}
}

func (s *SVG) ToString() string {
	inner := ""
	for _, element := range s.innerElements {
		inner += element
	}
	return fmt.Sprintf(`%v%v%v`, s.header, inner, closingTag())
}

func (s *SVG) AddElement(element string) {
	s.innerElements = append(s.innerElements, element)
}

func defaultHeader() string {
	return `<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 7 8" shape-rendering="crispEdges">`
}

func closingTag() string {
	return `</svg>`
}

func NewRandomSquare(x int, y int) string {
	return newRectElement(fmt.Sprintf("%x", rand.Uint32()), x, y, 1, 1)
}

func newRectElement(fillHex string, x int, y int, width int, height int) string {
	return fmt.Sprintf(`<rect fill="%v" x="%v" y="%v" width="%v" height="%v" />`, fillHex, x, y, width, height)
}
