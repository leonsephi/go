package main

import "testing"

func BenchmarkReflectUint(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		checkIDNotZero(uint(10))
	}
}

func BenchmarkReflectInt(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		checkIDNotZero(int(10))
	}
}
