package pin

import (
	"crypto/rand"
	"log"
	"math/big"
)

var randomNumberWrapper = func() (n *big.Int, err error) {
	return rand.Int(rand.Reader, big.NewInt(10))
}

func generateRandomNumber() int64 {
	nBig, err := randomNumberWrapper()
	if err != nil {
		log.Print(err.Error())
		panic(err)
	}
	n := nBig.Int64()
	return n
}

func isConsecutive(pin *[]int64, pos int) bool {
	for i := 0; i < len((*pin)[:pos])-1; i++ {
		if (*pin)[i] == (*pin)[i+1] {
			return true
		}
	}
	return false
}

func isIncrementalSeq(pin *[]int64, pos int) bool {
	count := 0
	for i := 0; i < len((*pin)[:pos])-1; i++ {
		if (*pin)[i+1]-(*pin)[i] == 1 {
			count++
		}
	}
	if count >= 2 {
		return true
	}
	return false
}

func isUnique(pins *[]string, pin string) bool {
	for _, p := range *pins {
		if p == pin {
			return false
		}
	}
	return true
}
