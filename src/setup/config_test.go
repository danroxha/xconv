package setup

import "testing"


func BenchmarkNewConfiguration(b *testing.B) {
	NewConfiguration()
}