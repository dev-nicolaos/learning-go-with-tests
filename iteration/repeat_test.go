package iteration

import "testing"

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 6)
	expected := "aaaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func ExampleRepeat() {
	fmt.Println(Repeat("T", 4))
	// Output: "TTTT"
}

func BenchmarkRepeat(b *testing.B) {
	for range b.N {
		Repeat("a", 7)
	}
}
