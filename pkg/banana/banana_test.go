package banana_test

import (
	"github.com/davidbyttow/mono-drone/pkg/banana"
	"testing"
)

func TestBanana(t *testing.T) {
	if banana.Banana() != "banana" {
		t.Fail()
	}
}
