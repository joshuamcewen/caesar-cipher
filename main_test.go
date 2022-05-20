package main

import "testing"

func TestShift(t *testing.T) {
	buffer := shift([]byte("ABZ abz 0!"))

	if buffer.String() != "CDB cdb 0!" {
		t.Errorf("Unexpected result: %s", buffer.String())
	}
}

func BenchmarkApplication(t *testing.B) {
	main()
}
