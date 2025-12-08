package cmd2

import (
	"path/filepath"
	"testing"
)


func TestMain(t *testing.T) {
	ans := Main(filepath.Join("..", "..", "day2", "input2test"))
	expected := 4174379265
	if ans != expected {
        t.Errorf("Main = %d; want %d", ans, expected)
    }
 }

 func BenchmarkMain(b *testing.B) {
	for b.Loop() {
		Main(filepath.Join("..", "..", "day2", "input"))
	}
 }