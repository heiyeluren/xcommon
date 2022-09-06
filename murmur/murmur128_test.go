package common

import (
	"log"
	"testing"
)

func Test_Sum128(t *testing.T) {
	h128_1,h128_2 := Sum128([]byte(`hello`))
	log.Println(h128_1,h128_2)
}