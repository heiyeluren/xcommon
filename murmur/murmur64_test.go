package xcommon

import (
	"log"
	"testing"
)

func Test_Sum64(t *testing.T) {
	h64 := Sum64([]byte(`hello`))
	log.Println(h64)
}