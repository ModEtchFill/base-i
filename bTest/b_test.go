package aTest
import (
	"fmt"
	"testing"
	)
func TestSomething(t *testing.T) {
	fmt.Printf("Test b Something\n")
	if 1 == 0 {
		t.Errorf("test fail")
	}
}
