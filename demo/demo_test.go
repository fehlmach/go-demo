package demo_test

import (
	"testing"

	"github.com/fehlmach/go-demo/demo"
)

func TestHelloWorld(t *testing.T) {
	if demo.HelloWorld() != "Hello World" {
		t.Fatal("Expected Hello World")
	}
}
