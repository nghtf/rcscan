package rcscan

import (
	"testing"
)

func TestNewEmpty(t *testing.T) {
	_, err := New("")
	if err == nil {
		t.Fatalf(`New("") = _, %v, want nil, error`, err)
	}
}

func TestNewFile404(t *testing.T) {
	_, err := New("404.file")
	if err == nil {
		t.Fatalf(`New("404.file") = _, %v, want nil, error`, err)
	}
}

func TestGetSection404(t *testing.T) {
	rc, _ := New("./example/example.rc")
	value, err := rc.Get("Section 404", "paramA")
	if err == nil {
		t.Fatalf(`Get("Section 404", "paramA") = %v, %v, want "", error`, value, err)
	}
}
