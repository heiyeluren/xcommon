package pool

import (
	"fmt"
	"testing"
)

func Test_Pool(t *testing.T) {
	pool := New(12)
	pool.NewTask(func(){
		fmt.Println("run task test")
	})
}