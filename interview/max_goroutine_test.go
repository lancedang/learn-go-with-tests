package interview

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

// panic: too many concurrent operations on a single file or socket (max 1048575)
func TestMaxGoroutine(t *testing.T) {
	var x = int(10000000)
	for i := 0; i < x; i++ {
		go func() {
			fmt.Printf("hi %v  %v \n", i, runtime.NumGoroutine())
		}()
	}
	time.Sleep(1 * time.Hour)
}

func TestMaxGoroutineControl(t *testing.T) {

}
