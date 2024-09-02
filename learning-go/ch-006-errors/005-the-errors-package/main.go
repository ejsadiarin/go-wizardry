package error

import (
	"errors"
)

func divide(x, y float64) (float64, error) {
	if y == 0 {
		// ?
	}
	return x / y, nil
}
