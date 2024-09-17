//write a test for a function that repeats a character 5 times.

package iteration

import "testing"

func TestMain(t *testing.T)  {
	repeated := repeat("a", 5)
	expected := "aaaaa"
	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
		
	}
} 	

func BenchmarkRepeat(b *testing.B){
	for i := 0; i < b.N; i++ {
		repeat("a", 5)
	}
}
