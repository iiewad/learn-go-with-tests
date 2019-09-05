package iteration

import "testing"

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}

func TestRepeat(t *testing.T) {
	repected := Repeat("a")
	expected := "aaaaa"

	if repected != expected {
		t.Errorf("expected '%s' but got '%s'", expected, repected)
	}
}
