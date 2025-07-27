package main

import (
    "testing"
)

func BenchmarkConcurrentMergeSort(b *testing.B) {
    for i := 0; i < b.N; i++ {
       concurrentMergeSort()
    }
}