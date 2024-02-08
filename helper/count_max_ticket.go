package helper

import "fmt"

func LeftZeroPad(number, padWidth int64) string {
	return fmt.Sprintf(fmt.Sprintf("%%0%dd", padWidth), number)
}
