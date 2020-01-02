package decotail

import (
	"fmt"
	"github.com/aybabtme/color/brush"
)

func DefineColor(index int, text string) string {
	switch index {
	case 0:
		return fmt.Sprint(brush.Red(text))
	case 1:
		return fmt.Sprint(brush.Blue(text))
	case 2:
		return fmt.Sprint(brush.Yellow(text))
	case 3:
		return fmt.Sprint(brush.Green(text))
	case 4:
		return fmt.Sprint(brush.Purple(text))
	case 5:
		return fmt.Sprint(brush.Cyan(text))
	default:
		return text

	}
}
