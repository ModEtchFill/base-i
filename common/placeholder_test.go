package common
import (
	"fmt"
	"testing"
	)
func TestPlaceholder(t *testing.T) {
	fmt.Printf("Test placeholdr\n")
	if 1 == 0 {
		t.Errorf("test fail")
	}
}
