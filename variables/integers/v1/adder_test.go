package integers

import "testing"

func TestAdd(t *testing.T) {
	sum := Add(2, 2)
	expected := 4
	if sum != expected {
		t.Errorf("Add(2,2)=%d; expected %d", sum, expected)
	}
}

// PS F:\go\goprojects\learning-go\variables\integers> go test -v
// === RUN   TestAdd
// --- PASS: TestAdd (0.00s)
// PASS
// ok      github.com/taiyangyi/learning-go/variables/integers     0.236s
