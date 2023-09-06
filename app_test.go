package main

import "testing"

func TestProgressTracker_Write(t *testing.T) {
	var readBytes int64
	pt := progressTracker{
		TotalSize: 10,
		Callback: func(rb int64) {
			readBytes += rb
		},
	}
	p := []byte("hello world")
	n, err := pt.Write(p)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if n != len(p) {
		t.Errorf("unexpected number of bytes written: got %d, want %d", n, len(p))
	}
	if readBytes != int64(len(p)) {
		t.Errorf("unexpected number of bytes read: got %d, want %d", readBytes, len(p))
	}
}
