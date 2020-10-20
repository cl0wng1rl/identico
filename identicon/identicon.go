package identicon

import (
	"strings"

	"github.com/gabrielbarker/identico/svg"
	"github.com/lucasb-eyer/go-colorful"
)

type Identicon interface {
	Create(input string) svg.SVG
}

type SimpleIdenticon struct{}

func (id *SimpleIdenticon) Create(input string) svg.SVG {
	data := id.generateData(input)
	s := svg.New()
	for j := 0; j < 8; j++ {
		for i := 0; i < 8; i++ {
			color := getColorForPixelColor(8*j+i, data)
			element := svg.NewRectElement(color, i, j, 1, 1)
			s.AddElement(element)
			element = svg.NewRectElement(color, 15-i, j, 1, 1)
			s.AddElement(element)
			element = svg.NewRectElement(color, i, 15-j, 1, 1)
			s.AddElement(element)
			element = svg.NewRectElement(color, 15-i, 15-j, 1, 1)
			s.AddElement(element)
		}
	}
	return *s
}

type simpleIdenticonData struct {
	firstColor  string
	secondColor string
	pixels      [64]pixelColor
}

type pixelColor int

const (
	CLEAR pixelColor = iota
	FIRST
	SECOND
)

func (id *SimpleIdenticon) generateData(input string) simpleIdenticonData {
	firstColor := colorful.Hsv(3.6*float64((len(input)*31)%100), 1, 1).Hex()
	secondColor := colorful.Hsv(3.6*float64((len(input)*13)%100), 1, 1).Hex()
	pixels := [64]pixelColor{}
	copy(pixels[:], getPixelColors(input, 64)[:64])
	return simpleIdenticonData{firstColor: firstColor, secondColor: secondColor, pixels: pixels}
}

func getColorForPixelColor(pixelIndex int, data simpleIdenticonData) string {
	if data.pixels[pixelIndex] == FIRST {
		return data.firstColor
	} else if data.pixels[pixelIndex] == SECOND {
		return data.secondColor
	} else {
		return "#00000000"
	}
}

func getPixelColors(input string, length int) []pixelColor {
	messageToRepeat := "message to encode"
	toEncode := strings.Repeat(messageToRepeat, 1+(length/len(messageToRepeat)))
	pixels := []pixelColor{}
	for i := 0; i < length; i++ {
		pixels = append(pixels, pixelColor((int(input[i%len(input)])+int(toEncode[i])+i)%3))
	}
	return pixels
}
