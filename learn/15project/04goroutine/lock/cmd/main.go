package main

import (
	"04goroutine/lock/service"
	"sync"
	"testing"
)

func main() {

}

func benchmark(b *testing.B, rw service.RW, read int, write int) {
	for i := 0; i < b.N; i++ {
		var wg sync.WaitGroup
		for j := 0; i < read*100; j++ {
			wg.Add(1)
			go func() {
				rw.Read()
				wg.Done()
			}()
		}

		for j := 0; i < write*100; j++ {
			wg.Add(1)
			go func() {
				rw.Write()
				wg.Done()
			}()
		}
		wg.Wait()
	}
}

func BenchmarkReadMore(b *testing.B)    { benchmark(b, &service.Lock{}, 9, 1) }
func BenchmarkReadMoreRW(b *testing.B)  { benchmark(b, &service.RWLock{}, 9, 1) }
func BenchmarkWriteMore(b *testing.B)   { benchmark(b, &service.Lock{}, 1, 9) }
func BenchmarkWriteMoreRW(b *testing.B) { benchmark(b, &service.RWLock{}, 1, 9) }
func BenchmarkEqual(b *testing.B)       { benchmark(b, &service.Lock{}, 5, 5) }
func BenchmarkEqualRW(b *testing.B)     { benchmark(b, &service.RWLock{}, 5, 5) }
