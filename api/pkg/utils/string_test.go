package utils

import (
	"fmt"
	"testing"
)

func TestTransNumToString(t *testing.T) {
	var numberSlince = []int64{1, 10, 100, 240, 1000, 10000, 100000, 1000000, 10000000}

	for _, number := range numberSlince {
		str := NumToString(number)
		fmt.Printf("%d short key is %s \n", number, str)
	}
}
