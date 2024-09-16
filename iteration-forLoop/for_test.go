//write a test for a function that repeats a character 5 times.

package iteration

import "testing"

func TestMain(t *testing.T)  {
	repeated := repeat("a")
	expected := "aaaaa"
	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
		
	}
} 	
