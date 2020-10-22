package identicon

import (
	"github.com/gabrielbarker/identico/svg"
)

type Identicon interface {
	Create(input string) svg.SVG
}
