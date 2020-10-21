// svg/svg.go

package svg

import (
	"fmt"
	"math/rand"
	"strings"
)

type SVG struct {
	header        string
	innerElements []string
}

func New(width int) *SVG {
	return &SVG{header: defaultHeader(width), innerElements: []string{}}
}

func (s *SVG) ToString() string {
	inner := strings.Join(s.innerElements, "")
	return strings.Join([]string{s.header, inner, closingTag()}, "")
}

func (s *SVG) AddElement(element string) {
	s.innerElements = append(s.innerElements, element)
}

func (s *SVG) AddNewRectElement(fillHex string, x int, y int) {
	s.innerElements = append(s.innerElements, newRectElement(fillHex, x, y, 1, 1))
}

func defaultHeader(width int) string {
	return fmt.Sprintf(`<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 %v %v" shape-rendering="crispEdges">`, width, width)
}

func closingTag() string {
	return `</svg>`
}

func NewRandomSquare(x int, y int) string {
	return newRectElement(randomHexColor(), x, y, 1, 1)
}

func newRectElement(fillHex string, x int, y int, width int, height int) string {
	return fmt.Sprintf(`<rect fill="%v" x="%v" y="%v" width="%v" height="%v" />`, fillHex, x, y, width, height)
}

func randomHexColor() string {
	return fmt.Sprintf("#%v%v%v", randomHexByte(), randomHexByte(), randomHexByte())
}

func randomHexByte() string {
	return fmt.Sprintf("%02x", rand.Intn(256))
}
