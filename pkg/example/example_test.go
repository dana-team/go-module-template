package example_test

import "testing"

func TestHelloWorld(t *testing.T) {
	if helloworld() != "Hello World!!" {
		t.Fatal("Test fail")
	}
}
