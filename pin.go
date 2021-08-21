package pin

import (
	"fmt"
	"strings"
)

const MAX_BATCH int = 1000
const MAX_PIN_SIZE int = 4

func generatePIN() string {
	pin := make([]int64, MAX_PIN_SIZE)
	for y, pos := 0, 1; y < MAX_PIN_SIZE; y++ {
		pin[y] = generateRandomNumber()
		for isConsecutive(&pin, pos) || isIncrementalSeq(&pin, pos) {
			pin[y] = generateRandomNumber()
		}
		pos++
	}
	return strings.Trim(strings.Replace(fmt.Sprint(pin), " ", "", -1), "[]")
}

func GenerateBatchPIN() []string {
	pin := ""
	pins := make([]string, MAX_BATCH)
	for i := 0; i < MAX_BATCH; i++ {
		for !isUnique(&pins, pin) {
			pin = generatePIN()
		}
		pins[i] = pin
	}
	return pins
}
