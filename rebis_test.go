package rebis

import "testing"

var x []Item

func BenchmarkEmptyMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		x = make([]Item, 3)
	}
}
