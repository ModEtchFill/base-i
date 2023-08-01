package aTest
import (
	"fmt"
	"testing"
	)
func TestSomethingB(t *testing.T) {
	fmt.Printf("Test hi Something\n")
	if 1 == 0 {
		t.Errorf("test fail")
	}
}
