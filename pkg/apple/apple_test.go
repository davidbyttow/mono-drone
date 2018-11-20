package apple_test

import (
	"github.com/davidbyttow/mono-drone/pkg/apple"
	"testing"
)

func TestApple(t *testing.T) {
	if apple.Apple() != "apple" {
		t.Fail()
	}
}
