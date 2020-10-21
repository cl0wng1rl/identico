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
	quadrantWidth := 8
	data := id.generateData(input, quadrantWidth)
	maxIndex := 2*quadrantWidth - 1
	s := svg.New(quadrantWidth * 2)
	for j := 0; j < quadrantWidth; j++ {
		for i := 0; i < quadrantWidth; i++ {
			color := getColorForPixelColor(quadrantWidth*j+i, data)
			s.AddNewRectElement(color, i, j)
			s.AddNewRectElement(color, maxIndex-i, j)
			s.AddNewRectElement(color, i, maxIndex-j)
			s.AddNewRectElement(color, maxIndex-i, maxIndex-j)
		}
	}
	return *s
}

type simpleIdenticonData struct {
	firstColor  string
	secondColor string
	pixels      []pixelColor
}

type pixelColor int

const (
	CLEAR pixelColor = iota
	FIRST
	SECOND
)

func (id *SimpleIdenticon) generateData(input string, quadrantWidth int) simpleIdenticonData {
	key := getInputKey(input)
	firstColor := colorful.Hsv(3.6*float64((key*31)%100), 1, 1).Hex()
	secondColor := colorful.Hsv(3.6*float64((key*13)%100), 1, 1).Hex()
	pixels := getPixelColors(input, quadrantWidth*quadrantWidth)
	return simpleIdenticonData{firstColor: firstColor, secondColor: secondColor, pixels: pixels}
}

func getInputKey(input string) int {
	key := 0
	for _, char := range input {
		key += int(char)
	}
	return key % 100
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
