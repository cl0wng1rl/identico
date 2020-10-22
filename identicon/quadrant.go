package identicon

import (
	"strconv"
	"strings"

	"github.com/gabrielbarker/identico/hash"
	"github.com/gabrielbarker/identico/svg"
	"github.com/lucasb-eyer/go-colorful"
)

type QuadrantIdenticon struct{}

func (id *QuadrantIdenticon) Create(input string, quadrantWidth int) svg.SVG {
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

type QuadrantIdenticonData struct {
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

func convertIntsToPixelColors(vals string) []pixelColor {
	pixelColors := make([]pixelColor, len(vals))
	for i, v := range vals {
		pixelColors[i] = pixelColor(v - '0')
	}
	return pixelColors
}

func (id *QuadrantIdenticon) generateData(input string, quadrantWidth int) QuadrantIdenticonData {
	hashString := hash.Hash(input, quadrantWidth*quadrantWidth+4)
	firstColor := colorful.Hsv(convert2BitHexTo360Value(hashString[:2]), 1, 1).Hex()
	secondColor := colorful.Hsv(convert2BitHexTo360Value(hashString[2:4]), 1, 1).Hex()
	pixels := convertIntsToPixelColors(hash.ConvertHashToBase(hashString[4:], 3))
	return QuadrantIdenticonData{firstColor: firstColor, secondColor: secondColor, pixels: pixels}
}

func convert2BitHexTo360Value(hexNum string) float64 {
	val, _ := strconv.ParseInt(hexNum, 16, 16)
	return float64(360 * val / 256)
}

func getInputKey(input string) int {
	key := 0
	for _, char := range input {
		key += int(char)
	}
	return key % 100
}

func getColorForPixelColor(pixelIndex int, data QuadrantIdenticonData) string {
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
